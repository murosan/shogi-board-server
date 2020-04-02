package handlers

import (
	"net/http"

	"github.com/murosan/shogi-board-server/app/server/handler"
)

// OKHandler is a handler that just says OK.
type OKHandler struct{}

func NewOKHandler() handler.Handler {
	return &OKHandler{}
}

func (hdr *OKHandler) Func(ctx *handler.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func (*OKHandler) Description() string {
	return "" // TODO
}

func (*OKHandler) Methods() []string {
	return []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
	}
}
