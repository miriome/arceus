// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: protobuf/hello_world.proto

package protobufconnect

import (
	protobuf "arceus/app/common/gen/protobuf"
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
	// HelloWorldServiceName is the fully-qualified name of the HelloWorldService service.
	HelloWorldServiceName = "protobuf.v1.HelloWorldService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// HelloWorldServiceHelloProcedure is the fully-qualified name of the HelloWorldService's Hello RPC.
	HelloWorldServiceHelloProcedure = "/protobuf.v1.HelloWorldService/Hello"
)

// HelloWorldServiceClient is a client for the protobuf.v1.HelloWorldService service.
type HelloWorldServiceClient interface {
	Hello(context.Context, *connect.Request[protobuf.HelloRequest]) (*connect.Response[protobuf.HelloResponse], error)
}

// NewHelloWorldServiceClient constructs a client for the protobuf.v1.HelloWorldService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHelloWorldServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) HelloWorldServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &helloWorldServiceClient{
		hello: connect.NewClient[protobuf.HelloRequest, protobuf.HelloResponse](
			httpClient,
			baseURL+HelloWorldServiceHelloProcedure,
			opts...,
		),
	}
}

// helloWorldServiceClient implements HelloWorldServiceClient.
type helloWorldServiceClient struct {
	hello *connect.Client[protobuf.HelloRequest, protobuf.HelloResponse]
}

// Hello calls protobuf.v1.HelloWorldService.Hello.
func (c *helloWorldServiceClient) Hello(ctx context.Context, req *connect.Request[protobuf.HelloRequest]) (*connect.Response[protobuf.HelloResponse], error) {
	return c.hello.CallUnary(ctx, req)
}

// HelloWorldServiceHandler is an implementation of the protobuf.v1.HelloWorldService service.
type HelloWorldServiceHandler interface {
	Hello(context.Context, *connect.Request[protobuf.HelloRequest]) (*connect.Response[protobuf.HelloResponse], error)
}

// NewHelloWorldServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHelloWorldServiceHandler(svc HelloWorldServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	helloWorldServiceHelloHandler := connect.NewUnaryHandler(
		HelloWorldServiceHelloProcedure,
		svc.Hello,
		opts...,
	)
	return "/protobuf.v1.HelloWorldService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case HelloWorldServiceHelloProcedure:
			helloWorldServiceHelloHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedHelloWorldServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedHelloWorldServiceHandler struct{}

func (UnimplementedHelloWorldServiceHandler) Hello(context.Context, *connect.Request[protobuf.HelloRequest]) (*connect.Response[protobuf.HelloResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("protobuf.v1.HelloWorldService.Hello is not implemented"))
}
