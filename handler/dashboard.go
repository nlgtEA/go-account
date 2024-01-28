package handler

import (
	"net/http"

	"github.com/nlgtEA/go-account/db/database"
	"github.com/nlgtEA/go-account/view/dashboard"
)

type DashboardHandler struct {
	Queries *database.Queries
}

func (dh *DashboardHandler) HandleShow(w http.ResponseWriter, r *http.Request) {
	dashboard.Show(dashboard.DashboardProps{UserEmail: "thuan.ngo@gmail.com"}).Render(r.Context(), w)
}
