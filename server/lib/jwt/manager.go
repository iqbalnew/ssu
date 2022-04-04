package manager

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	customAES "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/aes"
	authPb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/auth_service"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	Username     string              `json:"username"`
	UserID       uint64              `json:"user_id"`
	SessionID    string              `json:"session_id"`
	DateTime     string              `json:"date_time"`
	UserType     string              `json:"user_type"`
	ProductRoles []*ProductAuthority `json:"product_roles"`
	Authorities  []string            `json:"authorities"`
	CompanyIDs   string              `json:"company_ids"`
}

type CurrentUser struct {
	UserClaims
	FilterMe    string   `json:"filter_me"`
	StatusOrder []string `json:"status_order"`
	TaskFilter  string   `json:"task_filter"`
}

type VerifyTokenRes struct {
	IsValid      bool                `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
	IsExpired    bool                `protobuf:"varint,2,opt,name=isExpired,proto3" json:"isExpired,omitempty"`
	UserID       uint64              `protobuf:"varint,3,opt,name=userID,proto3" json:"userID,omitempty"`
	Username     string              `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	UserType     string              `protobuf:"bytes,5,opt,name=userType,proto3" json:"userType,omitempty"`
	ProductRoles []*ProductAuthority `protobuf:"bytes,6,rep,name=productRoles,proto3" json:"productRoles,omitempty"`
}

type ProductAuthority struct {
	ProductName string   `protobuf:"bytes,1,opt,name=productName,proto3" json:"productName,omitempty"`
	Authorities []string `protobuf:"bytes,2,rep,name=authorities,proto3" json:"authorities,omitempty"`
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey, tokenDuration}
}

func (manager *JWTManager) Generate(username string, userID uint64, sessionID string, dateTime string) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		Username: username,
		UserID:   userID,
	}

	if sessionID != "" && dateTime != "" {
		claims.SessionID = sessionID
		claims.DateTime = dateTime
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func (manager *JWTManager) GetMeFromJWT(ctx context.Context, accessToken string) (*CurrentUser, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := md["authorization"]
		if len(values) > 0 {
			split := strings.Split(values[0], " ")
			accessToken = split[0]
			if len(split) > 1 {
				accessToken = split[1]
			}
		}

	}

	if accessToken == "" {
		logrus.Errorf("access token is empty")
		return nil, status.Error(codes.Unauthenticated, "session is empty")
	}

	userClaims, err := manager.Verify(accessToken)
	if err != nil {
		logrus.Errorf("[api.task][func:GetMeFromJWT][01] failed to verify token '%s', error: %v", accessToken, err)
		return nil, status.Errorf(codes.Unauthenticated, "Session expired")
	}

	now := time.Now()

	fmt.Printf("token verify expired: %v|%v|%v", !(now.Unix() <= userClaims.ExpiresAt), now.String(), time.Unix(userClaims.ExpiresAt, 0).String())
	if !(now.Unix() <= userClaims.ExpiresAt) {
		return nil, status.Errorf(codes.Unauthenticated, "Session expired")
	}

	logrus.Println(userClaims.ProductRoles)
	for _, v := range userClaims.ProductRoles {
		if v.ProductName == "User" {
			for _, j := range v.Authorities {
				data := strings.Split(j, ":")
				if data[1] != "-" {
					userClaims.Authorities = append(userClaims.Authorities, data[1])
					break
				}
			}
		}
	}
	logrus.Println(userClaims.Authorities)

	currentUser := &CurrentUser{
		UserClaims: *userClaims,
	}
	// - Maker: 1. Draft, 2. Returned, 3. Pending, 4. Request for Delete, 5. Approved, 6. Rejected
	// - Signer: 1. Pending, 2. Request for Delete, 3. Approved, 4. Rejected
	if len(userClaims.Authorities) > 0 {
		switch strings.ToLower(userClaims.Authorities[0]) {
		case "maker":
			currentUser.StatusOrder = []string{"2", "3", "1", "6", "4", "5"}
			currentUser.FilterMe = "status:<>0,status:<>7"

		case "signer":
			currentUser.StatusOrder = []string{"1", "6", "4", "5"}
			currentUser.FilterMe = "status:<>0,status:<>2,status:<>3,status:<>7"

		default:
			return nil, status.Errorf(codes.PermissionDenied, "Authority Denied")
		}
	} else {
		return nil, status.Errorf(codes.PermissionDenied, "Authority Denied")
	}

	currentUser.TaskFilter = ""
	if currentUser.UserType == "ca" || currentUser.UserType == "cu" {
		currentUser.TaskFilter = "data.user.companyID:"

		key := getEnv("JWT_AES_KEY", "Odj12345*12345678901234567890123")
		aes := customAES.NewCustomAES(key)

		decrypted, err := aes.Decrypt(userClaims.CompanyIDs)
		if err != nil {
			logrus.Errorf("[api.auth][func:VerifyToken][05] Failed to decrypt companyIDs: %v", err)
			return nil, status.Errorf(codes.Internal, "Server error")
		}

		if decrypted != "" {
			var ids []uint64
			err = json.Unmarshal([]byte(decrypted), &ids)
			if err != nil {
				logrus.Errorf("[api.auth][func:VerifyToken][06] Failed to unmarshal companyIDs: %v", err)
				return nil, status.Errorf(codes.Internal, "Server error")
			}

			for i, v := range ids {
				if i == 0 {
					currentUser.TaskFilter = currentUser.TaskFilter + fmt.Sprintf("%d", v)
				} else {
					currentUser.TaskFilter = currentUser.TaskFilter + fmt.Sprintf(",%d", v)
				}
			}
		}
	}

	return currentUser, nil
}

func (manager *JWTManager) GetMeFromAuthService(ctx context.Context, accessToken string) (*VerifyTokenRes, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	authConn, err := grpc.Dial(getEnv("AUTH_SERVICE", ":9105"), opts...)
	if err != nil {
		logrus.Errorln("Failed connect to Task Service: %v", err)
		return nil, status.Errorf(codes.Internal, "Error Internal")
	}
	defer authConn.Close()

	authClient := authPb.NewApiServiceClient(authConn)

	dataUser, err := authClient.VerifyToken(ctx, &authPb.VerifyTokenReq{
		AccessToken: accessToken,
	})
	if err != nil {
		return nil, err
	}
	if dataUser == nil {
		return nil, status.Errorf(codes.Aborted, "Failed To Get Data User")
	}

	user := &VerifyTokenRes{
		IsValid:   dataUser.IsValid,
		IsExpired: dataUser.IsExpired,
		UserID:    dataUser.UserID,
		Username:  dataUser.Username,
		UserType:  dataUser.UserType,
	}

	for _, v := range dataUser.ProductRoles {
		role := &ProductAuthority{
			ProductName: v.ProductName,
			Authorities: v.Authorities,
		}
		user.ProductRoles = append(user.ProductRoles, role)
	}
	return user, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
