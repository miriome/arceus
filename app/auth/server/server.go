package server

import (
	"connectrpc.com/connect"
	"context"
	"database/sql"
	"errors"
	"github.com/arceus/app/auth/gen/miromie-local/model"
	"github.com/arceus/app/auth/gen/miromie-local/table"
	"github.com/arceus/app/auth/gen/protobuf"
	. "github.com/go-jet/jet/v2/mysql"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Server struct {
	*sql.DB
}

func (s *Server) Login(ctx context.Context, req *connect.Request[protobuf.LoginRequest]) (*connect.Response[protobuf.LoginResponse], error) {
	key := "8OLkOSCfc31EtZQBBpvJnVwjyPJ702nI"

	stmt := SELECT(table.Users.AllColumns).FROM(table.Users).WHERE(table.Users.Username.EQ(String(req.Msg.Username))).LIMIT(1)
	var dest []model.Users
	err := stmt.Query(s.DB, &dest)

	if err != nil {
		log.Fatal(err)
	}

	if len(dest) == 0 {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("user not found"))
	}
	user := dest[0]
	// Convert PHP bcrypt hash which starts with $2y$ to $2a$ because Go's bcrypt package understands only $2a$ prefix
	if user.Password[:4] == "$2y$" {
		user.Password = "$2a$" + user.Password[4:]
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Msg.Password))
	if err != nil {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("username or password is incorrect"))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
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
