// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: coordservice/coordservice.proto

/*
Package coordservice is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package coordservice

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

var (
	filter_CoordsService_GetCoords_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_CoordsService_GetCoords_0(ctx context.Context, marshaler runtime.Marshaler, client CoordsServiceClient, req *http.Request, pathParams map[string]string) (CoordsService_GetCoordsClient, runtime.ServerMetadata, error) {
	var protoReq Units
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_CoordsService_GetCoords_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.GetCoords(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterCoordsServiceHandlerServer registers the http handlers for service CoordsService to "mux".
// UnaryRPC     :call CoordsServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterCoordsServiceHandlerFromEndpoint instead.
func RegisterCoordsServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server CoordsServiceServer) error {

	mux.Handle("GET", pattern_CoordsService_GetCoords_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterCoordsServiceHandlerFromEndpoint is same as RegisterCoordsServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCoordsServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterCoordsServiceHandler(ctx, mux, conn)
}

// RegisterCoordsServiceHandler registers the http handlers for service CoordsService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCoordsServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCoordsServiceHandlerClient(ctx, mux, NewCoordsServiceClient(conn))
}

// RegisterCoordsServiceHandlerClient registers the http handlers for service CoordsService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "CoordsServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CoordsServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CoordsServiceClient" to call the correct interceptors.
func RegisterCoordsServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CoordsServiceClient) error {

	mux.Handle("GET", pattern_CoordsService_GetCoords_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/.CoordsService/GetCoords")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CoordsService_GetCoords_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CoordsService_GetCoords_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CoordsService_GetCoords_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"coords"}, ""))
)

var (
	forward_CoordsService_GetCoords_0 = runtime.ForwardResponseStream
)
