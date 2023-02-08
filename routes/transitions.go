package routes

import (
	"net/http"

	"github.com/TrickingApi/trickingapi/models"
	"github.com/gin-gonic/gin"
)

// GetAllTransitions godoc
// @Description Reads and returns list of transitions from the transitions enum
// @Summary Get all transitions
// @Tags transitions
// @Accept */*
// @Produce json
// @Success 200 {slice} string
// @Router /transitions [get]
func GetAllTransitionsHandler() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, models.Transitions)
	}
	return gin.HandlerFunc(fn)
}
