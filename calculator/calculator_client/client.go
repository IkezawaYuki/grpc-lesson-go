package main

import (
	"context"
	"fmt"
	sumpb "github.com/IkezawaYuki/protobuf-lesson-go/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
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
	//doStreamingClient(c)
	//doBiDiStreamClient(c)
	doErrorUnary(c)
}

func doErrorUnary(c sumpb.CalculateServiceClient) {
	doErrorCall(c, 10)
	doErrorCall(c, -5)

}

func doErrorCall(c sumpb.CalculateServiceClient, number int32) {
	fmt.Println("Starting to do a SquareRoot Unary RPC...")
	res, err := c.SquareRoot(context.Background(), &sumpb.SquareRootRequest{
		Number: number,
	})

	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
			}
		} else {
			log.Fatalf("Big Error calling SquareRoot: %v", err)
		}
	}
	fmt.Printf("Request of square root of %v: %v", number, res.GetNumberRoot())
}

func doBiDiStreamClient(c sumpb.CalculateServiceClient) {
	nums := []int32{1, 2, 1, 4, 1, 6, 1, 8, 1, 10}
	waitc := make(chan interface{})

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		panic(err)
	}

	go func() {
		for _, n := range nums {
			time.Sleep(1000 * time.Millisecond)
			err := stream.Send(&sumpb.FindMaximumRequest{
				Num: n,
			})
			if err != nil {
				panic(err)
			}
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				break
			}
			if err != nil {
				panic(err)
			}
			result := resp.GetMax()
			fmt.Printf("max is... : %v\n", result)
		}
	}()

	<-waitc
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
