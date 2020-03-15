package main

import (
	"context"
	"fmt"
	"github.com/IkezawaYuki/protobuf-lesson-go/blog/blogpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	fmt.Println("blog client started")

	opts := grpc.WithInsecure()

	tls := false
	if tls {
		certFile := "ssl/ca.crt"
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("error while loading CA trust certificate: %v", sslErr)
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	fmt.Println("creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Yuki",
		Title:    "MY FIRST BLOG",
		Content:  "こんにちは",
	}

	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: blog,
	})
	if err != nil {
		log.Fatalf("unexpected err: %v", err)
	}
	fmt.Printf("Blog has been created: %v", createBlogRes)
}
