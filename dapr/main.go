package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dapr/go-sdk/service/common"
	"github.com/dapr/go-sdk/service/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	s, err := grpc.NewService(":50001")
	if err != nil {
		log.Fatal(err)
	}

	// Add router
	s.AddServiceInvocationHandler("echo", func(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
		m, _ := metadata.FromIncomingContext(ctx)
		fmt.Println(m)
		fmt.Println(in.ContentType, in.Data, in.QueryString, in.Verb)

		return nil, nil
	})

	log.Fatal(s.Start())
}
