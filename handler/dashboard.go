package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/nlgtEA/go-account/db/database"
	"github.com/nlgtEA/go-account/view/dashboard"
)

type DashboardHandler struct {
	Queries *database.Queries
}

func (dh *DashboardHandler) HandleShow(w http.ResponseWriter, r *http.Request) {
  // Fake id, please forgive me 🤧
	userID, err := uuid.Parse("bf912e47-690a-4922-be27-33e484dfbb5a")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "some thing went wrong: %v", err)
		return
	}

	assets, err := dh.Queries.GetAssetsByUserID(r.Context(), database.GetAssetsByUserIDParams{
		UserID: userID,
		Offset: 0,
		Limit:  10,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "some thing went wrong: %v", err)
		return
	}

	displayProps := dashboard.DashboardProps{UserEmail: "thuan.ngo@gmail.com", Assets: assets}

	dashboard.Show(displayProps).Render(r.Context(), w)
}
