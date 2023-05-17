package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/TrickingApi/trickingapi/models"
	"github.com/TrickingApi/trickingapi/routes"
	"github.com/TrickingApi/trickingapi/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var allTricks []models.Trick
var idToTrickMap map[string]models.Trick
var categoriesToTrickSliceMap map[models.TrickCategory][]models.Trick
var aliasesToTrickIds map[string]string
var transitions map[string]models.Transition
var landingStances map[string]models.LandingStance

// organizes tricks by id and category upon server init
func init() {
	loadTricksData()
	loadTransitionData()
	loadLandingStances()
}

func loadTricksData() {
	content, err := os.ReadFile("data/tricks.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	if err := json.Unmarshal(content, &allTricks); err != nil {
		panic(err)
	}

	// initialize maps
	idToTrickMap = make(map[string]models.Trick)
	categoriesToTrickSliceMap = make(map[models.TrickCategory][]models.Trick)
	aliasesToTrickIds = make(map[string]string)

	// initialize empty slices for each category
	for _, category := range models.Categories {
		emptySlice := []models.Trick{}
		categoriesToTrickSliceMap[category] = emptySlice
	}

	// organize tricks by id, category, and alias
	for _, trick := range allTricks {
		idToTrickMap[trick.Id] = trick
		aliasesToTrickIds[utils.FormatAlias(trick.Name)] = trick.Id
		aliasesToTrickIds[utils.FormatAlias(trick.Id)] = trick.Id

		var categories = trick.Categories

		for _, category := range categories {
			tc := models.TrickCategory(category)

			if tricks, ok := categoriesToTrickSliceMap[tc]; ok {
				tricks = append(tricks, trick)
				categoriesToTrickSliceMap[tc] = tricks
			}
		}

		for _, alias := range trick.Aliases {
			formattedAlias := utils.FormatAlias(alias)
			aliasesToTrickIds[formattedAlias] = trick.Id
		}
	}
}

func loadTransitionData() {
	content, err := os.ReadFile("data/transitions.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	if err := json.Unmarshal(content, &transitions); err != nil {
		panic(err)
	}
}

func loadLandingStances() {
	content, err := os.ReadFile("data/landingstances.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	if err := json.Unmarshal(content, &landingStances); err != nil {
		panic(err)
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
	router.SetTrustedProxies(nil)
	config := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET"},
		AllowHeaders:    []string{"*"}, // allow all headers
		ExposeHeaders:   []string{"Content-Length", "Content-Type"},
	}
	//config.AddHeaders(")
	router.Use(cors.New(config))
	router.GET("/", routes.GetRootHandler())
	router.GET("/tricks", routes.GetAllTricksHandler(idToTrickMap))
	router.GET("/tricks/ids", routes.GetAllTricksByIdsHandler(idToTrickMap))
	router.GET("/tricks/names", routes.GetAllTrickNamesHandler(idToTrickMap))
	router.GET("/tricks/:name", routes.GetTrickHandler(idToTrickMap, aliasesToTrickIds))
	router.GET("/categories", routes.GetAllCategoriesHandler(categoriesToTrickSliceMap))
	router.GET("/categories/tricks", routes.GetAllTricksByCategoriesHandler(categoriesToTrickSliceMap))
	router.GET("/categories/:name", routes.GetTricksForCategoryHandler(categoriesToTrickSliceMap))
	router.GET("/transitions/ids", routes.GetAllTransitionIdsHandler())
	router.GET("/transitions", routes.GetAllTransitionsHandler(transitions))
	router.GET("/transitions/:id", routes.GetTransitionHandler(transitions))
	router.GET("/landingstances", routes.GetAllLandingStancesHandler(landingStances))
	router.GET("/landingstances/ids", routes.GetAllLandingStanceIdsHandler())
	router.GET("/landingstances/:id", routes.GetLandingStanceByIdHandler(landingStances))
	router.Run()
}
