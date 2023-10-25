package main

import (
	"net/http"

	pb "github.com/miromie/arceus/app/common/protobuf"
	"github.com/miromie/arceus/app/common/server"
)

func main() {
	sv := &server.Server{} // implements Server interface
	twirpHandler := pb.NewHelloWorldServer(sv)

	http.ListenAndServe(":8080", twirpHandler)
}
