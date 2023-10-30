package main

import (
	connect "arceus/app/auth/gen/protobuf/protobufconnect"
	"arceus/app/auth/server"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {

	user := "root"
	password := "winson"
	dbName := "miromie-local"
	host := "localhost"
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
	fmt.Println("Serving 8081 localhost")
	certPath := "../localhost.cert"
	keyPath := "../localhost.key"

	err = http.ListenAndServeTLS("localhost:8081", certPath, keyPath, mux)

	if err != nil {
		log.Fatal(err)

	}

}
