package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"regexp"

	"google.golang.org/grpc"

	"fmt"
	"net"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"

	"bitbucket.bri.co.id/scm/addons/addons-task-service/server/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	"github.com/urfave/cli"

	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

const defaultPort = 9090
const serviceName = "Task"

var s *grpc.Server

func main() {

	initConfig()

	app := cli.NewApp()
	app.Name = ""
	app.Commands = []cli.Command{
		grpcServerCmd(),
		gatewayServerCmd(),
		grpcGatewayServerCmd(),
		runMigrationCmd(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}

func grpcServerCmd() cli.Command {
	return cli.Command{
		Name:  "grpc-server",
		Usage: "starts a gRPC server",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port",
				Value: defaultPort,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.Int("port")

			startDBConnection()

			go func() {
				if err := grpcServer(port); err != nil {
					logrus.Fatalf("failed RPC serve: %v", err)
				}
			}()

			// Wait for Control C to exit
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			// Block until a signal is received
			<-ch

			closeDBConnections()

			logrus.Println("Stopping RPC server")
			s.Stop()
			logrus.Println("RPC server stopped")
			return nil
		},
	}
}

func gatewayServerCmd() cli.Command {
	return cli.Command{
		Name:  "gw-server",
		Usage: "starts a Gateway server",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port",
				Value: 3000,
			},
			cli.StringFlag{
				Name:  "grpc-endpoint",
				Value: ":" + fmt.Sprint(defaultPort),
				Usage: "the address of the running gRPC server to transcode to",
			},
		},
		Action: func(c *cli.Context) error {
			port, grpcEndpoint := c.Int("port"), c.String("grpc-endpoint")

			go func() {
				if err := httpGatewayServer(port, grpcEndpoint); err != nil {
					logrus.Fatalf("failed JSON Gateway serve: %v", err)
				}
			}()

			// Wait for Control C to exit
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			// Block until a signal is received
			<-ch

			logrus.Println("JSON Gateway server stopped")

			return nil
		},
	}
}

func grpcGatewayServerCmd() cli.Command {
	return cli.Command{
		Name:  "grpc-gw-server",
		Usage: "Starts gRPC and Gateway server",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port1",
				Value: defaultPort,
			},
			cli.IntFlag{
				Name:  "port2",
				Value: 3000,
			},
			cli.StringFlag{
				Name:  "grpc-endpoint",
				Value: ":" + fmt.Sprint(defaultPort),
				Usage: "the address of the running gRPC server to transcode to",
			},
		},
		Action: func(c *cli.Context) error {
			rpcPort, httpPort, grpcEndpoint := c.Int("port1"), c.Int("port2"), c.String("grpc-endpoint")

			startDBConnection()

			go func() {
				if err := grpcServer(rpcPort); err != nil {
					logrus.Fatalf("failed RPC serve: %v", err)
				}
			}()

			go func() {
				if err := httpGatewayServer(httpPort, grpcEndpoint); err != nil {
					logrus.Fatalf("failed JSON Gateway serve: %v", err)
				}
			}()

			// Wait for Control C to exit
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			// Block until a signal is received
			<-ch

			closeDBConnections()

			logrus.Println("Stopping RPC server")
			s.GracefulStop()
			logrus.Println("RPC server stopped")
			logrus.Println("JSON Gateway server stopped")

			return nil
		},
	}
}

func grpcServer(port int) error {
	// RPC
	logrus.Printf("Starting %s Service ................", serviceName)
	logrus.Printf("Starting RPC server on port %d...", port)
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	apiServer := api.New(db_main, announcementConn)
	authInterceptor := api.NewAuthInterceptor(apiServer.GetManager())

	unaryInterceptorOpt := grpc.UnaryInterceptor(api.UnaryInterceptors(authInterceptor))
	streamInterceptorOpt := grpc.StreamInterceptor(api.StreamInterceptors(authInterceptor))

	s = grpc.NewServer(unaryInterceptorOpt, streamInterceptorOpt)
	pb.RegisterTaskServiceServer(s, apiServer)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	return s.Serve(list)
}

func httpGatewayServer(port int, grpcEndpoint string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Connect to the GRPC server
	conn, err := grpc.Dial(grpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	rmux := runtime.NewServeMux(
		runtime.WithErrorHandler(CustomHTTPError),
	)
	// opts := []grpc.DialOption{grpc.WithInsecure()}
	// err := pb.RegisterBaseServiceHandlerFromEndpoint(ctx, rmux, grpcEndpoint, opts)
	client := pb.NewTaskServiceClient(conn)
	err = pb.RegisterTaskServiceHandlerClient(ctx, rmux, client)
	if err != nil {
		return err
	}

	// Serve the swagger-ui and swagger file
	mux := http.NewServeMux()
	mux.Handle("/", rmux)

	mux.HandleFunc("/api/task/docs/swagger.json", serveSwagger)
	fs := http.FileServer(http.Dir("www/swagger-ui"))
	mux.Handle("/api/task/docs/", http.StripPrefix("/api/task/docs/", fs))

	// Start
	logrus.Printf("Starting JSON Gateway server on port %d...", port)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), cors(mux))
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "www/swagger.json")
}

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "https://addons-web-client.vercel.app" {
		return true
	}
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, OPTION, POST, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
