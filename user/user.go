// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package user

import (
	"context"

	"sdp-user/pb/user.proto"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DepartmentInfo    = user_proto.DepartmentInfo
	DirInfo           = user_proto.DirInfo
	EntityInfoInput   = user_proto.EntityInfoInput
	EntityInfoOutput  = user_proto.EntityInfoOutput
	GetEntityInfoReq  = user_proto.GetEntityInfoReq
	GetEntityInfoResp = user_proto.GetEntityInfoResp
	RoleInfo          = user_proto.RoleInfo
	UserInfo          = user_proto.UserInfo

	User interface {
		// 查询用户会话信息
		GetEntityInfo(ctx context.Context, in *GetEntityInfoReq, opts ...grpc.CallOption) (*GetEntityInfoResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

// 查询用户会话信息
func (m *defaultUser) GetEntityInfo(ctx context.Context, in *GetEntityInfoReq, opts ...grpc.CallOption) (*GetEntityInfoResp, error) {
	client := user_proto.NewUserClient(m.cli.Conn())
	return client.GetEntityInfo(ctx, in, opts...)
}