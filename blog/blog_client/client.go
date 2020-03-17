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
		Content:  "hello ke-suke",
	}

	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: blog,
	})
	if err != nil {
		log.Fatalf("unexpected err: %v\n", err)
	}
	blogID := createBlogRes.GetBlog().GetId()
	fmt.Printf("Blog has been created: %v\n", createBlogRes)

	fmt.Println("reading the blog")
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: blogID,
	})
	if readBlogErr != nil {
		fmt.Printf("error happened while reading: %v\n", readBlogErr)
	}
	fmt.Printf("Blog was read: %v\n", readBlogRes)

	// update Blog
	fmt.Println("update blog")
	newBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Yuki",
		Title:    "MY EDITED BLOG",
		Content:  "hello hello everybody",
	}

	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
		Blog: newBlog,
	})
	if updateErr != nil {
		fmt.Printf("error happened while updating: %v\n", updateErr)
	}
	fmt.Printf("blog was updated: %v\n", updateRes)

}
