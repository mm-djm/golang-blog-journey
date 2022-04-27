package controller

import (
	"encoding/json"
	"golang-blog-journey/model"
	rand "golang-blog-journey/util/cipher"
	log "golang-blog-journey/util/log"
	"strings"
	"time"
)

type DescribeSimpleBlogRequest struct {
	ArticleId string `json:"article_id"`
	ApiBaseRequest
}

func (req *DescribeSimpleBlogRequest) Process() ApiBaseResponse {
	log.Infow("DescribeSimpleBlog", "Body", req)
	data, err := model.GetSimpleBlogData(req.ArticleId, req.RequestID)
	if err != nil {
		log.Errorw("DescribeSimpleBlog", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("GET_BLOG_DATA_ERROR")
	}
	if len(data) == 0 {
		return ErrorResponse("BLOG_ID_NOT_EXIST")
	}
	v := data[0]
	err = model.UpdateReadCount(req.ArticleId, req.RequestID, v.Count)
	if err != nil {
		log.Errorw("DescribeSimpleBlog", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("GET_BLOG_DATA_ERROR")
	}
	resp := OkResponse()
	resp["dataSet"] = buildSimpleBlogData(v)
	log.Infow("DescribeSimpleBlog", "RequestID", req.RequestID, "response", resp)
	return resp
}

func buildSimpleBlogData(data *model.Blog) map[string]interface{} {
	dataSet := map[string]interface{}{
		"id":      data.Id,
		"name":    data.Name,
		"tag":     strings.Split(data.Tag, ","),
		"date":    time.Unix(data.Date, 0).Format("2006-01-02"),
		"count":   data.Count,
		"content": data.Content,
	}
	return dataSet
}

func (req *DescribeSimpleBlogRequest) SetReqData(b []byte) {
	if err := json.Unmarshal(b, req); err != nil {
		log.Errorw("json unmarshal", "err", err)
		return
	}
	if req.RequestID == "" {
		req.RequestID = rand.RandRequestID()
	}
}

func DescribeSimpleBlogHandler() BlogIfc {
	return &DescribeSimpleBlogRequest{}
}
