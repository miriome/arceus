# syntax=docker/dockerfile:1

FROM golang:1.21.3 as build

WORKDIR /auth
RUN mkdir -p auth
COPY auth ./auth

## TODO: Once we have proper git going on, dont need to copy like ths anymore.
RUN mkdir -p middleware
COPY middleware ./middleware

WORKDIR /auth/auth
# copy prod env
COPY auth/prod.toml env.toml



RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /miromie-apiservice-auth

FROM gcr.io/distroless/base-debian11 as release

WORKDIR /

COPY --from=build /miromie-apiservice-auth /miromie-apiservice-auth
COPY --from=build /auth/auth/env.toml /env.toml

EXPOSE 12345

#USER nonroot:nonroot

CMD ["/miromie-apiservice-auth"]
