package controller

import (
	"encoding/json"
	"golang-blog-journey/model"
	rand "golang-blog-journey/util/cipher"
	log "golang-blog-journey/util/log"
)

type UpdateBlogRequest struct {
	ApiBaseRequest
	ArticleId string   `json:"article_id"`
	Name      string   `json:"name"`
	Tag       []string `json:"tag"`
	Abstract  string   `json:"abstract"`
	Content   string   `json:"content"`
}

func (req *UpdateBlogRequest) Process() ApiBaseResponse {
	resp := OkResponse()
	ok, err := model.CheckBlogExist(req.ArticleId, req.RequestID)
	if err != nil {
		log.Errorw("UpdateBlog", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("UPDATE_BLOG_ERROR")
	}
	if !ok {
		return ErrorResponse("BLOG_ID_NOT_EXIST")
	}

	err = model.UpdateBlog(req.Name, req.Tag, req.ArticleId, req.Abstract, req.Content, req.RequestID)
	if err != nil {
		log.Errorw("UpdateBlog", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("UPDATE_BLOG_ERROR")
	}
	return resp
}

func (req *UpdateBlogRequest) SetReqData(b []byte) {
	if err := json.Unmarshal(b, req); err != nil {
		log.Errorw("json unmarshal", "err", err)
		return
	}
	if req.RequestID == "" {
		req.RequestID = rand.RandRequestID()
	}
}

func UpdateBlogHandler() BlogIfc {
	return &UpdateBlogRequest{}
}
