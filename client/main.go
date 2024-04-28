package main

import (
	"context"
	"fmt"
	"log"
	"time"

	demo "client/pkg/note/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := demo.NewPersonClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	get, err := c.GetPerson(ctx, &demo.PersonRequest{Id: 1})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(get)

	set, err := c.SetPerson(ctx, &demo.PersonMessage{
		Id:        2,
		FirstName: "Sanzhar",
		LastName:  "Sanzhar",
		Age:       20,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(set)
}
