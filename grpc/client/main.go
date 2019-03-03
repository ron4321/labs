package main

import (
	"context"
	"flag"
	"labs-grpc/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
)
var (
	message = flag.String("m", "default", "send message body")
)
func main() {

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	chatClient := chat.NewChatClient(conn)
	resp,err := chatClient.Send(ctx, &chat.ChatMessageRequest{
		Body:*message,
	})

	if err != nil {
		log.Fatalf("could not chat: %v", err)
	}
	log.Printf("Client: %s", resp.Body)
}