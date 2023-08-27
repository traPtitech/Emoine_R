// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: emoine_r/v1/admin_api.proto

package emoine_rv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// AdminAPIServiceName is the fully-qualified name of the AdminAPIService service.
	AdminAPIServiceName = "emoine_r.v1.AdminAPIService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AdminAPIServiceCreateEventProcedure is the fully-qualified name of the AdminAPIService's
	// CreateEvent RPC.
	AdminAPIServiceCreateEventProcedure = "/emoine_r.v1.AdminAPIService/CreateEvent"
	// AdminAPIServiceUpdateEventProcedure is the fully-qualified name of the AdminAPIService's
	// UpdateEvent RPC.
	AdminAPIServiceUpdateEventProcedure = "/emoine_r.v1.AdminAPIService/UpdateEvent"
	// AdminAPIServiceDeleteEventProcedure is the fully-qualified name of the AdminAPIService's
	// DeleteEvent RPC.
	AdminAPIServiceDeleteEventProcedure = "/emoine_r.v1.AdminAPIService/DeleteEvent"
	// AdminAPIServiceGetTokensProcedure is the fully-qualified name of the AdminAPIService's GetTokens
	// RPC.
	AdminAPIServiceGetTokensProcedure = "/emoine_r.v1.AdminAPIService/GetTokens"
	// AdminAPIServiceGenerateTokenProcedure is the fully-qualified name of the AdminAPIService's
	// GenerateToken RPC.
	AdminAPIServiceGenerateTokenProcedure = "/emoine_r.v1.AdminAPIService/GenerateToken"
	// AdminAPIServiceRevokeTokenProcedure is the fully-qualified name of the AdminAPIService's
	// RevokeToken RPC.
	AdminAPIServiceRevokeTokenProcedure = "/emoine_r.v1.AdminAPIService/RevokeToken"
)

