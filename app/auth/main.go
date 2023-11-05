package main

import (
	"database/sql"
	"fmt"
	connect "github.com/arceus/app/auth/gen/protobuf/protobufconnect"
	"github.com/arceus/app/auth/server"
	"github.com/arceus/app/middleware"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {

	user := "root"
	password := "winson"
	dbName := "miromie-local"
	host := "host.docker.internal"
	port := 3306

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&collation=utf8mb4_unicode_ci", user, password, host, port, dbName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	sv := &server.Server{
		DB: db,
	} // implements Server interface
	mux := http.NewServeMux()

	path, handler := connect.NewAuthServiceHandler(sv)

	mux.Handle(path, handler)
	fmt.Println("Serving 12345 localhost")
	certPath := "./localhost.cert"
	keyPath := "./localhost.key"

	err = http.ListenAndServeTLS(":12345", certPath, keyPath, middleware.LogRoute(mux))

	if err != nil {
		log.Fatal(err)

	}

}
