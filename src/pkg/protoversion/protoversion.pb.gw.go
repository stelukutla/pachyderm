// Code generated by protoc-gen-grpc-gateway
// source: pkg/protoversion/protoversion.proto
// DO NOT EDIT!

/*
Package protoversion is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package protoversion

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/gengo/grpc-gateway/utilities"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"go.pedge.io/google-protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var _ codes.Code
var _ io.Reader
var _ = runtime.String
var _ = json.Marshal
var _ = utilities.PascalFromSnake

func request_Api_GetVersion_0(ctx context.Context, client ApiClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq google_protobuf.Empty

	return client.GetVersion(ctx, &protoReq)
}

// RegisterApiHandlerFromEndpoint is same as RegisterApiHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterApiHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string) (err error) {
	conn, err := grpc.Dial(endpoint)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				glog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				glog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterApiHandler(ctx, mux, conn)
}

// RegisterApiHandler registers the http handlers for service Api to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterApiHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewApiClient(conn)

	mux.Handle("GET", pattern_Api_GetVersion_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_Api_GetVersion_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, err)
			return
		}

		forward_Api_GetVersion_0(ctx, w, req, resp)

	})

	return nil
}

var (
	pattern_Api_GetVersion_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"version"}, ""))
)

var (
	forward_Api_GetVersion_0 = runtime.ForwardResponseMessage
)