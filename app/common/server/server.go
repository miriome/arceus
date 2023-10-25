package server

import (
	"context"

	pb "github.com/miromie/arceus/app/common/protobuf"
)

type Server struct{}

func (s *Server) Hello(ctx context.Context, req *pb.HelloReq) (resp *pb.HelloResp, err error) {
	return &pb.HelloResp{
		Text: "hello",
	}, nil

}
