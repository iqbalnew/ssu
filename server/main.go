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
	addonsLogger "bitbucket.bri.co.id/scm/addons/addons-task-service/server/logger"
	mongoClient "bitbucket.bri.co.id/scm/addons/addons-task-service/server/mongodb"

	"bitbucket.bri.co.id/scm/addons/addons-task-service/server/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"

	"github.com/spf13/viper"
	"github.com/urfave/cli"

	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	// "github.com/gorilla/handlers"
)

const defaultPort = 9090
const serviceName = "Task"

var s *grpc.Server

func main() {

	// fn := logOutput()
	// defer fn()
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

			sid, err := shortid.New(1, shortid.DefaultABC, 2342)
			if err != nil {
				panic(err)
			}

			var logger *addonsLogger.Logger
			if getEnv("ENV", "LOCAL") != "LOCAL" {
				logrus.Println("[stating utility] Connecting to Fluentd ")
				logger = addonsLogger.NewLogger(config.LoggerPort, config.LoggerHost, config.LoggerTag)
				defer logger.Close()
			}

			go func() {
				if err := grpcServer(port, sid, logger); err != nil {
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

			sid, err := shortid.New(1, shortid.DefaultABC, 2342)
			if err != nil {
				panic(err)
			}

			var logger *addonsLogger.Logger
			if getEnv("ENV", "LOCAL") != "LOCAL" {
				logrus.Println("[stating utility] Connecting to Fluentd ")
				logger = addonsLogger.NewLogger(config.LoggerPort, config.LoggerHost, config.LoggerTag)
				defer logger.Close()
			}

			go func() {
				if err := httpGatewayServer(port, grpcEndpoint, sid, logger); err != nil {
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

			sid, err := shortid.New(1, shortid.DefaultABC, 2342)
			if err != nil {
				panic(err)
			}

			var logger *addonsLogger.Logger
			if getEnv("ENV", "LOCAL") != "LOCAL" {
				logrus.Println("[stating utility] Connecting to Fluentd ")
				logger = addonsLogger.NewLogger(config.LoggerPort, config.LoggerHost, config.LoggerTag)
				defer logger.Close()
			}

			go func() {
				if err := grpcServer(rpcPort, sid, logger); err != nil {
					logrus.Fatalf("failed RPC serve: %v", err)
				}
			}()

			go func() {
				if err := httpGatewayServer(httpPort, grpcEndpoint, sid, logger); err != nil {
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

func grpcServer(port int, sid *shortid.Shortid, logger *addonsLogger.Logger) error {
	// RPC
	logrus.Printf("Starting %s Service ................", serviceName)
	logrus.Printf("Starting RPC server on port %d...", port)
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	var mongodbClient *mongoClient.MongoDB
	if getEnv("ENV", "LOCAL") != "LOCAL" {
		logrus.Println("[starting utilit] Connecting to Mongo ")
		mongodbClient = mongoClient.NewCLient(config.MongoURI, config.MongoDB, "logs")
		defer mongodbClient.Close()
	}

	apiServer := api.New(db_main, sid, announcementConn, mongodbClient, logger)
	authInterceptor := api.NewAuthInterceptor(apiServer.GetManager())

	unaryInterceptorOpt := grpc.UnaryInterceptor(api.UnaryInterceptors(authInterceptor))
	streamInterceptorOpt := grpc.StreamInterceptor(api.StreamInterceptors(authInterceptor))

	s = grpc.NewServer(unaryInterceptorOpt, streamInterceptorOpt)
	pb.RegisterTaskServiceServer(s, apiServer)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	return s.Serve(list)
}

func httpGatewayServer(port int, grpcEndpoint string, sid *shortid.Shortid, logger *addonsLogger.Logger) error {
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
		runtime.WithForwardResponseOption(httpResponseModifier),
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

	var handler http.Handler = mux
	handler = setHeaderHandler(handler, sid)
	handler = logRequestHandler(handler, logger)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
		// ReadTimeout:  120 * time.Second,
		// WriteTimeout: 120 * time.Second,
		// IdleTimeout:  120 * time.Second, // introduced in Go 1.8
	}

	return srv.ListenAndServe()
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "www/swagger.json")
}

func allowedOrigin(origin string) bool {
	if stringInSlice(viper.GetString("cors"), config.CorsAllowedOrigins) {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}
