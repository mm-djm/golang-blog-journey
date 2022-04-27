package controller

import (
	"encoding/json"
	"golang-blog-journey/model"
	rand "golang-blog-journey/util/cipher"
	log "golang-blog-journey/util/log"
	"strings"
	"time"
)

type DescribeBlogRequest struct {
	ApiBaseRequest
}

func (req *DescribeBlogRequest) Process() ApiBaseResponse {
	log.Infow("DescribeBlog", "Body", req)
	data, err := model.GetBlogData(req.RequestID)
	if err != nil {
		log.Errorw("DescribeBlog", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("GET_BLOG_DATA_ERROR")
	}

	resp := OkResponse()
	resp["dataSet"] = buildBlogData(data)
	log.Infow("DescribeBlog", "RequestID", req.RequestID, "response", resp)
	return resp
}

func buildBlogData(data []*model.Blog) []map[string]interface{} {
	dataSet := make([]map[string]interface{}, 0)
	for _, v := range data {
		dataSet = append(dataSet, map[string]interface{}{
			"id":       v.Id,
			"name":     v.Name,
			"tag":      strings.Split(v.Tag, ","),
			"date":     time.Unix(v.Date, 0).Format("2006-01-02"),
			"count":    v.Count,
			"abstract": v.Abstract,
		})
	}
	return dataSet
}

func (req *DescribeBlogRequest) SetReqData(b []byte) {
	if err := json.Unmarshal(b, req); err != nil {
		log.Errorw("json unmarshal", "err", err)
		return
	}
	if req.RequestID == "" {
		req.RequestID = rand.RandRequestID()
	}
}

func DescribeBlogHandler() BlogIfc {
	return &DescribeBlogRequest{}
}
