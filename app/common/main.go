package main

import (
	connect "arceus/app/common/gen/protobuf/protobufconnect"
	"arceus/app/common/server"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

func main() {

	sv := &server.Server{} // implements Server interface
	mux := http.NewServeMux()
	path, handler := connect.NewHelloWorldServiceHandler(sv)
	mux.Handle(path, handler)

	http.ListenAndServe(":8080", h2c.NewHandler(mux, &http2.Server{}))
}
