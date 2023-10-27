package server

import (
	"arceus/app/common/gen/protobuf"
	"connectrpc.com/connect"
	"context"
)

type Server struct{}

func (s *Server) Hello(context.Context, *connect.Request[protobuf.HelloRequest]) (*connect.Response[protobuf.HelloResponse], error) {
	res := connect.NewResponse(&protobuf.HelloResponse{Text: "hello"})
	return res, nil
}

//func (s *Server) Hello(ctx context.Context, req *connect.Request[pb.HelloRequest]) (resp *connect.Response[pb.HelloResponse], err error) {
//	res := connect.NewResponse(&pb.HelloResponse{Text: "hello"})
//	return res, nil
//
//}