// AdminAPIServiceClient is a client for the emoine_r.v1.AdminAPIService service.
type AdminAPIServiceClient interface {
	// イベントを作成します
	CreateEvent(context.Context, *connect_go.Request[v1.CreateEventRequest]) (*connect_go.Response[v1.CreateEventResponse], error)
	// イベント情報を更新します
	UpdateEvent(context.Context, *connect_go.Request[v1.UpdateEventRequest]) (*connect_go.Response[emptypb.Empty], error)
	// イベントを削除します
	DeleteEvent(context.Context, *connect_go.Request[v1.DeleteEventRequest]) (*connect_go.Response[emptypb.Empty], error)
	// 該当するイベントのトークン一覧を取得します
	GetTokens(context.Context, *connect_go.Request[v1.GetTokensRequest]) (*connect_go.Response[v1.GetTokensResponse], error)
	// イベント用のトークンを生成します
	GenerateToken(context.Context, *connect_go.Request[v1.GenerateTokenRequest]) (*connect_go.Response[v1.GenerateTokenResponse], error)
	// イベント用のトークンを無効化します
	RevokeToken(context.Context, *connect_go.Request[v1.RevokeTokenRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewAdminAPIServiceClient constructs a client for the emoine_r.v1.AdminAPIService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAdminAPIServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AdminAPIServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &adminAPIServiceClient{
		createEvent: connect_go.NewClient[v1.CreateEventRequest, v1.CreateEventResponse](
			httpClient,
			baseURL+AdminAPIServiceCreateEventProcedure,
			opts...,
		),
		updateEvent: connect_go.NewClient[v1.UpdateEventRequest, emptypb.Empty](
			httpClient,
			baseURL+AdminAPIServiceUpdateEventProcedure,
			opts...,
		),
		deleteEvent: connect_go.NewClient[v1.DeleteEventRequest, emptypb.Empty](
			httpClient,
			baseURL+AdminAPIServiceDeleteEventProcedure,
			opts...,
		),
		getTokens: connect_go.NewClient[v1.GetTokensRequest, v1.GetTokensResponse](
			httpClient,
			baseURL+AdminAPIServiceGetTokensProcedure,
			opts...,
		),
		generateToken: connect_go.NewClient[v1.GenerateTokenRequest, v1.GenerateTokenResponse](
			httpClient,
			baseURL+AdminAPIServiceGenerateTokenProcedure,
			opts...,
		),
		revokeToken: connect_go.NewClient[v1.RevokeTokenRequest, emptypb.Empty](
			httpClient,
			baseURL+AdminAPIServiceRevokeTokenProcedure,
			opts...,
		),
	}
}

// adminAPIServiceClient implements AdminAPIServiceClient.
type adminAPIServiceClient struct {
	createEvent   *connect_go.Client[v1.CreateEventRequest, v1.CreateEventResponse]
	updateEvent   *connect_go.Client[v1.UpdateEventRequest, emptypb.Empty]
	deleteEvent   *connect_go.Client[v1.DeleteEventRequest, emptypb.Empty]
	getTokens     *connect_go.Client[v1.GetTokensRequest, v1.GetTokensResponse]
	generateToken *connect_go.Client[v1.GenerateTokenRequest, v1.GenerateTokenResponse]
	revokeToken   *connect_go.Client[v1.RevokeTokenRequest, emptypb.Empty]
}

// CreateEvent calls emoine_r.v1.AdminAPIService.CreateEvent.
func (c *adminAPIServiceClient) CreateEvent(ctx context.Context, req *connect_go.Request[v1.CreateEventRequest]) (*connect_go.Response[v1.CreateEventResponse], error) {
	return c.createEvent.CallUnary(ctx, req)
}

// UpdateEvent calls emoine_r.v1.AdminAPIService.UpdateEvent.
func (c *adminAPIServiceClient) UpdateEvent(ctx context.Context, req *connect_go.Request[v1.UpdateEventRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.updateEvent.CallUnary(ctx, req)
}

// DeleteEvent calls emoine_r.v1.AdminAPIService.DeleteEvent.
func (c *adminAPIServiceClient) DeleteEvent(ctx context.Context, req *connect_go.Request[v1.DeleteEventRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.deleteEvent.CallUnary(ctx, req)
}

// GetTokens calls emoine_r.v1.AdminAPIService.GetTokens.
func (c *adminAPIServiceClient) GetTokens(ctx context.Context, req *connect_go.Request[v1.GetTokensRequest]) (*connect_go.Response[v1.GetTokensResponse], error) {
	return c.getTokens.CallUnary(ctx, req)
}

// GenerateToken calls emoine_r.v1.AdminAPIService.GenerateToken.
func (c *adminAPIServiceClient) GenerateToken(ctx context.Context, req *connect_go.Request[v1.GenerateTokenRequest]) (*connect_go.Response[v1.GenerateTokenResponse], error) {
	return c.generateToken.CallUnary(ctx, req)
}

// RevokeToken calls emoine_r.v1.AdminAPIService.RevokeToken.
func (c *adminAPIServiceClient) RevokeToken(ctx context.Context, req *connect_go.Request[v1.RevokeTokenRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.revokeToken.CallUnary(ctx, req)
}

// AdminAPIServiceHandler is an implementation of the emoine_r.v1.AdminAPIService service.
type AdminAPIServiceHandler interface {
	// イベントを作成します
	CreateEvent(context.Context, *connect_go.Request[v1.CreateEventRequest]) (*connect_go.Response[v1.CreateEventResponse], error)
	// イベント情報を更新します
	UpdateEvent(context.Context, *connect_go.Request[v1.UpdateEventRequest]) (*connect_go.Response[emptypb.Empty], error)
	// イベントを削除します
	DeleteEvent(context.Context, *connect_go.Request[v1.DeleteEventRequest]) (*connect_go.Response[emptypb.Empty], error)
	// 該当するイベントのトークン一覧を取得します
	GetTokens(context.Context, *connect_go.Request[v1.GetTokensRequest]) (*connect_go.Response[v1.GetTokensResponse], error)
	// イベント用のトークンを生成します
	GenerateToken(context.Context, *connect_go.Request[v1.GenerateTokenRequest]) (*connect_go.Response[v1.GenerateTokenResponse], error)
	// イベント用のトークンを無効化します
	RevokeToken(context.Context, *connect_go.Request[v1.RevokeTokenRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewAdminAPIServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAdminAPIServiceHandler(svc AdminAPIServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	adminAPIServiceCreateEventHandler := connect_go.NewUnaryHandler(
		AdminAPIServiceCreateEventProcedure,
		svc.CreateEvent,
		opts...,
	)
	adminAPIServiceUpdateEventHandler := connect_go.NewUnaryHandler(
		AdminAPIServiceUpdateEventProcedure,
		svc.UpdateEvent,
		opts...,
	)
	adminAPIServiceDeleteEventHandler := connect_go.NewUnaryHandler(
		AdminAPIServiceDeleteEventProcedure,
		svc.DeleteEvent,
		opts...,
	)
	adminAPIServiceGetTokensHandler := connect_go.NewUnaryHandler(
		AdminAPIServiceGetTokensProcedure,
		svc.GetTokens,
		opts...,
	)
	adminAPIServiceGenerateTokenHandler := connect_go.NewUnaryHandler(
		AdminAPIServiceGenerateTokenProcedure,
		svc.GenerateToken,
		opts...,
	)
	adminAPIServiceRevokeTokenHandler := connect_go.NewUnaryHandler(
		AdminAPIServiceRevokeTokenProcedure,
		svc.RevokeToken,
		opts...,
	)
	return "/emoine_r.v1.AdminAPIService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AdminAPIServiceCreateEventProcedure:
			adminAPIServiceCreateEventHandler.ServeHTTP(w, r)
		case AdminAPIServiceUpdateEventProcedure:
			adminAPIServiceUpdateEventHandler.ServeHTTP(w, r)
		case AdminAPIServiceDeleteEventProcedure:
			adminAPIServiceDeleteEventHandler.ServeHTTP(w, r)
		case AdminAPIServiceGetTokensProcedure:
			adminAPIServiceGetTokensHandler.ServeHTTP(w, r)
		case AdminAPIServiceGenerateTokenProcedure:
			adminAPIServiceGenerateTokenHandler.ServeHTTP(w, r)
		case AdminAPIServiceRevokeTokenProcedure:
			adminAPIServiceRevokeTokenHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAdminAPIServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAdminAPIServiceHandler struct{}

func (UnimplementedAdminAPIServiceHandler) CreateEvent(context.Context, *connect_go.Request[v1.CreateEventRequest]) (*connect_go.Response[v1.CreateEventResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("emoine_r.v1.AdminAPIService.CreateEvent is not implemented"))
}

func (UnimplementedAdminAPIServiceHandler) UpdateEvent(context.Context, *connect_go.Request[v1.UpdateEventRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("emoine_r.v1.AdminAPIService.UpdateEvent is not implemented"))
}

func (UnimplementedAdminAPIServiceHandler) DeleteEvent(context.Context, *connect_go.Request[v1.DeleteEventRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("emoine_r.v1.AdminAPIService.DeleteEvent is not implemented"))
}

func (UnimplementedAdminAPIServiceHandler) GetTokens(context.Context, *connect_go.Request[v1.GetTokensRequest]) (*connect_go.Response[v1.GetTokensResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("emoine_r.v1.AdminAPIService.GetTokens is not implemented"))
}

func (UnimplementedAdminAPIServiceHandler) GenerateToken(context.Context, *connect_go.Request[v1.GenerateTokenRequest]) (*connect_go.Response[v1.GenerateTokenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("emoine_r.v1.AdminAPIService.GenerateToken is not implemented"))
}

func (UnimplementedAdminAPIServiceHandler) RevokeToken(context.Context, *connect_go.Request[v1.RevokeTokenRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("emoine_r.v1.AdminAPIService.RevokeToken is not implemented"))
}
