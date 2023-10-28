package main

import (
	connect "arceus/app/common/gen/protobuf/protobufconnect"
	"arceus/app/common/server"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	sv := &server.Server{} // implements Server interface
	mux := http.NewServeMux()
	workFile := os.Getenv("GOWORK")
	workPath := filepath.Dir(workFile)

	path, handler := connect.NewHelloWorldServiceHandler(sv)

	mux.Handle(path, handler)
	http.ListenAndServeTLS(":8080", fmt.Sprintf("%v/localhost.cert", workPath), fmt.Sprintf("%v/localhost.key", workPath), handler)

}
