# syntax=docker/dockerfile:1

FROM golang:1.21.3 as build

WORKDIR /auth

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./


RUN CGO_ENABLED=0 GOOS=linux go build -o /miromie-apiservice-auth

FROM gcr.io/distroless/base-debian11 as release

WORKDIR /

COPY --from=build /miromie-apiservice-auth /miromie-apiservice-auth
COPY --from=build /auth/localhost.key /localhost.key
COPY --from=build /auth/localhost.cert /localhost.cert

EXPOSE 8081

#USER nonroot:nonroot

CMD ["/miromie-apiservice-auth"]
