package main

import (
	"context"
	"fmt"
	"github.com/IkezawaYuki/protobuf-lesson-go/blog/blogpb"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

var collection *mongo.Collection

type server struct {
}

func (s *server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	blogID := req.GetBlogId()
	oid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	data := &blogItem{}
	filter := bson.M{"_id": oid}
	res := collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with speciifed ID: %v", err),
		)
	}
	return &blogpb.ReadBlogResponse{
		Blog: &blogpb.Blog{
			Id:       data.ID.Hex(),
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		},
	}, nil
}

func (s *server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	fmt.Println("create blog request")

	blog := req.GetBlog()

	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}
	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprint("cannot convert to OID"),
		)
	}
	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       oid.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil
}

func (s *server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	fmt.Println("Delete blgo request")
	oid, err := primitive.ObjectIDFromHex(req.GetBlogId())
	if err != nil {
		status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	filter := bson.M{"_id": oid}
	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("cannot delete object in mongodb: %v", err),
		)
	}
	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("cannot found blog in mongodb: %v", err),
		)
	}

	return &blogpb.DeleteBlogResponse{
		BlogId: req.GetBlogId(),
	}, nil
}

func (s *server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	fmt.Println("Update blog request")
	blog := req.GetBlog()
	oid, err := primitive.ObjectIDFromHex(blog.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("cannot parse ID"),
		)
	}

	data := &blogItem{}
	filter := bson.M{"_id": oid}
	res := collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with speciifed ID: %v", err),
		)
	}

	data.AuthorID = blog.GetAuthorId()
	data.Content = blog.GetContent()
	data.Title = blog.GetTitle()

	_, updateErr := collection.ReplaceOne(ctx, filter, data)
	if updateErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("cannot update object in mongodb: %v", updateErr),
		)
	}

	return &blogpb.UpdateBlogResponse{
		Blog: dataToBlogPb(data),
	}, nil
}

func (s *server) ListBlog(req *blogpb.ListBlogRequest, stream blogpb.BlogService_ListBlogServer) error {
	fmt.Println("list blog request")
	cur, err := collection.Find(context.Background(), nil)
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("unknown internal error: %v", err),
		)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		data := &blogItem{}
		cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("error while decoding from mongodb: %v", err))
		}
		stream.Send(&blogpb.ListBlogResponse{
			Blog: dataToBlogPb(data),
		})
	}
	if err := cur.Err(); err != nil {

	}
	return nil
}

func dataToBlogPb(data *blogItem) *blogpb.Blog {
	return &blogpb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Title:    data.Title,
		Content:  data.Content,
	}
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("connecting to mongodb")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("mydb").Collection("blog")

	fmt.Println("blog service started")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	tls := false
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalf("failed to certificates: %v", sslErr)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("stopping the server")
	s.Stop()
	fmt.Println("stopping the listener")
	lis.Close()
	fmt.Println("closing mongodb connection")
	client.Disconnect(context.Background())
	fmt.Println("end of program")
}
