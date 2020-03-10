package main

import (
	"context"
	"fmt"
	sumpb "github.com/IkezawaYuki/protobuf-lesson-go/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := sumpb.NewCalculateServiceClient(conn)
	//doUnary(c)
	//doStream(c)
	doStreamingClient(c)
}

func doStreamingClient(c sumpb.CalculateServiceClient) {
	nums := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		panic(err)
	}
	for _, r := range nums {
		fmt.Printf("sending %v\n", r)
		stream.Send(&sumpb.AverageRequest{
			Num: r,
		})
		time.Sleep(1000 * time.Millisecond)
	}
	result, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	fmt.Println("answer is ", result.GetResult())
}

func doStream(c sumpb.CalculateServiceClient) {
	req := &sumpb.NumRequest{
		Num: 120,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		panic(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(res.GetResult())
	}
}

func doUnary(c sumpb.CalculateServiceClient) {
	req := &sumpb.SumRequest{
		A: 10,
		B: 5,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
