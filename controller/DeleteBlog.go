package controller

import (
	"encoding/json"
	"golang-blog-journey/model"
	rand "golang-blog-journey/util/cipher"
	log "golang-blog-journey/util/log"
	"os"
	"path/filepath"
)

type DeleteBlogRequest struct {
	ApiBaseRequest
	Id string `json:"article_id"`
}

func (req *DeleteBlogRequest) Process() ApiBaseResponse {
	log.Infow("DeleteBlog", "Body", req)
	err := model.RemoveBlog(req.Id, req.RequestID)
	if err != nil {
		log.Errorw("DeleteBlog", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("REMOVE_BLOG_ERROR")
	}
	fp := filepath.Join(ArticlePath, req.Id)
	err = os.RemoveAll(fp)
	if err != nil {
		log.Errorw("DeleteBlog", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("REMOVE_BLOG_ERROR")
	}
	resp := OkResponse()
	log.Infow("DeleteBlog", "RequestID", req.RequestID, "response", resp)
	return resp
}

func (req *DeleteBlogRequest) SetReqData(b []byte) {
	if err := json.Unmarshal(b, req); err != nil {
		log.Errorw("json unmarshal", "err", err)
		return
	}
	if req.RequestID == "" {
		req.RequestID = rand.RandRequestID()
	}
}

func DeleteBlogHandler() BlogIfc {
	return &DeleteBlogRequest{}
}
