package handler

import "net/http"

type HandlerInterface interface {
	GetHandlers() []HandlerStruct
	http.Handler
}
