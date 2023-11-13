package logic

import (
	"context"

	"sdp-user/internal/svc"
	"sdp-user/pb/user.proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEntityInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEntityInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEntityInfoLogic {
	return &GetEntityInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户会话信息
func (l *GetEntityInfoLogic) GetEntityInfo(in *user_proto.GetEntityInfoReq) (*user_proto.GetEntityInfoResp, error) {
	// todo: add your logic here and delete this line

	return &user_proto.GetEntityInfoResp{}, nil
}
