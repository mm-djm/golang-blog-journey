package controller

import (
	"encoding/json"
	"golang-blog-journey/model"
	rand "golang-blog-journey/util/cipher"
	log "golang-blog-journey/util/log"
	jwt "golang-blog-journey/util/middleware"
)

type DeleteAdminRequest struct {
	ApiBaseRequest
	Email string `json:"email"`
}

func (req *DeleteAdminRequest) Process() ApiBaseResponse {
	log.Infow("DeleteAdmin", "Body", req)
	err := model.RemoveAdmin(req.Email, req.RequestID)
	if err != nil {
		log.Errorw("DeleteAdmin", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("REMOVE_ADMIN_ERROR")
	}
	//delete existed verified token
	if _, ok := jwt.JWTMap[req.Email]; ok {
		delete(jwt.JWTMap, req.Email)
	}
	resp := OkResponse()
	log.Infow("DeleteAdmin", "RequestID", req.RequestID, "response", resp)
	return resp
}

func (req *DeleteAdminRequest) SetReqData(b []byte) {
	if err := json.Unmarshal(b, req); err != nil {
		log.Errorw("json unmarshal", "err", err)
		return
	}
	if req.RequestID == "" {
		req.RequestID = rand.RandRequestID()
	}
}

func DeleteAdminHandler() BlogIfc {
	return &DeleteAdminRequest{}
}
