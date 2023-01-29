package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/TrickingApi/trickingapi/models"
	"github.com/iancoleman/strcase"
)

var BASE_URL string = "https://www.loopkickstricking.com/"
var OTHER_NAMES string = "Other Names:"
var PRE_REQS string = "Prerequisites:"
var DEFAULT_START_PATH string = "tricktionary/explore"

// List of Production Scrapes (1/28/23):
// https://www.loopkickstricking.com/tricktionary/explore
// https://www.loopkickstricking.com/tricktionary/vertical-kicks
// https://www.loopkickstricking.com/tricktionary/outside-tricks
// https://www.loopkickstricking.com/tricktionary/inside-tricks

func Scrape(idToTrickMap map[string]models.Trick, startPath string, category string) {
	trickURLs := extractTrickURLs(startPath)

	// Scrape the trick information from each trick page
	for _, url := range trickURLs {
		currentTrick := extractTrickInformation(url, category)
		if _, ok := idToTrickMap[currentTrick.Id]; !ok {
			idToTrickMap[currentTrick.Id] = currentTrick
		} else {
			fmt.Println("Merging existing data with new data for: ", currentTrick.Name)
			prevTrick := idToTrickMap[currentTrick.Id]
			mergedTrick := mergeTricks(prevTrick, currentTrick)
			idToTrickMap[mergedTrick.Id] = mergedTrick
		}
	}

	fmt.Println("Filling Prereq Connections: ")
	for _, currentTrick := range idToTrickMap {
		updatePrereqNextTrickEdges(currentTrick, idToTrickMap)
	}

	writeToFile(idToTrickMap)
}

func extractTrickURLs(startPath string) []string {
	var start string
	if start = startPath; len(startPath) > 0 {
	} else {
		start = DEFAULT_START_PATH
	}

	// Make the request
	resp, err := http.Get(BASE_URL + start)
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

	return trickURLs
}

func extractTrickInformation(url string, category string) models.Trick {
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
	name := strings.ReplaceAll(doc.Find(".heading-8").Text(), "\u00a0", " ")
	otherNamesText := strings.ReplaceAll(doc.Find(".nickname-wrapper > .nickname").Text(), OTHER_NAMES, "")
	otherNamesText = strings.ReplaceAll(otherNamesText, "\u00a0", "")
	aliases := strings.Split(strings.Trim(otherNamesText, " "), ",")
	description := strings.ReplaceAll(doc.Find(".paragraph-5").Text(), "\u00a0", " ")
	prereqsText := doc.Find(".prereq-wrapper > .nickname.prereq + .nickname").Text()
	prereqsText = strings.ReplaceAll(prereqsText, "\u00a0", " ")
	prereqs := strings.Split(strings.ReplaceAll(prereqsText, PRE_REQS, ""), ", ")

	var categories []models.TrickCategory

	if len(category) > 0 {
		if tc := models.GetCategoryFromString(category); tc != "" {
			categories = append(categories, tc)
		}
	}

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

	return *currentTrick
}

func updatePrereqNextTrickEdges(currentTrick models.Trick, idToTrickMap map[string]models.Trick) {
	// searches for all prereqs and adds current trick to their nextTricks links
	for _, prereq := range currentTrick.Prerequisites {
		fmt.Println("Checking if prereq exists: ", prereq)
		prereqId := generateIdFromName(prereq)
		if prereqTrick, ok := idToTrickMap[prereqId]; ok {
			if len(prereqTrick.Name) != 0 && !contains(prereqTrick.NextTricks, currentTrick.Name) {
				fmt.Println("Prereq connection created between: " + prereqTrick.Name + " - " + currentTrick.Name)
				prereqTrick.NextTricks = append(prereqTrick.NextTricks, currentTrick.Name)
				idToTrickMap[prereqId] = prereqTrick
			}
		}
	}
}

func writeToFile(idToTrickMap map[string]models.Trick) {
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

func contains(slice []string, element string) bool {
	for _, elem := range slice {
		if elem == element {
			return true
		}
	}
	return false
}

func mergeTricks(trickA models.Trick, trickB models.Trick) models.Trick {
	if trickA.Id != trickB.Id || trickA.Name != trickB.Name {
		panic("Merge tricks is meant for merging two trick objects with the same trickIds")
	}

	var mergedCategories []models.TrickCategory
	var mergedPrereqs, mergedNextTricks []string
	mergedCategories, okCategories := mergeUniqueElements(trickA.Categories, trickB.Categories).([]models.TrickCategory)
	mergedPrereqs, okPrereqs := mergeUniqueElements(trickA.Prerequisites, trickB.Prerequisites).([]string)
	mergedNextTricks, okNextTricks := mergeUniqueElements(trickA.NextTricks, trickB.NextTricks).([]string)

	if okCategories && okPrereqs && okNextTricks {
		trickA.Categories = mergedCategories
		trickA.Prerequisites = mergedPrereqs
		trickA.NextTricks = mergedNextTricks
	}
	return trickA
}

func mergeUniqueElements(s1, s2 interface{}) interface{} {
	aValue := reflect.ValueOf(s1)
	bValue := reflect.ValueOf(s2)

	if aValue.Kind() != reflect.Slice || bValue.Kind() != reflect.Slice {
		panic("Inputs must be slices")
	}

	aType := reflect.TypeOf(s1).Elem()

	// Create a map to store unique elements
	uniqueElements := make(map[interface{}]bool)

	// Add elements from s1 to the map
	for i := 0; i < aValue.Len(); i++ {
		uniqueElements[aValue.Index(i).Interface()] = true
	}
	// Add elements from s2 to the map
	for i := 0; i < bValue.Len(); i++ {
		uniqueElements[bValue.Index(i).Interface()] = true
	}

	// Create a slice to store the unique elements
	resultSlice := reflect.MakeSlice(reflect.SliceOf(aType), 0, len(uniqueElements))

	for key := range uniqueElements {
		resultSlice = reflect.Append(resultSlice, reflect.ValueOf(key))
	}

	return resultSlice.Interface()
}
