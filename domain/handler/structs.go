package handler

import "net/http"

type HandlerStruct struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

func (h HandlerStruct) HttpHandlerFunc() (method string, path string, handlerFunc http.HandlerFunc) {
	return h.Method, h.Path, h.HandlerFunc
}
