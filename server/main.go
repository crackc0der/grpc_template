package main

import (
	"context"
	"fmt"
	"log"
	"net"
	demo "server/pkg/note/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	demo.UnimplementedPersonServer
}

func (s *server) GetPerson(ctx context.Context, g *demo.PersonRequest) (*demo.PersonMessage, error) {
	if g.Id == 1 {
		return &demo.PersonMessage{
			Id:        1,
			FirstName: "Pavel",
			LastName:  "Natsaev",
			Age:       35,
		}, nil
	} else {
		return nil, fmt.Errorf("Not found")
	}

}

func (s *server) SetPerson(ctx context.Context, m *demo.PersonMessage) (*demo.PersonResponse, error) {
	return &demo.PersonResponse{Id: 7}, nil
}

func main() {
	conn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	reflection.Register(s)
	demo.RegisterPersonServer(s, &server{})
	s.Serve(conn)
}
