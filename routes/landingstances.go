package routes

import (
	"net/http"

	"github.com/TrickingApi/trickingapi/models"
	"github.com/gin-gonic/gin"
)

// GetAllLandingStanceIds godoc
// @Description returns a list of all landing stances ids
// @Summary Get All Landing Stances ids
// @Tags landingstances
// @Accept */*
// @Produce json
// @Success 200 {object} []models.LandingStance
// @Router /landingstances/ids [get]

func GetAllLandingStanceIdsHandler() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, models.LandingStances)
	}
	return gin.HandlerFunc(fn)
}
