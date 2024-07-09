package handler

import (
	"net/http"

	"github.com/LucasBiazon/Gopportunitie.git/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "failed to get openings")
		return
	}
	sendSuccess(ctx, "show-openings", openings)
}
