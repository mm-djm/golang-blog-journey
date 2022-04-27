package controller

import (
	"encoding/json"
	"golang-blog-journey/model"
	rand "golang-blog-journey/util/cipher"
	log "golang-blog-journey/util/log"
)

type DescribeAdminRequest struct {
	ApiBaseRequest
}

func (req *DescribeAdminRequest) Process() ApiBaseResponse {
	log.Infow("DescribeAdmin", "Body", req)
	data, err := model.GetAdminData(req.RequestID)
	if err != nil {
		log.Errorw("DescribeAdmin", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("GET_ADMIN_DATA_ERROR")
	}

	resp := OkResponse()
	resp["dataSet"] = buildAdminData(data)
	log.Infow("DescribeAdmin", "RequestID", req.RequestID, "response", resp)
	return resp
}

func buildAdminData(data []*model.Admin) []map[string]interface{} {
	dataSet := make([]map[string]interface{}, 0)
	for _, v := range data {
		dataSet = append(dataSet, map[string]interface{}{
			"email": v.Email,
			"name":  v.Name,
			"role":  v.Role,
		})
	}
	return dataSet
}

func (req *DescribeAdminRequest) SetReqData(b []byte) {
	if err := json.Unmarshal(b, req); err != nil {
		log.Errorw("json unmarshal", "err", err)
		return
	}
	if req.RequestID == "" {
		req.RequestID = rand.RandRequestID()
	}
}

func DescribeAdminHandler() BlogIfc {
	return &DescribeAdminRequest{}
}
