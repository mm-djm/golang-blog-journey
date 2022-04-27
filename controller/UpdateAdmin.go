package controller

import (
	"encoding/json"
	"golang-blog-journey/model"
	cipher "golang-blog-journey/util/cipher"
	log "golang-blog-journey/util/log"
)

type UpdateAdminRequest struct {
	ApiBaseRequest
	Email    string `json:"email"`
	Password string `json:"user_password"`
	Role     string `json:"role"`
}

func (req *UpdateAdminRequest) Process() ApiBaseResponse {
	log.Infow("UpdateAdmin", "Body", req)
	isExist, err := model.CheckUserNameExist(req.Email, req.RequestID)
	if err != nil {
		log.Errorw("UpdateAdmin", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("UPDATE_ADMIN_ERROR")
	}
	if !isExist {
		log.Errorw("UpdateAdmin", "RequestID", req.RequestID, "err", "user not exist")
		return ErrorResponse("UPDATE_ADMIN_ERROR")
	}

	var pwd string
	if req.Password != "" {
		pwdData := []byte(req.Password)
		pwd, err = cipher.EncryptByAes(pwdData)
		if err != nil {
			log.Errorw("UpdateAdmin", "RequestID", req.RequestID, "err", err)
			return ErrorResponse("UPDATE_ADMIN_ERROR")
		}
	}

	err = model.UpdateAdmin("", pwd, req.Role, req.Email, req.RequestID)
	if err != nil {
		log.Errorw("UpdateAdmin", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("UPDATE_ADMIN_ERROR")
	}
	resp := OkResponse()
	log.Infow("UpdateAdmin", "RequestID", req.RequestID, "response", resp)
	return resp
}

func (req *UpdateAdminRequest) SetReqData(b []byte) {
	if err := json.Unmarshal(b, req); err != nil {
		log.Errorw("json unmarshal", "err", err)
		return
	}
	if req.RequestID == "" {
		req.RequestID = cipher.RandRequestID()
	}
}

func UpdateAdminHandler() BlogIfc {
	return &UpdateAdminRequest{}
}
