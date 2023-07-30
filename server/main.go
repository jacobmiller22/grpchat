package server

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/jacobmiller22/grpchat/chatpb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpchatServer struct {
	pb.UnimplementedChatMessageServiceServer
}

func Main() {
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln("Failed to listen")
	}
	fmt.Println("Listening on port 5000")
	server := grpc.NewServer()
	pb.RegisterChatMessageServiceServer(server, &GrpchatServer{})
	fmt.Println("Registered grpc server")
	server.Serve(listener)
}

func (s *GrpchatServer) StreamChatMessage(stream pb.ChatMessageService_StreamChatMessageServer) error {
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return nil
		}

		stream.Send(&pb.ChatMessage{Text: "replying", Tmstmp: timestamppb.Now(), From: "jacob"})

	}
}
