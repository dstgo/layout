// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.24.4
// source: api/layout/app/application_app_service.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAppCreateApp = "/layout.api.layout.app.v1.App/CreateApp"
const OperationAppDeleteApp = "/layout.api.layout.app.v1.App/DeleteApp"
const OperationAppGetApp = "/layout.api.layout.app.v1.App/GetApp"
const OperationAppListApp = "/layout.api.layout.app.v1.App/ListApp"
const OperationAppUpdateApp = "/layout.api.layout.app.v1.App/UpdateApp"
const OperationAppUpdateAppStatus = "/layout.api.layout.app.v1.App/UpdateAppStatus"

type AppHTTPServer interface {
	// CreateApp CreateApp 创建应用信息
	CreateApp(context.Context, *CreateAppRequest) (*CreateAppReply, error)
	// DeleteApp DeleteApp 删除应用信息
	DeleteApp(context.Context, *DeleteAppRequest) (*DeleteAppReply, error)
	// GetApp GetApp 获取指定的应用信息
	GetApp(context.Context, *GetAppRequest) (*GetAppReply, error)
	// ListApp ListApp 获取应用信息列表
	ListApp(context.Context, *ListAppRequest) (*ListAppReply, error)
	// UpdateApp UpdateApp 更新应用信息
	UpdateApp(context.Context, *UpdateAppRequest) (*UpdateAppReply, error)
	// UpdateAppStatus UpdateAppStatus 更新应用信息状态
	UpdateAppStatus(context.Context, *UpdateAppStatusRequest) (*UpdateAppStatusReply, error)
}

func RegisterAppHTTPServer(s *http.Server, srv AppHTTPServer) {
	r := s.Route("/")
	r.GET("/layout/client/v1/app", _App_GetApp0_HTTP_Handler(srv))
	r.GET("/layout/api/v1/app", _App_GetApp1_HTTP_Handler(srv))
	r.GET("/layout/api/v1/apps", _App_ListApp0_HTTP_Handler(srv))
	r.POST("/layout/api/v1/app", _App_CreateApp0_HTTP_Handler(srv))
	r.PUT("/layout/api/v1/app", _App_UpdateApp0_HTTP_Handler(srv))
	r.PUT("/layout/api/v1/app/status", _App_UpdateAppStatus0_HTTP_Handler(srv))
	r.DELETE("/layout/api/v1/app", _App_DeleteApp0_HTTP_Handler(srv))
}

func _App_GetApp0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAppRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppGetApp)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetApp(ctx, req.(*GetAppRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAppReply)
		return ctx.Result(200, reply)
	}
}

func _App_GetApp1_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAppRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppGetApp)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetApp(ctx, req.(*GetAppRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAppReply)
		return ctx.Result(200, reply)
	}
}

func _App_ListApp0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListAppRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppListApp)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListApp(ctx, req.(*ListAppRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAppReply)
		return ctx.Result(200, reply)
	}
}

func _App_CreateApp0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateAppRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppCreateApp)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CreateApp(ctx, req.(*CreateAppRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateAppReply)
		return ctx.Result(200, reply)
	}
}

func _App_UpdateApp0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateAppRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUpdateApp)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateApp(ctx, req.(*UpdateAppRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateAppReply)
		return ctx.Result(200, reply)
	}
}

func _App_UpdateAppStatus0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateAppStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUpdateAppStatus)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateAppStatus(ctx, req.(*UpdateAppStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateAppStatusReply)
		return ctx.Result(200, reply)
	}
}

func _App_DeleteApp0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteAppRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppDeleteApp)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteApp(ctx, req.(*DeleteAppRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteAppReply)
		return ctx.Result(200, reply)
	}
}

type AppHTTPClient interface {
	CreateApp(ctx context.Context, req *CreateAppRequest, opts ...http.CallOption) (rsp *CreateAppReply, err error)
	DeleteApp(ctx context.Context, req *DeleteAppRequest, opts ...http.CallOption) (rsp *DeleteAppReply, err error)
	GetApp(ctx context.Context, req *GetAppRequest, opts ...http.CallOption) (rsp *GetAppReply, err error)
	ListApp(ctx context.Context, req *ListAppRequest, opts ...http.CallOption) (rsp *ListAppReply, err error)
	UpdateApp(ctx context.Context, req *UpdateAppRequest, opts ...http.CallOption) (rsp *UpdateAppReply, err error)
	UpdateAppStatus(ctx context.Context, req *UpdateAppStatusRequest, opts ...http.CallOption) (rsp *UpdateAppStatusReply, err error)
}

type AppHTTPClientImpl struct {
	cc *http.Client
}

func NewAppHTTPClient(client *http.Client) AppHTTPClient {
	return &AppHTTPClientImpl{client}
}

func (c *AppHTTPClientImpl) CreateApp(ctx context.Context, in *CreateAppRequest, opts ...http.CallOption) (*CreateAppReply, error) {
	var out CreateAppReply
	pattern := "/layout/api/v1/app"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppCreateApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) DeleteApp(ctx context.Context, in *DeleteAppRequest, opts ...http.CallOption) (*DeleteAppReply, error) {
	var out DeleteAppReply
	pattern := "/layout/api/v1/app"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppDeleteApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetApp(ctx context.Context, in *GetAppRequest, opts ...http.CallOption) (*GetAppReply, error) {
	var out GetAppReply
	pattern := "/layout/api/v1/app"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppGetApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) ListApp(ctx context.Context, in *ListAppRequest, opts ...http.CallOption) (*ListAppReply, error) {
	var out ListAppReply
	pattern := "/layout/api/v1/apps"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppListApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UpdateApp(ctx context.Context, in *UpdateAppRequest, opts ...http.CallOption) (*UpdateAppReply, error) {
	var out UpdateAppReply
	pattern := "/layout/api/v1/app"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppUpdateApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UpdateAppStatus(ctx context.Context, in *UpdateAppStatusRequest, opts ...http.CallOption) (*UpdateAppStatusReply, error) {
	var out UpdateAppStatusReply
	pattern := "/layout/api/v1/app/status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppUpdateAppStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}