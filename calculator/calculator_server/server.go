package main

import (
	"context"
	"fmt"
	sumpb "github.com/IkezawaYuki/protobuf-lesson-go/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"io"
	"math"
	"net"
	"time"
)

type server struct {
}

func (s *server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	result := req.A + req.B
	res := &sumpb.SumResponse{
		Result: result,
	}
	return res, nil
}

func (s *server) PrimeNumberDecomposition(req *sumpb.NumRequest, stream sumpb.CalculateService_PrimeNumberDecompositionServer) error {
	N := int(req.GetNum())
	k := 2
	for N > 1 {
		if N%k == 0 {
			resp := sumpb.NumResponse{
				Result: int32(k),
			}
			stream.Send(&resp)
			time.Sleep(1000 * time.Millisecond)
			N = N / k
		} else {
			k = k + 1
		}
	}
	return nil
}

func (s *server) ComputeAverage(stream sumpb.CalculateService_ComputeAverageServer) error {
	var sum float32 = 0
	var n float32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			result := sum / n
			return stream.SendAndClose(&sumpb.AverageResponse{
				Result: result,
			})
		}
		if err != nil {
			panic(err)
		}
		sum += float32(req.GetNum())
		n = n + 1
	}
}

func (s *server) FindMaximum(stream sumpb.CalculateService_FindMaximumServer) error {
	number := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		num := int(req.GetNum())
		if number < num {
			number = num
		}
		err = stream.Send(&sumpb.FindMaximumResponse{
			Max: int32(number),
		})
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (s *server) SquareRoot(ctx context.Context, req *sumpb.SquareRootRequest) (*sumpb.SquareRootResponse, error) {
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}
	return &sumpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	sumpb.RegisterCalculateServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
