package middleware

import (
	"fmt"
	"net/http"
)

func LogRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("%s\n", request.URL)

		next.ServeHTTP(writer, request)
		//reqToken := splitToken[1]

	})
}
