package server

import (
	"context"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/vldolganov/go_second_microservice/proto"
)

const (
	op = "server.Start"
)

type notesServer struct {
	pb.NotesServer
}

func Start(log *logrus.Logger, port string){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed TCP connect: %s %s", op, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotesServer(grpcServer, &notesServer{})

	log.Infof("Server success starting: %s", op)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to grpc server connect: %s %s", op, err)
	}
}

func(s *notesServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{
		Text: req.Text + "пришел во второй",
	}, nil
}
