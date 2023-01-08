package routes

import (
	"github.com/TrickingApi/trickingapi/models"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllTricksHandler(t *testing.T) {
	c, w, dummyTrick := SetUpTests()

	allTricks := []models.Trick{dummyTrick}
	
	var handler = GetAllTricksHandler(&allTricks)
	handler(c)

	assert.Equal(t, 200, w.Code)

	var got []models.Trick
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
			t.Fatal(err)
	}
	assert.Equal(t, got, allTricks)
}

func TestGetAllTrickNamesHandler(t *testing.T) {
	c, w, dummyTrick := SetUpTests()

	allTricks := []models.Trick{dummyTrick}

	var handler = GetAllTrickNamesHandler(&allTricks)
	handler(c)

	trickNames := []string{"Quad Full In Frappe"}
	assert.Equal(t, 200, w.Code)

	var got []string
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
			t.Fatal(err)
	}
	assert.Equal(t, got, trickNames)
}

func TestGetTrickHandlerValid(t *testing.T) {
	c, w, dummyTrick := SetUpTests()
	c.Params = []gin.Param{
    {
        Key: "name",
        Value: "quadFullInFrappeOut",
    },
	}

	idToTrickMap := map[string]models.Trick{"quadFullInFrappeOut": dummyTrick}
	var handler = GetTrickHandler(idToTrickMap)
	handler(c)

	assert.Equal(t, 200, w.Code)

	var got models.Trick
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got, dummyTrick)
}

func TestGetTrickHandlerUnknown(t *testing.T) {
	c, w, dummyTrick := SetUpTests()
	c.Params = []gin.Param{
    {
        Key: "name",
        Value: "quadBtwist",
    },
	}

	idToTrickMap := map[string]models.Trick{"quadFullInFrappeOut": dummyTrick}
	var handler = GetTrickHandler(idToTrickMap)
	handler(c)

	assert.Equal(t, 404, w.Code)

	expectedErrorObj := models.TrickError{
		Message: "Oops, looks like you've requested an unknown trick! Do you think this is a mistake? Consider contributing at https://github.com/TrickingApi/trickingapi",
		Success: false,
		Data: "Unknown trick: quadBtwist",
	}

	var got models.TrickError
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got, expectedErrorObj)
}