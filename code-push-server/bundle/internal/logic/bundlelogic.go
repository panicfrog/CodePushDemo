package logic

import (
	"context"
	"fmt"

	"code-push-server/bundle/internal/svc"
	"code-push-server/bundle/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BundleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBundleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BundleLogic {
	return &BundleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BundleLogic) Bundle(req *types.Request) (resp *types.Response, err error) {
	return &types.Response{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}, nil
}
