package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/TrickingApi/trickingapi/models"
	"github.com/gin-gonic/gin"
)

// GetAllTransitionIds godoc
// @Description Returns list of transition Ids from the transitions enum
// @Summary Get all transition ids
// @Tags transitions
// @Accept */*
// @Produce json
// @Success 200 {slice} string
// @Router /transitions/ids [get]
func GetAllTransitionIdsHandler() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, models.TransitionIds)
	}
	return gin.HandlerFunc(fn)
}

// GetAllTransitions godoc
// @Description Reads and returns list of transitions from the static transitions.json file at https://github.com/TrickingApi/trickingapi
// @Summary Get all transitions
// @Tags transitions
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Transition
// @Router /transitions [get]
func GetAllTransitionsHandler(transitions map[string]models.Transition) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		result := make([]models.Transition, 0, len(transitions))

		for _, value := range transitions {
			result = append(result, value)
		}
		c.IndentedJSON(http.StatusOK, result)
	}
	return gin.HandlerFunc(fn)
}

// GetTransition godoc
// @Description reads list of known Transition objects and returns a function that returns transition matching the id param in the request
// @Summary Get transition by id from TrickingApi/data/transitions
// @Tags transitions
// @Accept */*
// @Produce json
// @Success	200 {object} models.Transition
// @Failure 404 {object} models.TrickError
// @Router /transitions/:id [get]
func GetTransitionHandler(transitions map[string]models.Transition) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		if transition, ok := transitions[strings.ToUpper(id)]; ok {
			c.IndentedJSON(http.StatusOK, transition)
			return
		}
		error := models.TrickError{
			Message: "Transition not found",
			Success: false,
			Data:    fmt.Sprintf("Unknown transition id: %s", id),
		}
		c.IndentedJSON(http.StatusNotFound, error)
	}
	return gin.HandlerFunc(fn)
}
