package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/TrickingApi/trickingapi/models"
	"github.com/TrickingApi/trickingapi/routes"
	"github.com/TrickingApi/trickingapi/utils"
	"github.com/gin-gonic/gin"
)

var allTricks []models.Trick
var idToTrickMap map[string]models.Trick
var categoriesToTrickSliceMap map[models.TrickCategory][]models.Trick

// organizes tricks by id and category upon server init
func init() {
	content, err := os.ReadFile("data/tricks.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	if err := json.Unmarshal(content, &allTricks); err != nil {
		panic(err)
	}

	idToTrickMap = make(map[string]models.Trick)

	categoriesToTrickSliceMap = make(map[models.TrickCategory][]models.Trick)

	for _, category := range models.Categories {
		emptySlice := []models.Trick{}
		categoriesToTrickSliceMap[category] = emptySlice
	}

	for _, trick := range allTricks {
		idToTrickMap[trick.Id] = trick

		var categories = trick.Categories

		for _, category := range categories {
			tc := models.TrickCategory(category)

			if tricks, ok := categoriesToTrickSliceMap[tc]; ok {
				tricks = append(tricks, trick)
				categoriesToTrickSliceMap[tc] = tricks
			}
		}
	}
}

// @title Tricking Api
// @version 0.1
// @description Consumption-only API for barebones Tricking Terminology Data
// @termsOfService https://github.com/TrickingApi/trickingapi

// @contact.name Mikael Mantis
// @contact.url https://github.com/TrickingApi/trickingapi
// @contact.email mikael@trickingapi.dev

// @license.name MIT
// @license.url https://github.com/TrickingApi/trickingapi/blob/main/LICENSE.md

// @host trickingapi.dev
// @BasePath /api
// @schemes https
func main() {
	if len(os.Args) > 1 && os.Args[1] == "-scrape" {
		var startPath = ""
		var category = ""
		if len(os.Args) > 2 {
			startPath = os.Args[2]
		}

		if len(os.Args) > 3 {
			category = os.Args[3]
		}

		utils.Scrape(idToTrickMap, startPath, category)
		return
	}

	router := gin.Default()
	router.GET("/tricks", routes.GetAllTricksHandler(idToTrickMap))
	router.GET("/tricks/names", routes.GetAllTrickNamesHandler(idToTrickMap))
	router.GET("/tricks/:name", routes.GetTrickHandler(idToTrickMap))
	router.GET("/categories", routes.GetAllCategoriesHandler(categoriesToTrickSliceMap))
	router.GET("/categories/tricks", routes.GetAllTricksByCategoriesHandler(categoriesToTrickSliceMap))
	router.GET("/categories/:name", routes.GetTricksForCategoryHandler(categoriesToTrickSliceMap))

	router.StaticFile("swagger", "./docs/swagger.json")
	router.Run()
}
