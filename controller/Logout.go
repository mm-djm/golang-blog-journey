package controller

import (
	"golang-blog-journey/util/cipher"
	log "golang-blog-journey/util/log"
	jwt "golang-blog-journey/util/middleware"

	"github.com/gin-gonic/gin"
)

type LogoutRequest struct {
	RequestID string `json:"RequestID"`
}

func (req *LogoutRequest) Process(c *gin.Context) ApiBaseResponse {
	if req.RequestID == "" {
		req.RequestID = cipher.RandRequestID()
	}
	log.Infow("Logout", "Body", req)

	tokenHeader, err := c.Cookie("user")
	if err != nil {
		log.Errorw("Logout", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("LOGOUT_ERROR")
	}
	j := jwt.NewJWT()
	claims, err := j.ParseToken(tokenHeader)
	if err != nil {
		log.Errorw("Logout", "RequestID", req.RequestID, "err", err)
		return ErrorResponse("LOGOUT_ERROR")
	}
	delete(jwt.JWTMap, claims.Email)

	resp := OkResponse()
	log.Infow("Logout", "RequestID", req.RequestID, "response", resp)
	return resp
}
