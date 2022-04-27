package controller

import (
	"encoding/json"
	"golang-blog-journey/model"
	cipher "golang-blog-journey/util/cipher"
	rand "golang-blog-journey/util/cipher"
	log "golang-blog-journey/util/log"
)

type CreateBlogRequest struct {
	ApiBaseRequest
	Name     string   `json:"name"`
	Tag      []string `json:"tag"`
	Abstract string   `json:"abstract"`
	Content  string   `json:"content"`
}

func (req *CreateBlogRequest) Process() ApiBaseResponse {
	resp := OkResponse()
	randId, err := cipher.RandId(8)
	if err != nil {
		log.Errorw("CreateBlog", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("ADD_BLOG_ERROR")
	}

	err = model.AddBlog(req.Name, req.Tag, randId, req.Abstract, req.Content, req.RequestID)
	if err != nil {
		log.Errorw("CreateBlog", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("ADD_BLOG_ERROR")
	}
	resp["article_id"] = randId
	return resp
}

func (req *CreateBlogRequest) SetReqData(b []byte) {
	if err := json.Unmarshal(b, req); err != nil {
		log.Errorw("json unmarshal", "err", err)
		return
	}
	if req.RequestID == "" {
		req.RequestID = rand.RandRequestID()
	}
}

func CreateBlogHandler() BlogIfc {
	return &CreateBlogRequest{}
}
