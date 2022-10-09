package main

import (
	"context"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	deliveryGRPC "github.com/godpeny/gos/delivery/grpc"
	"github.com/godpeny/gos/ent"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io"
	"net"
	"os"
)

var (
	addr      = flag.String("addr", ":8080", "endpoint of the gRPC service")
	network   = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
	enableTls = false
)

func main() {
	// set grpc logging
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)
	log.Infoln("GRPC logger set up properly.")

	// set up tcp server
	lis, err := net.Listen(*network, *addr)
	if err != nil {
		log.Fatalf("error listening server : %v", err)
	}
	defer func() {
		if err = lis.Close(); err != nil {
			log.Fatalf("Failed to close %s %s: %v", *network, *addr, err)
		}
	}()

	// set db
	entClient, err := ent.Open("mysql", "root:toor@tcp(localhost:3306)/gos?charset=utf8&parseTime=True&loc=UTC")
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	defer entClient.Close()
	// Run the auto migration tool.
	if err = entClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Infoln("DB set up properly.")

	// set up grpc
	var opts []grpc.ServerOption
	if enableTls {
		// implement when tls enabled.
		//
		//if *certFile == "" {
		//	*certFile = data.Path("x509/server_cert.pem")
		//}
		//if *keyFile == "" {
		//	*keyFile = data.Path("x509/server_key.pem")
		//}
		//creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		//if err != nil {
		//	log.Fatalf("Failed to generate credentials %v", err)
		//}
		//opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	// register multiple services
	deliveryGRPC.RegisterUserServiceServer(grpcServer, deliveryGRPC.NewUserService(entClient))
	log.Infoln("GRPC set up properly.")

	// TODO(@godpeny): implementing graceful shutdown when received OS signals.
	ctx := context.Background()
	go func() {
		defer func() {
			log.Infoln("GRPC Server(TCP) gracefully shutting down...")
			grpcServer.GracefulStop()
		}()
		<-ctx.Done()
	}()

	log.Infoln("GRPC Server(TCP) running...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed starting http server: %v", err)
	}
}
