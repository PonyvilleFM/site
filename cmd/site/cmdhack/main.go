package main

import (
	"log"

	"github.com/PonyvilleFM/site/schema"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:31337", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := schema.NewUsersClient(conn)
	_, err = client.Register(context.Background(), &schema.RegisterCall{
		User: &schema.User{
			Username:      "AzureDiamond",
			Email:         "its@only.stars.for.me",
			IsAdmin:       true,
			IsDj:          true,
			TwitterHandle: "@AzureDiamond",
		},
		Password:        "hunter2",
		PasswordConfirm: "hunter2",
	})
	if err != nil {
		log.Fatal(err)
	}
}
