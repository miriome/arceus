module github.com/arceus/app/auth

go 1.21.3

require (
	connectrpc.com/connect v1.12.0
	github.com/go-jet/jet/v2 v2.10.1
	github.com/go-sql-driver/mysql v1.7.1
	github.com/golang-jwt/jwt/v5 v5.0.0
	golang.org/x/crypto v0.14.0
	google.golang.org/protobuf v1.31
	github.com/arceus/app/middleware v0.0.0
)
replace github.com/arceus/app/middleware => ../middleware

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
