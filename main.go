package main

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"io/ioutil"
	"log"
	"github.com/TrickingApi/trickingapi/models"
	"github.com/TrickingApi/trickingapi/routes"
)

var allTricks []models.Trick
var idToTrickMap map[string]models.Trick
var categoriesToTrickSliceMap map[models.TrickCategory][]models.Trick

func init() {
	content, err := ioutil.ReadFile("data/tricks.json")
	if err != nil {
			log.Fatal("Error when opening file: ", err)
	}

	if err := json.Unmarshal(content, &allTricks); err != nil {
		panic(err)
	}

	idToTrickMap = make(map[string]models.Trick)
	categoriesToTrickSliceMap = make(map[models.TrickCategory][]models.Trick)

	for _, trick := range allTricks {
		idToTrickMap[trick.Id] = trick

		var categories = trick.Categories

		for _, category := range categories {
			tc := models.TrickCategory(category)

			if tricks, ok := categoriesToTrickSliceMap[tc]; ok {
				tricks = append(tricks, trick)
				categoriesToTrickSliceMap[tc] = tricks
			} else {
				slice := []models.Trick{trick}
				categoriesToTrickSliceMap[tc] = slice
			}
		}
		
	}
}

// @title Tricking Api
// @version 0.1
// @description Consumption-only API for barebones Tricking Data
// @termsOfService https://github.com/TrickingApi/trickingapi

// @contact.name Mikael Mantis
// @contact.url https://github.com/TrickingApi/trickingapi
// @contact.email mikael.mantis7@gmail.com

// @license.name MIT
// @license.url https://github.com/TrickingApi/trickingapi/blob/main/LICENSE.md

// @host api.trickingapi.dev
// @BasePath /
// @schemes https
func main() {
	router := gin.Default()
	router.GET("/tricks", routes.GetAllTricksHandler(&allTricks))
	router.GET("/tricks/names", routes.GetAllTrickNamesHandler(&allTricks))
	router.GET("/tricks/:name", routes.GetTrickHandler(idToTrickMap))
	router.GET("/categories", routes.GetTricksByCategoriesHandler(categoriesToTrickSliceMap))
	router.GET("/categories/:name", routes.GetCategoryToTricksHandler(categoriesToTrickSliceMap))

	router.StaticFile("swagger", "./docs/swagger.json")
	router.Run()
}
