package main

import (
	"context"
	sumpb "github.com/IkezawaYuki/protobuf-lesson-go/calculator/calculatorpb"
	"google.golang.org/grpc"
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

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err)
	}
	c := grpc.NewServer()
	sumpb.RegisterCalculateServiceServer(c, &server{})
	if err := c.Serve(lis); err != nil {
		panic(err)
	}
}