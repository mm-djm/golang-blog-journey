package controller

import (
	"encoding/json"
	"golang-blog-journey/model"
	rand "golang-blog-journey/util/cipher"
	log "golang-blog-journey/util/log"
)

type DescribeSimpleAdminRequest struct {
	Email string `json:"email"`
	ApiBaseRequest
}

func (req *DescribeSimpleAdminRequest) Process() ApiBaseResponse {
	log.Infow("DescribeSimpleAdmin", "Body", req)
	data, err := model.GetSimpleAdminData(req.Email, req.RequestID)
	if err != nil {
		log.Errorw("DescribeSimpleAdmin", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("GET_ADMIN_DATA_ERROR")
	}

	resp := OkResponse()
	resp["dataSet"] = buildSimpleAdminData(data)
	log.Infow("DescribeSimpleAdmin", "RequestID", req.RequestID, "response", resp)
	return resp
}

func buildSimpleAdminData(data *model.Admin) map[string]interface{} {
	dataSet := map[string]interface{}{
		"email": data.Email,
		"name":  data.Name,
		"role":  data.Role,
	}
	return dataSet
}

func (req *DescribeSimpleAdminRequest) SetReqData(b []byte) {
	if err := json.Unmarshal(b, req); err != nil {
		log.Errorw("json unmarshal", "err", err)
		return
	}
	if req.RequestID == "" {
		req.RequestID = rand.RandRequestID()
	}
}

func DescribeSimpleAdminHandler() BlogIfc {
	return &DescribeSimpleAdminRequest{}
}
