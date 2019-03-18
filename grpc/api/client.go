package api

import (
	"context"
	"log"
	"time"

	pb "github.com/joyciapp/joyci-core/grpc/proto"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func connect(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func newClient(conn *grpc.ClientConn) pb.JoyciCoreClient {
	return pb.NewJoyciCoreClient(conn)
}

func newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 15*time.Second)
}

// GitClone clones a git repository
func GitClone(repository string) {
	conn := connect(address)
	defer conn.Close()

	c := newClient(conn)

	ctx, cancel := newContext()
	defer cancel()

	if err := ctx.Err(); err != nil {
		log.Fatal("error context:", err)
	}

	_, err := c.GitClone(ctx, &pb.GitCloneRequest{Repository: repository})
	if err != nil {
		log.Fatal("error on clone a repository:", err)
	}
}

// ExecuteCommands execute bash commands
func ExecuteCommands(commands ...string) {
	conn := connect(address)
	defer conn.Close()

	c := newClient(conn)

	ctx, cancel := newContext()
	defer cancel()

	if err := ctx.Err(); err != nil {
		log.Fatal("error context:", err)
	}

	_, err := c.ExecuteCommands(ctx, &pb.ExecuteCommandsRequest{Commands: commands})
	if err != nil {
		log.Fatal("error on execute commands:", err)
	}
}
