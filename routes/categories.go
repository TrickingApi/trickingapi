package routes

import (
	"fmt"
	"github.com/TrickingApi/trickingapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategoryToTricksHandler(categoriesToTrickSliceMap map[models.TrickCategory][]models.Trick) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		category := models.TrickCategory(c.Param("name"))

		if slice, ok := categoriesToTrickSliceMap[category]; ok {
			c.IndentedJSON(http.StatusOK, slice)
		} else {
			errorString := fmt.Sprintf("Unknown category: %s", category)
			error := models.TrickError {
				Message: "Oops, looks like you've requested an unknown category of tricks! Do you think this is a mistake? Consider contributing at https://github.com/TrickingApi/trickingapi",
				Success: false,
				Data: errorString,
			}
			c.IndentedJSON(http.StatusNotFound, error)			
		}
	}

	return gin.HandlerFunc(fn)
}

func GetTricksByCategoriesHandler(categoriesToTrickSliceMap map[models.TrickCategory][]models.Trick) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, categoriesToTrickSliceMap)
	}
	return gin.HandlerFunc(fn)
}