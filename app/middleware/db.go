package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "winson"
	dbName   = "miromie-local"
)

func injectDb(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		connectString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

		db, err := sql.Open("mysql", connectString)
		if err != nil {
			panic(err)
			return
		}

		key := "8OLkOSCfc31EtZQBBpvJnVwjyPJ702nI"
		authHeader := request.Header.Get("Authorization")
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) < 2 {
			http.Error(writer, "Token not found", http.StatusUnauthorized)
			return
		}
		tokenString := splitToken[1]

		token, error := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { return []byte(key), nil })
		if error != nil {
			http.Error(writer, fmt.Sprintf("%v", error), http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			http.Error(writer, "Invalid token claims.", http.StatusForbidden)
			return
		}

		uid := claims["id"]
		ctx := context.WithValue(request.Context(), "uid", uid)

		next.ServeHTTP(writer, request.WithContext(ctx))
		//reqToken := splitToken[1]

	})
}
