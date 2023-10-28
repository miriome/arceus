package server

import (
	"arceus/app/auth/gen/protobuf"
	"connectrpc.com/connect"
	"context"
	"github.com/golang-jwt/jwt/v5"
)

type Server struct{}

func (s *Server) Login(context.Context, *connect.Request[protobuf.LoginRequest]) (*connect.Response[protobuf.LoginResponse], error) {
	key := "8OLkOSCfc31EtZQBBpvJnVwjyPJ702nI"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       "1",
		"username": "test",
		"email":    "test@e.com",
	})
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	res := connect.NewResponse(&protobuf.LoginResponse{
		Jwt: tokenString,
	})
	return res, nil
}
