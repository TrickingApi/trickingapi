package routes

import (
	"fmt"
	"net/http"

	"github.com/TrickingApi/trickingapi/models"
	"github.com/gin-gonic/gin"
)

// GetAllTricks godoc
// @Description Reads and returns list of tricks from the static tricks.json file at https://github.com/TrickingApi/trickingapi
// @Summary Get All Tricks from TrickingApi/data/tricks
// @Tags tricks
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]models.Trick
// @Router /tricks [get]
func GetAllTricksHandler(idToTrickMap map[string]models.Trick) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, &idToTrickMap)
	}
	return gin.HandlerFunc(fn)
}

// GetAllTrickNames godoc
// @Description Returns the names of all tricks from the static trickNames.json file at https://github.com/TrickingApi/trickingapi
// @Summary Get All Trick Names and their ids from TrickingApi/data/tricks
// @Tags tricks
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]string
// @Router /tricks/names [get]
func GetAllTrickNamesHandler(idToTrickMap map[string]models.Trick) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var idToTrickNames = make(map[string]string)
		for id, trick := range idToTrickMap {
			idToTrickNames[id] = trick.Name
		}
		c.IndentedJSON(http.StatusOK, idToTrickNames)
	}
	return gin.HandlerFunc(fn)
}

// GetTrick godoc
// @Description reads list of known Trick objects and returns trick matching the name param in the request
// @Summary Get Trick by Specific Name from TrickingApi/data/tricks
// @Tags tricks
// @Accept */*
// @Produce json
// @Success 200 {object} models.Trick
// @Failure 404 {object} models.TrickError
// @Router /tricks/:name [get]
func GetTrickHandler(idToTrickMap map[string]models.Trick) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		name := c.Param("name")

		if trick, ok := idToTrickMap[name]; ok {
			c.IndentedJSON(http.StatusOK, trick)
		} else {
			errorString := fmt.Sprintf("Unknown trick: %s", name)
			error := models.TrickError{
				Message: "Oops, looks like you've requested an unknown trick! Do you think this is a mistake? Consider contributing at https://github.com/TrickingApi/trickingapi",
				Success: false,
				Data:    errorString,
			}
			c.IndentedJSON(http.StatusNotFound, error)
		}
	}

	return gin.HandlerFunc(fn)
}
