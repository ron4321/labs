package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"labs-grpc/proto"
	"log"
	"net"
)

//go:generate protoc -I ../proto --go_out=plugins=grpc:../proto ../proto/chat.proto
const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	chatService := NewChatService()
	chat.RegisterChatServer(s, chatService)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type ChatService struct{

}
func NewChatService() *ChatService{
	return &ChatService{}
}
func (s *ChatService) Send(context context.Context, message *chat.ChatMessageRequest) (*chat.ChatMessageResponse, error){
	fmt.Println("Server: received message -> "+ message.Body)
	return &chat.ChatMessageResponse{
		Body:"server received [" + message.Body + "]",
	}, nil
}