package routes

import (
	"net/http"

	"github.com/TrickingApi/trickingapi/models"
	"github.com/gin-gonic/gin"
)

// GetAllLandingStances godoc
// @Description returns an object with all landing stances
// @Summary Get All Landing Stances
// @Tags landingstances
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]models.LandingStance
// @Router /landingstances [get]
func GetAllLandingStancesHandler(landingstances map[string]models.LandingStance) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		result := make([]models.LandingStance, 0, len(landingstances))

		for _, value := range landingstances {
			result = append(result, value)
		}
		c.IndentedJSON(http.StatusOK, result)
	}
	return gin.HandlerFunc(fn)
}

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

// GetLandingStanceById godoc
// @Description returns a landing stance for a given id
// @Summary Get Landing Stance by id
// @Tags landingstances
// @Accept */*
// @Produce json
// @Success 200 {object} models.LandingStance
// @Failure 404 {object} models.TrickError
// @Router /landingstances/:id [get]
func GetLandingStanceByIdHandler(landingstances map[string]models.LandingStance) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		if landingstance, ok := landingstances[id]; ok {
			c.IndentedJSON(http.StatusOK, landingstance)
		} else {
			c.IndentedJSON(http.StatusNotFound, models.TrickError{Message: "Landing stance not found"})
		}
	}
	return gin.HandlerFunc(fn)
}
