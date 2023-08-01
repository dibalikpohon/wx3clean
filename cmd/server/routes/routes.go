package routes

import (
	"net/http"
	"wx3clean/cmd/server/provider"
	"wx3clean/domain/handler"
)

func NewHttpServer() (mux *http.ServeMux) {
	mux = http.NewServeMux()

	var handlers []handler.HandlerInterface
	handlers = append(handlers,
		provider.NewDokterHandler())

	return
}
