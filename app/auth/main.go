package main

import (
	connect "arceus/app/auth/gen/protobuf/protobufconnect"
	"arceus/app/auth/server"
	"fmt"
	"log"
	"net/http"
)

func main() {

	sv := &server.Server{} // implements Server interface
	mux := http.NewServeMux()

	path, handler := connect.NewAuthServiceHandler(sv)

	mux.Handle(path, handler)
	fmt.Println("Serving 8081 localhost")
	certPath := "../localhost.cert"
	keyPath := "../localhost.key"
	err := http.ListenAndServeTLS("localhost:8081", certPath, keyPath, mux)
	if err != nil {
		log.Fatal(err)
	}

}
