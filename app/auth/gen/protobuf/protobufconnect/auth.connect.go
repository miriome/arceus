// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: protobuf/auth.proto

package protobufconnect

import (
	protobuf "auth/gen/protobuf"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// AuthServiceName is the fully-qualified name of the AuthService service.
	AuthServiceName = "protobuf.AuthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthServiceLoginProcedure is the fully-qualified name of the AuthService's Login RPC.
	AuthServiceLoginProcedure = "/protobuf.AuthService/Login"
)

// AuthServiceClient is a client for the protobuf.AuthService service.
type AuthServiceClient interface {
	Login(context.Context, *connect.Request[protobuf.LoginRequest]) (*connect.Response[protobuf.LoginResponse], error)
}

// NewAuthServiceClient constructs a client for the protobuf.AuthService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AuthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authServiceClient{
		login: connect.NewClient[protobuf.LoginRequest, protobuf.LoginResponse](
			httpClient,
			baseURL+AuthServiceLoginProcedure,
			opts...,
		),
	}
}

// authServiceClient implements AuthServiceClient.
type authServiceClient struct {
	login *connect.Client[protobuf.LoginRequest, protobuf.LoginResponse]
}

// Login calls protobuf.AuthService.Login.
func (c *authServiceClient) Login(ctx context.Context, req *connect.Request[protobuf.LoginRequest]) (*connect.Response[protobuf.LoginResponse], error) {
	return c.login.CallUnary(ctx, req)
}

// AuthServiceHandler is an implementation of the protobuf.AuthService service.
type AuthServiceHandler interface {
	Login(context.Context, *connect.Request[protobuf.LoginRequest]) (*connect.Response[protobuf.LoginResponse], error)
}

// NewAuthServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthServiceHandler(svc AuthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	authServiceLoginHandler := connect.NewUnaryHandler(
		AuthServiceLoginProcedure,
		svc.Login,
		opts...,
	)
	return "/protobuf.AuthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthServiceLoginProcedure:
			authServiceLoginHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthServiceHandler struct{}

func (UnimplementedAuthServiceHandler) Login(context.Context, *connect.Request[protobuf.LoginRequest]) (*connect.Response[protobuf.LoginResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("protobuf.AuthService.Login is not implemented"))
}
