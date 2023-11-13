// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"github.com/cjstarcc/sdp-user/internal/logic"
	"github.com/cjstarcc/sdp-user/internal/svc"
	"github.com/cjstarcc/sdp-user/pb/user.proto"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user_proto.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// 查询用户会话信息
func (s *UserServer) GetEntityInfo(ctx context.Context, in *user_proto.GetEntityInfoReq) (*user_proto.GetEntityInfoResp, error) {
	l := logic.NewGetEntityInfoLogic(ctx, s.svcCtx)
	return l.GetEntityInfo(in)
}
