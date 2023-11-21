package main

import (
	"connectrpc.com/grpchealth"
	"database/sql"
	"fmt"
	connect "github.com/arceus/app/auth/gen/protobuf/protobufconnect"
	"github.com/arceus/app/auth/server"
	"github.com/arceus/app/middleware"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("env")
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
		os.Exit(1)
	}
	user := viper.Get("DB_USER")
	password := viper.Get("DB_PASSWORD")
	dbName := viper.Get("DB_NAME")
	host := viper.Get("DB_HOST")
	port := viper.GetInt("DB_PORT")
	fmt.Printf("Getting db from host: %s, port: %d", host, port)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	sv := &server.Server{
		DB: db,
	} // implements Server interface
	mux := http.NewServeMux()

	checker := grpchealth.NewStaticChecker(
		connect.AuthServiceName,
	)

	mux.Handle(connect.NewAuthServiceHandler(sv))
	mux.Handle(grpchealth.NewHandler(checker))

	serverPort := viper.Get("SERVER_PORT")
	fmt.Printf("Serving %s localhost\n", serverPort)
	err = http.ListenAndServe(
		fmt.Sprintf(":%s", serverPort),
		h2c.NewHandler(middleware.LogRoute(mux), &http2.Server{}),
	)

	if err != nil {
		log.Fatal(err)

	}

}
