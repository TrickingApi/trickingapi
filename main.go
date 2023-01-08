package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"tricks/web-service/models"
)

var allTricks []models.Trick

func init() {
	content, err := ioutil.ReadFile("data/tricks.json")
	if err != nil {
			log.Fatal("Error when opening file: ", err)
	}


	if err := json.Unmarshal(content, &allTricks); err != nil {
		panic(err)
	}
	fmt.Println(allTricks)
}

func main() {
	router := gin.Default()
	router.GET("/tricks", getAllTricks)

	router.Run("localhost:8080")
}

func getAllTricks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, allTricks)
}

