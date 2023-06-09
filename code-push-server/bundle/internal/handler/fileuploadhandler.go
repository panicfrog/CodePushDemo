package handler

import (
	"context"
	"fmt"
	"net/http"

	"code-push-server/bundle/internal/logic"
	"code-push-server/bundle/internal/svc"
	"code-push-server/bundle/internal/types"

	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		bucket := "appjsbundle"
		putPolicy := storage.PutPolicy{
			Scope:      bucket,
			FsizeLimit: svcCtx.Config.MaxBytes,
		}
		mac := qbox.NewMac(svcCtx.Config.AccessKey, svcCtx.Config.SecretKey)
		upToken := putPolicy.UploadToken(mac)
		cfg := storage.Config{}
		cfg.Region = &storage.ZoneHuanan
		cfg.UseHTTPS = true
		cfg.UseCdnDomains = false

		formUploader := storage.NewFormUploader(&cfg)
		ret := storage.PutRet{}
		putExtra := storage.PutExtra{}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		defer file.Close()
		// 生成唯一文件名
		uuid := uuid.New().String()
		key := req.Key
		if len(key) <= 0 {
			key = fmt.Sprintf("%s_%s", uuid, fileHeader.Filename)
		}
		err = formUploader.Put(context.Background(), &ret, upToken, key, file, fileHeader.Size, &putExtra)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		req.Hash = ret.Hash
		req.Key = ret.Key
		req.PersistentID = ret.PersistentID
		req.Size = fileHeader.Size

		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
