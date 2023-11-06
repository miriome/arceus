# syntax=docker/dockerfile:1

FROM golang:1.21.3 as build

WORKDIR /auth
RUN mkdir -p auth
RUN mkdir -p middleware
## TODO: Once we have proper git going on, dont need to copy like ths anymore.
COPY ./auth ./auth
COPY ./middleware ./middleware
WORKDIR /auth/auth
# copy prod env
ADD auth/.env.prod .env 

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /miromie-apiservice-auth

FROM gcr.io/distroless/base-debian11 as release

WORKDIR /

COPY --from=build /miromie-apiservice-auth /miromie-apiservice-auth
COPY --from=build /auth/auth/localhost.key /localhost.key
COPY --from=build /auth/auth/localhost.cert /localhost.cert

EXPOSE 12345

#USER nonroot:nonroot

CMD ["/miromie-apiservice-auth"]
