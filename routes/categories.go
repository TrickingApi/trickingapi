package routes

import (
	"fmt"
	"github.com/TrickingApi/trickingapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTricksForCategory godoc
// @Description Reads and returns a list of tricks for a specific category from the tricks.json file at https://github.com/TrickingApi/trickingapi
// @Summary Get All Tricks Grouped Under A Category from TrickingApi/data/tricks
// @Tags root, categories, tricks
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Trick
// @Router /categories/:name [get]
func GetTricksForCategoryHandler(categoriesToTrickSliceMap map[models.TrickCategory][]models.Trick) gin.HandlerFunc {
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

// GetAllTricksByCategories godoc
// @Description Reads and returns a mapping of categories to list of tricks from the tricks.json file at https://github.com/TrickingApi/trickingapi
// @Summary Get All Tricks Grouped by Categories from TrickingApi/data/tricks
// @Tags root, categories, tricks
// @Accept */*
// @Produce json
// @Success 200 {object} map[models.TrickCategory][]models.Trick
// @Router /categories/tricks [get]
func GetAllTricksByCategoriesHandler(categoriesToTrickSliceMap map[models.TrickCategory][]models.Trick) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, categoriesToTrickSliceMap)
	}
	return gin.HandlerFunc(fn)
}

// GetAllCategories godoc
// @Description Reads and returns all existing categories of tricks from the tricks.json file at https://github.com/TrickingApi/trickingapi
// @Sumary Get All Category Names from TrickingApi/data/tricks
// @Tags root, categories
// @Accept */*
// @Produce json
// @Success 200 {object} []string
// @Router /categories [get]
func GetAllCategoriesHandler(categoriesToTrickSliceMap map[models.TrickCategory][]models.Trick) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// Tbh, would be a lot more efficient if there was a categories.json
		keys := make([]models.TrickCategory, len(categoriesToTrickSliceMap))
		i := 0
		for k := range categoriesToTrickSliceMap {
				keys[i] = k
				i++
		}
		c.IndentedJSON(http.StatusOK, keys)
	}
	return gin.HandlerFunc(fn)
}