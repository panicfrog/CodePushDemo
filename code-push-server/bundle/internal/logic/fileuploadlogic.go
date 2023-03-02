package logic

import (
	"context"

	"code-push-server/bundle/internal/svc"
	"code-push-server/bundle/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileResponse, err error) {
	// todo: add your logic here and delete this line
	return &types.FileResponse{
		Identity: req.Hash,
		Size:     req.Size,
	}, nil
}
