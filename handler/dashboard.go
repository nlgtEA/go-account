package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/nlgtEA/go-account/db/database"
	"github.com/nlgtEA/go-account/util"
	"github.com/nlgtEA/go-account/view/dashboard"
)

type DashboardHandler struct {
	Queries *database.Queries
}

func (dh *DashboardHandler) HandleShow(w http.ResponseWriter, r *http.Request) {
	// Fake id, please forgive me 🤧
	userID, err := uuid.Parse("bf912e47-690a-4922-be27-33e484dfbb5a")
	if err != nil {
		util.ResponseWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	assets, err := dh.Queries.GetAssetsByUserID(r.Context(),
		database.GetAssetsByUserIDParams{
			UserID: userID,
			Offset: 0,
			Limit:  10,
		})

	if err != nil {
		util.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	displayAssets := []dashboard.DisplayAsset{}

    totalValue := 0
	for _, asset := range assets {
        totalValue += int(asset.CurrentValue)
		displayAssets = append(displayAssets, dashboard.DisplayAsset{
			Name:           asset.Name,
			FormattedValue: util.FormatVND(int(asset.CurrentValue)),
		})
	}

    displayAssets = append(displayAssets, dashboard.DisplayAsset{
        Name: "Total",
        FormattedValue: util.FormatVND(totalValue),
    })

	displayProps := dashboard.DashboardProps{
		UserEmail:     "thuan.ngo@gmail.com",
		DisplayAssets: displayAssets,
	}

	dashboard.Show(displayProps).Render(r.Context(), w)
}
