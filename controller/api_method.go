package controller

import (
	"sync"

	"golang-blog-journey/util/log"
)

type APIHandler func() BlogIfc

var handlerAdminMap map[string]APIHandler
var handlerAPIMap map[string]APIHandler
var l sync.Mutex
var ErrorMap = make(map[string]map[string]interface{})

type BlogIfc interface {
	Process() ApiBaseResponse
	SetReqData(b []byte)
}

func registerAdminHandler(action string, h APIHandler) {
	l.Lock()
	handlerAdminMap[action] = h
	l.Unlock()
}

func registerAPIHandler(action string, h APIHandler) {
	l.Lock()
	handlerAPIMap[action] = h
	l.Unlock()
}

func RegisterError(name, content string, retcode int) {
	ErrorMap[name] = ApiBaseResponse{
		"Error":   content,
		"RetCode": retcode,
	}
}

func GetHandler(route string, action string) (APIHandler, error) {
	var ok bool
	var h APIHandler
	l.Lock()
	switch route {
	case "/admin":
		h, ok = handlerAdminMap[action]
	case "/api":
		h, ok = handlerAPIMap[action]
	default:
		l.Unlock()
		return nil, nil
	}

	l.Unlock()
	if !ok {
		log.Errorw("missing action")
	}

	return h, nil
}

type ApiBaseRequest struct {
	Action    string `json:"Action"`
	RequestID string `json:"RequestID"`
}

type (
	ApiBaseResponse map[string]interface{}
)

func OkResponse() ApiBaseResponse {
	resp := ApiBaseResponse{
		"RetCode": 0,
	}
	return resp
}

func ErrorResponse(name string) ApiBaseResponse {
	resp := ErrorMap[name]
	return resp
}

func init() {
	handlerAdminMap = map[string]APIHandler{}
	handlerAPIMap = map[string]APIHandler{}
}
