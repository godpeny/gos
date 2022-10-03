package main

import (
	"context"
	"flag"
	"fmt"
	deliveryGRPC "github.com/godpeny/gos/delivery/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"math/rand"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:8001", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// get client
	client := deliveryGRPC.NewUserServiceClient(conn)

	// Contact the server and create
	ctx := context.Background()
	user := randomUser()
	created, err := client.Create(ctx, &deliveryGRPC.CreateUserRequest{
		User: user,
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed creating user: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("user created with id: %d", created.Id)

	// On a separate RPC invocation, retrieve the user we saved previously.
	get, err := client.Get(ctx, &deliveryGRPC.GetUserRequest{
		Id: created.Id,
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed retrieving user: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("retrieved user with id=%d: %v", get.Id, get)
}

func randomUser() *deliveryGRPC.User {
	return &deliveryGRPC.User{
		Username: fmt.Sprintf("user_%d", rand.Int()),
	}
}
