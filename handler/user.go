package handler

import "net/http"

type UserHandler struct{}

func (uh *UserHandler) HandlerUserShow(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("U ri sarang haess channa"))
}
