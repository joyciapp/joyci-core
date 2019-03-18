package api

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/joyciapp/joyci-core/cmd/bash"
	"github.com/joyciapp/joyci-core/cmd/git"
	pb "github.com/joyciapp/joyci-core/grpc/proto"
	"google.golang.org/grpc"
)

var (
	pwd, _    = os.Getwd()
	workDir   = "/tmp/build/"
	volumeDir = pwd + workDir
)

const port = ":50051"

// Server structs representing the GRPC Api server
type Server struct{}

// GitClone implementation
func (s *Server) GitClone(ctx context.Context, request *pb.GitCloneRequest) (*empty.Empty, error) {
	git := git.New().VolumeDir(volumeDir).Build()
	git.Clone(request.Repository)

	return new(empty.Empty), nil
}

// ExecuteCommands implementation
func (s *Server) ExecuteCommands(ctx context.Context, request *pb.ExecuteCommandsRequest) (*empty.Empty, error) {
	appName := "joyci-core" // TODO: explore how to pass app name into the context
	bash := bash.New().VolumeDir(volumeDir + "/" + appName).WorkDir(workDir + "/" + appName).Build()
	bash.Execute(request.Commands...)

	return new(empty.Empty), nil
}

// Serve start grpc server
func Serve() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterJoyciCoreServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Joyci Core GRPC server started at ", port)
}
