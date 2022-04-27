package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"golang-blog-journey/util/log"
)

func Admin(c *gin.Context) {
	var raw = new(ApiBaseRequest)
	bs, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorw("read body", "err", err)
		return
	}

	if err := json.Unmarshal(bs, raw); err != nil {
		log.Errorw("json unmarshal", "err", err)
		return
	}
	handler, err := GetHandler("/admin", raw.Action)
	if err != nil {
		log.Errorw("get handler", "err", err)
		return
	}
	reqObj := handler()
	reqObj.SetReqData(bs)
	json := reqObj.Process()
	c.JSON(http.StatusOK, json)
}

func API(c *gin.Context) {
	var raw = new(ApiBaseRequest)
	bs, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorw("read body", "err", err)
		return
	}

	if err := json.Unmarshal(bs, raw); err != nil {
		log.Errorw("json unmarshal", "err", err)
		return
	}
	handler, err := GetHandler("/api", raw.Action)
	if err != nil {
		log.Errorw("get handler", "err", err)
		return
	}
	reqObj := handler()
	reqObj.SetReqData(bs)
	json := reqObj.Process()
	c.JSON(http.StatusOK, json)
}

func LoginAdmin(c *gin.Context) {
	var raw = new(LoginRequest)
	bs, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorw("read body", "err", err)
		return
	}

	if err := json.Unmarshal(bs, raw); err != nil {
		log.Errorw("json unmarshal", "err", err)
		return
	}
	json := raw.Process(c)
	c.JSON(http.StatusOK, json)
}

func RegisterAdmin(c *gin.Context) {
	var raw = new(RegisterRequest)
	bs, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorw("read body", "err", err)
		return
	}

	if err := json.Unmarshal(bs, raw); err != nil {
		log.Errorw("json unmarshal", "err", err)
		return
	}
	json := raw.Process()
	c.JSON(http.StatusOK, json)
}

func LogoutAdmin(c *gin.Context) {
	var raw = new(LogoutRequest)
	bs, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorw("read body", "err", err)
		return
	}

	if err := json.Unmarshal(bs, raw); err != nil {
		log.Errorw("json unmarshal", "err", err)
		return
	}
	json := raw.Process(c)
	c.JSON(http.StatusOK, json)
}
