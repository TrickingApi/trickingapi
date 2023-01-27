package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/TrickingApi/trickingapi/models"
	"github.com/iancoleman/strcase"
)

var BASE_URL string = "https://www.loopkickstricking.com/"
var OTHER_NAMES string = "Other Names:"
var PRE_REQS string = "Prerequisites:"

func Scrape(idToTrickMap map[string]models.Trick) {
	// Make the request
	resp, err := http.Get(BASE_URL + "tricktionary/explore")
	fmt.Println("Scraping")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("Validating response")
	// Validate the response
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	fmt.Println("Reading Response")
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Create a goquery document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		log.Fatal(err)
	}

	// Extract the trick URLs
	var trickURLs []string
	doc.Find(".w-dyn-item > a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			trickURLs = append(trickURLs, BASE_URL+link)
		}
	})

	// Scrape the trick information from each trick page
	for _, url := range trickURLs {
		// Make the request
		fmt.Println("Fetching: ", url)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		// Validate the response
		if resp.StatusCode != http.StatusOK {
			log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
		}

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Create a goquery document
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
		if err != nil {
			log.Fatal(err)
		}

		// Extract the trick information
		name := doc.Find(".heading-8").Text()
		otherNamesText := strings.ReplaceAll(doc.Find(".nickname-wrapper > .nickname").Text(), OTHER_NAMES, "")
		aliases := strings.Split(strings.Trim(otherNamesText, " "), ",")
		description := doc.Find(".paragraph-5").Text()
		prereqsText := doc.Find(".prereq-wrapper > .nickname.prereq + .nickname").Text()
		prereqs := strings.Split(strings.ReplaceAll(prereqsText, PRE_REQS, ""), ", ")

		var categories []models.TrickCategory
		var nextTricks []string
		currentTrick := &models.Trick{
			Id:            generateIdFromName(name),
			Name:          name,
			Aliases:       aliases,
			Categories:    categories,
			Prerequisites: prereqs,
			NextTricks:    nextTricks,
			Description:   description,
		}
		fmt.Println("Created Trick: ", currentTrick.Name)

		_, ok := idToTrickMap[currentTrick.Id]
		if !ok {
			idToTrickMap[currentTrick.Id] = *currentTrick
		}

		// searches for all prereqs and adds current trick to their nextTricks links
		fmt.Println("Filling Prereq Connections: ")
		for _, prereq := range currentTrick.Prerequisites {
			prereqId := generateIdFromName(prereq)
			prereqTrick, ok := idToTrickMap[prereqId]

			if ok {
				fmt.Println("Prereq connection created between {} - {}", prereqTrick.Name, currentTrick.Name)
				prereqTrick.NextTricks = append(prereqTrick.NextTricks, currentTrick.Name)
			}
		}
	}

	// convert idToTrickMap to array of Tricks, marshal to new json file
	var newTricksList []models.Trick
	for _, trick := range idToTrickMap {
		newTricksList = append(newTricksList, trick)
	}

	// Convert the array of tricks to a JSON string
	jsonData, err := json.MarshalIndent(newTricksList, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println("")
	fmt.Println("Writing to newTricks.json")
	// Write the JSON string to a file
	file, err := os.Create("newTricks.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		panic(err)
	}
}

func generateIdFromName(trickName string) string {
	return strcase.ToLowerCamel(trickName)
}
