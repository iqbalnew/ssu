package manager

import (
	"context"
	"fmt"
	"os"
	"time"

	authPb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/auth_service"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	Username  string `json:"username"`
	UserID    uint64 `json:"user_id"`
	SessionID string `json:"session_id"`
	DateTime  string `json:"date_time"`
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
