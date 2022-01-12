package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"fmt"
	"net"

	pb "github.com/ordentco/go-base/server/pb"

	"github.com/ordentco/go-base/server/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

const (
	defaultPort = 3035
)

func main() {

	initConfig()

	app := cli.NewApp()
	app.Name = ""
	app.Commands = []cli.Command{
		grpcServerCmd(),
		gatewayServerCmd(),
		allServerCmd(),
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

			// initDBConnection()

			logrus.Printf("Starting RPC server on port %d...", port)
			list, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
			if err != nil {
				return err
			}

			// interceptorOpt := grpc.UnaryInterceptor(api.Interceptors())
			interceptorOpt := grpc.UnaryInterceptor(nil)
			s := grpc.NewServer(interceptorOpt)
			pb.RegisterGoBaseServiceServer(s, api.New())

			if err := s.Serve(list); err != nil {
				return err
			}

			// closeDBConnections()

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
				Value: 3001,
			},
			cli.StringFlag{
				Name:  "grpc-endpoint",
				Value: "127.0.0.1:3035",
				Usage: "the address of the running gRPC server to transcode to",
			},
		},
		Action: func(c *cli.Context) error {
			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			mux := runtime.NewServeMux()
			fs := http.FileServer(http.Dir("./swagger-ui"))
			http.Handle("/doc", http.StripPrefix("/doc", fs))

			opts := []grpc.DialOption{grpc.WithInsecure()}
			err := pb.RegisterGoBaseServiceHandlerFromEndpoint(ctx, mux, c.String("grpc-endpoint"), opts)
			if err != nil {
				return err
			}

			logrus.Printf("Starting JSON Gateway server on port %d...", c.Int("port"))

			return http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", c.Int("port")), mux)
		},
	}
}

func allServerCmd() cli.Command {
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
				Value: 3001,
			},
			cli.StringFlag{
				Name:  "grpc-endpoint",
				Value: "127.0.0.1:3035",
				Usage: "the address of the running gRPC server to transcode to",
			},
		},
		Action: func(c *cli.Context) error {
			port1, port2 := c.Int("port1"), c.Int("port2")

			// initDBConnection()

			// RPC
			logrus.Printf("Starting RPC server on port %d...", port1)
			list, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port1))
			if err != nil {
				return err
			}

			// interceptorOpt := grpc.UnaryInterceptor(api.Interceptors())
			interceptorOpt := grpc.UnaryInterceptor(nil)

			s := grpc.NewServer(interceptorOpt)
			pb.RegisterGoBaseServiceServer(s, api.New())

			// Gatewaty
			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			mux := runtime.NewServeMux()

			opts := []grpc.DialOption{grpc.WithInsecure()}
			err = pb.RegisterGoBaseServiceHandlerFromEndpoint(ctx, mux, c.String("grpc-endpoint"), opts)
			if err != nil {
				return err
			}

			go func() {
				if err := s.Serve(list); err != nil {
					log.Fatalf("failed RPC serve: %v", err)
				}
			}()

			go func() {
				logrus.Printf("Starting JSON Gateway server on port %d...", port2)

				if err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port2), mux); err != nil {
					log.Fatalf("failed Gateway serve: %v", err)
				}
			}()

			// Wait for Control C to exit
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			// Block until a signal is received
			<-ch

			// closeDBConnections()

			fmt.Println("Stopping the server")
			s.Stop()

			return nil
		},
	}
}

func swaggerHandler(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	fs := http.FileServer(http.Dir("../swagger-ui"))
	// // w.Write([]byte("hello " + pathParams["name"]))
	// // Serve Swagger Doc
	// // fs := http.FileServer(http.Dir("../swagger-ui"))
	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))
	// http.ServeFile(w, r, "./swagger-ui")
	// path := filepath.Join("./swagger-ui", "index.html")
	// fileName := "index.html"
	// if pathParams["fileName"] != "" {
	// 	path = filepath.Join("./swagger-ui", pathParams["fileName"])
	// 	fileName = pathParams["fileName"]
	// }
	// fmt.Println("--> %v", fileName)

	// tmpl, err := template.ParseFiles(path)
	// if err != nil {
	// 	log.Fatalf("Swagger Parse File Failure : %v", err)
	// }

	// fmt.Println("--> %v", fp)

	// if err := tmpl.ExecuteTemplate(w, fileName, nil); err != nil {
	// 	log.Fatalf("Execute Swagger Template Failure : %v", err)
	// }
}
