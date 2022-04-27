package controller

import (
	"fmt"
	"golang-blog-journey/model"
	cipher "golang-blog-journey/util/cipher"
	log "golang-blog-journey/util/log"
	jwt "golang-blog-journey/util/middleware"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	RequestID string `json:"RequestID"`
}

func (req *LoginRequest) Process(c *gin.Context) ApiBaseResponse {
	if req.RequestID == "" {
		req.RequestID = cipher.RandRequestID()
	}
	log.Infow("Login", "Body", req)
	data, err := model.GetSimpleAdminData(req.Email, req.RequestID)
	if err != nil {
		log.Errorw("Login", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("LOGIN_ERROR")
	}

	pwd, err := cipher.DecryptByAes(data.Password)
	if err != nil {
		log.Errorw("Login", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("LOGIN_ERROR")
	}
	p := string(pwd)
	if p != req.Password {
		return ErrorResponse("LOGIN_ERROR")
	}

	claims := &jwt.MyClaims{
		Email: req.Email,
	}
	j := jwt.NewJWT()
	token, err := j.CreateToken(*claims)
	if err != nil {
		log.Errorw("CreateToken", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("LOGIN_ERROR")
	}

	jwt.RegisterMap(req.Email, data.Role)
	domain := fmt.Sprintf(".%s", c.Request.Host)
	c.SetCookie("user", token, 60*60*24, "/app/admin", domain, false, true)
	c.SetCookie("user", token, 60*60*24, "/app/api/logout", domain, false, true)
	resp := OkResponse()
	resp["Token"] = token
	log.Infow("Login", "RequestID", req.RequestID, "response", resp)
	return resp
}
