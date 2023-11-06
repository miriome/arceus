package main

import (
	"database/sql"
	"fmt"
	connect "github.com/arceus/app/auth/gen/protobuf/protobufconnect"
	"github.com/arceus/app/auth/server"
	"github.com/arceus/app/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	user := viper.Get("DB_USER")
	password := viper.Get("DB_PASSWORD")
	dbName := viper.Get("DB_NAME")
	host := viper.Get("DB_HOST")
	port := viper.Get("DB_PORT")

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

	serverPort := viper.Get("SERVER_PORT")

	err = http.ListenAndServeTLS(fmt.Sprintf(":%s", serverPort), certPath, keyPath, middleware.LogRoute(mux))

	if err != nil {
		log.Fatal(err)

	}

}
