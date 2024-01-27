package handler

import (
	"net/http"

	"github.com/nlgtEA/go-account/view/user"
)

type UserHandler struct{}

func (uh *UserHandler) HandlerUserShow(w http.ResponseWriter, r *http.Request) {
	page := user.Hello()
	page.Render(r.Context(), w)
	w.Write([]byte("U ri sarang haess channa"))
}
