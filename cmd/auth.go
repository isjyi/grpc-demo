/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"path"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/isjyi/grpc-demo/pb"
	"github.com/isjyi/grpc-demo/pkg/ui/data/swagger"
	"github.com/isjyi/grpc-demo/service/auth"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	port          string
	httpPort      string
	swaggerSwitch bool
	gatewaySwitch bool
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "用户接口",
	Long:  `这是用户登陆`,
	Run: func(cmd *cobra.Command, args []string) {
		if gatewaySwitch {
			if err := gatewayRun(); err != nil {
				log.Fatal("cannot start server: ", err)
			}
		} else {
			if err := grpcRun(); err != nil {
				log.Fatal("cannot start server: ", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.Flags().StringVarP(&port, "port", "p", ":2000", "grpc service start port")
	authCmd.Flags().StringVar(&httpPort, "http_port", ":5000", "grpc-gateway service start port")
	authCmd.Flags().BoolVar(&swaggerSwitch, "swagger", false, "Whether to enable swagger service (default false)")
	authCmd.Flags().BoolVar(&gatewaySwitch, "gateway", false, "Whether to enable http service (default false)")
}

func grpcRun() error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to listen: %v", err))
	}

	log.Printf("listen on %s\n", port)

	server := grpc.NewServer()

	pb.RegisterAuthServiceServer(server, auth.NewAuthSrv())

	if err := server.Serve(listener); err != nil {
		return errors.New(fmt.Sprintf("failed to serve: %v", err))
	}
	return nil
}

func gatewayRun() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterAuthServiceHandlerFromEndpoint(ctx, gwmux, port, opts)
	if err != nil {
		return errors.New(fmt.Sprintf("grpc-gateway to serve: %v", err))
	}
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	if swaggerSwitch {
		mux.HandleFunc("/swagger/", serveSwaggerFile)
		serveSwaggerUI(mux)
	}

	log.Printf("grpc-gateway on: %s\n", httpPort)

	return http.ListenAndServe(httpPort, mux)
}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join("./swagger", p)

	log.Printf("Serving swagger-file: %s", p)

	http.ServeFile(w, r, p)
}

func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "/third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
