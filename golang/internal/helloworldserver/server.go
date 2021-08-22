package helloworldserver

import (
	"context"
	pb "dddsample/rpc/helloworld"
)

// Server implements the Haberdasher service
type HelloWorldServer struct{}

func (s *HelloWorldServer) Hello(ctx context.Context, req *pb.HelloReq) (resp *pb.HelloResp, err error) {
	return &pb.HelloResp{Text: "Hello " + req.GetSubject()}, nil
}
