package controller

import (
	"errors"
	"golang-blog-journey/model"

	cipher "golang-blog-journey/util/cipher"
	log "golang-blog-journey/util/log"
)

type RegisterRequest struct {
	Name      string `json:"user_name"`
	Password  string `json:"user_password"`
	Email     string `json:"email"`
	RequestID string `json:"RequestID"`
}

func (req *RegisterRequest) Process() ApiBaseResponse {
	if req.RequestID == "" {
		req.RequestID = cipher.RandRequestID()
	}
	log.Infow("Register", "Body", req)
	if req.Email == "" {
		log.Errorw("Register", "RequestID", req.RequestID, "err", errors.New("need email"))
		return ErrorResponse("ADD_ADMIN_ERROR")
	}
	isExist, err := model.CheckUserNameExist(req.Email, req.RequestID)
	if err != nil {
		log.Errorw("Register", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("ADD_ADMIN_ERROR")
	}
	if isExist {
		log.Errorw("Register", "RequestID", req.RequestID, "err", "user is exist")
		return ErrorResponse("ADD_ADMIN_ERROR")
	}
	pwdData := []byte(req.Password)
	pwd, err := cipher.EncryptByAes(pwdData)
	if err != nil {
		log.Errorw("Register", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("ADD_ADMIN_ERROR")
	}

	err = model.AddAdmin(req.Name, pwd, req.Email, req.RequestID)
	if err != nil {
		log.Errorw("Register", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("ADD_ADMIN_ERROR")
	}
	resp := OkResponse()
	log.Infow("Register", "RequestID", req.RequestID, "response", resp)
	return resp
}
