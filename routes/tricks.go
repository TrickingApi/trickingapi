package routes

import (
	"fmt"
	"github.com/TrickingApi/trickingapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllTricksHandler(allTricks *[]models.Trick) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, *allTricks)
	}
	return gin.HandlerFunc(fn)
}

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
