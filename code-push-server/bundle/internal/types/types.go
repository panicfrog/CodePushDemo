// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me|he|her"`
}

type Response struct {
	Message string `json:"message"`
}

type FileUploadRequest struct {
	Hash         string `json:"hash,optional"`
	Key          string `json:"key,optional"`
	Size         int64  `json:"size,optional"`
	PersistentID string `json:"persistentID,optional"`
}

type FileResponse struct {
	Identity string `json:"identity"`
	Size     int64  `json:"size"`
}
