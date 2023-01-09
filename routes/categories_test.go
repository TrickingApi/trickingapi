package routes

import (
	"encoding/json"
	"github.com/TrickingApi/trickingapi/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestGetAllTricksByCategoriesHandler(t *testing.T) {
	c, w, dummyTrick := SetUpTests()
	categoriesToTrickSliceMap := createMockCategoriesMap(dummyTrick)

	var handler = GetAllTricksByCategoriesHandler(categoriesToTrickSliceMap)
	handler(c)

	assert.Equal(t, 200, w.Code)

	var got map[models.TrickCategory][]models.Trick
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got, categoriesToTrickSliceMap)
}

func TestGetTricksForCategoryHandlerQuad(t *testing.T) {
	c, w, dummyTrick := SetUpTests()
	c.Params = []gin.Param{
		{
			Key: "name",
			Value: "Quad",
		},
	}

	categoriesToTrickSliceMap := createMockCategoriesMap(dummyTrick)
	var handler = GetTricksForCategoryHandler(categoriesToTrickSliceMap)
	handler(c)

	assert.Equal(t, 200, w.Code)
	var got []models.Trick
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	expectedSlice := []models.Trick{dummyTrick}
	assert.Equal(t, got, expectedSlice)
}

func TestGetTricksForCategoryHandlerPsuedoDub(t *testing.T) {
	c, w, dummyTrick := SetUpTests()
	c.Params = []gin.Param{
		{
			Key: "name",
			Value: "Pseudo Double Flip",
		},
	}

	categoriesToTrickSliceMap := createMockCategoriesMap(dummyTrick)
	var handler = GetTricksForCategoryHandler(categoriesToTrickSliceMap)
	handler(c)

	assert.Equal(t, 200, w.Code)
	var got []models.Trick
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	expectedSlice := []models.Trick{dummyTrick}
	assert.Equal(t, got, expectedSlice)
}

func TestGetTricksForCategoryHandlerUnknown(t *testing.T) {
	c, w, dummyTrick := SetUpTests()
	c.Params = []gin.Param{
		{
			Key: "name",
			Value: "Multi-Level Movement",
		},
	}

	categoriesToTrickSliceMap := createMockCategoriesMap(dummyTrick)
	var handler = GetTricksForCategoryHandler(categoriesToTrickSliceMap)
	handler(c)

	assert.Equal(t, 404, w.Code)

	expectedErrorObj := models.TrickError{
		Message: "Oops, looks like you've requested an unknown category of tricks! Do you think this is a mistake? Consider contributing at https://github.com/TrickingApi/trickingapi",
		Success: false,
		Data: "Unknown category: Multi-Level Movement",
	}

	var got models.TrickError
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, got, expectedErrorObj)
}

func TestGetAllCategoriesHandler(t *testing.T) {
	c, w, dummyTrick := SetUpTests()
	categoriesToTrickSliceMap := createMockCategoriesMap(dummyTrick)
	var handler = GetAllCategoriesHandler(categoriesToTrickSliceMap)
	handler(c)

	assert.Equal(t, 200, w.Code)

	expectedResult := []string{"Quad", "Pseudo Double Flip", "Flip", "Twist"}

	var got []string
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	assert.ElementsMatch(t, got, expectedResult)
}

func createMockCategoriesMap(dummyTrick models.Trick) (map[models.TrickCategory][]models.Trick) {
	trickSlice := []models.Trick{dummyTrick}
	categoriesToTrickSliceMap := map[models.TrickCategory][]models.Trick{
		"Quad": trickSlice,
		"Pseudo Double Flip": trickSlice,
		"Flip": trickSlice,
		"Twist": trickSlice,
	}
	return categoriesToTrickSliceMap
}