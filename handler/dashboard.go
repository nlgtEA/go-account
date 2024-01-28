package handler

import (
	"net/http"

	"github.com/nlgtEA/go-account/view/dashboard"
)

type DashboardHandler struct{}

func (dh *DashboardHandler) HandleShow(w http.ResponseWriter, r *http.Request) {
	dashboard.Show().Render(r.Context(), w)
}
