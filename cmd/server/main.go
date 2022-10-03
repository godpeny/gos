package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	deliveryGRPC "github.com/godpeny/gos/delivery/grpc"
	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"github.com/godpeny/gos/ent"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
	"google.golang.org/grpc/grpclog"
)

var (
	enableTls = false
	port      = 8000
)

func main() {
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

	log.Println("db set up properly")

	// set grpc
	altsTC := alts.NewServerCreds(alts.DefaultServerOptions())
	grpcServer := grpc.NewServer(grpc.Creds(altsTC))
	// register multiple services
	deliveryGRPC.RegisterUserServiceServer(grpcServer, deliveryGRPC.NewUserService(entClient))

	wrappedServer := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(origin string) bool {
		// Allow all origins, DO NOT do this in production
		return true
	}))

	log.Println("grpc set up properly")

	// set router
	router := chi.NewRouter()
	router.Use(
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
		newGrpcWebMiddleware(wrappedServer).Handler,
	)

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	if enableTls {
		// implement when tls enabled.
		//
		// if err := httpServer.ListenAndServeTLS(*tlsCertFilePath, *tlsKeyFilePath); err != nil {
		// 	grpclog.Fatalf("failed starting http2 server: %v", err)
		// }
	} else {
		log.Println("server running")
		if err := httpServer.ListenAndServe(); err != nil {
			grpclog.Fatalf("failed starting http server: %v", err)
		}
	}
}

type GrpcWebMiddleware struct {
	*grpcweb.WrappedGrpcServer
}

func (m *GrpcWebMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.IsAcceptableGrpcCorsRequest(r) || m.IsGrpcWebRequest(r) {
			m.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func newGrpcWebMiddleware(grpcWeb *grpcweb.WrappedGrpcServer) *GrpcWebMiddleware {
	return &GrpcWebMiddleware{grpcWeb}
}
