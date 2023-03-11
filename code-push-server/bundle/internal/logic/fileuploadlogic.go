package logic

import (
	"context"

	"code-push-server/bundle/internal/model"
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
	_, err = l.svcCtx.FileModel.Insert(l.ctx, &model.File{Key: req.Key, Hash: req.Hash, Size: req.Size})
	if err != nil {
		return nil, err
	}
	return &types.FileResponse{
		Hash: req.Hash,
		Key:  req.Key,
		Size: req.Size,
	}, nil
}
