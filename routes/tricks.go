package routes

import (
	"fmt"
	"github.com/TrickingApi/trickingapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllTricks godoc
// @Description Reads and returns list of tricks from the static tricks.json file at https://github.com/TrickingApi/trickingapi
// @Summary Get All Tricks from TrickingApi/data/tricks
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Trick
// @Router /tricks [get]
func GetAllTricksHandler(allTricks *[]models.Trick) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, *allTricks)
	}
	return gin.HandlerFunc(fn)
}

// GetAllTrickNames godoc
// @Description Reads and returns the names of all tricks from the static tricks.json file at https://github.com/TrickingApi/trickingapi
// @Summary Get All Trick Names from TrickingApi/data/tricks
// @Tags root, names
// @Accept */*
// @Produce json
// @Success 200 {object} []string
// @Router /tricks/names [get]
func GetAllTrickNamesHandler(allTricks *[]models.Trick) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var trickNames []string
		for _, trick := range *allTricks {
			trickNames = append(trickNames, trick.Name)
		}
		c.IndentedJSON(http.StatusOK, trickNames)
	}
	return gin.HandlerFunc(fn)
}

// GetTrick godoc
// @Description reads list of known Trick objects and returns trick matching the name param in the request
// @Summary Get Trick by Specific Name from TrickingApi/data/tricks
// @Tags tricks, name
// @Accept */*
// @Produce json
// @Success 200 {object} models.Trick
// @Failure 404 {object} models.TrickError
// @Router /tricks:name [get]
func GetTrickHandler(idToTrickMap map[string]models.Trick) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		name := c.Param("name")

		if trick, ok := idToTrickMap[name]; ok {
			c.IndentedJSON(http.StatusOK, trick)
		} else {
			errorString := fmt.Sprintf("Unknown trick: %s", name)
			error := models.TrickError {
				Message: "Oops, looks like you've requested an unknown trick! Do you think this is a mistake? Consider contributing at https://github.com/TrickingApi/trickingapi",
				Success: false,
				Data: errorString,
			}
			c.IndentedJSON(http.StatusNotFound, error)
		}
	}

	return gin.HandlerFunc(fn)
}
