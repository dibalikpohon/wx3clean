package handler

import "net/http"

type HandlerStruct struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}
