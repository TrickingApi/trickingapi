package routes

import (
	"encoding/json"
	"testing"

	"github.com/TrickingApi/trickingapi/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTricksHandler(t *testing.T) {
	c, w := SetUpTests()
	dummyTrick := CreateMockTrick()

	var idToTrickMap = make(map[string]models.Trick)
	idToTrickMap[dummyTrick.Id] = dummyTrick

	var handler = GetAllTricksHandler(idToTrickMap)
	handler(c)

	assert.Equal(t, 200, w.Code)

	var got map[string]models.Trick
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got, idToTrickMap)
}

func TestGetAllTrickByIdsHandler(t *testing.T) {
	c, w := SetUpTests()
	dummyTrick := CreateMockTrick()

	var idToTrickMap = make(map[string]models.Trick)
	idToTrickMap[dummyTrick.Id] = dummyTrick

	var handler = GetAllTricksByIdsHandler(idToTrickMap)
	handler(c)

	trickNames := map[string]string{"quadFullInFrappeOut": "Quad Full In Frappe"}
	assert.Equal(t, 200, w.Code)

	var got map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got, trickNames)
}

func TestGetAllTrickNamesHandler(t *testing.T) {
	c, w := SetUpTests()
	dummyTrick := CreateMockTrick()

	var idToTrickMap = make(map[string]models.Trick)
	idToTrickMap[dummyTrick.Id] = dummyTrick

	var handler = GetAllTrickNamesHandler(idToTrickMap)
	handler(c)

	trickNames := []string{"Quad Full In Frappe", "Quad Full In Frappe Out", "Venti Frappe"}
	assert.Equal(t, 200, w.Code)

	var got []string
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got, trickNames)
}

func TestGetTrickHandlerValid(t *testing.T) {
	c, w := SetUpTests()
	dummyTrick := CreateMockTrick()

	c.Params = []gin.Param{
		{
			Key:   "name",
			Value: "quadFullInFrappeOut",
		},
	}

	idToTrickMap := map[string]models.Trick{"quadFullInFrappeOut": dummyTrick}
	aliasesToTrickIds := map[string]string{"quadFullInFrappeOut": "quadFullInFrappeOut"}
	var handler = GetTrickHandler(idToTrickMap, aliasesToTrickIds)
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
	c, w := SetUpTests()
	dummyTrick := CreateMockTrick()

	c.Params = []gin.Param{
		{
			Key:   "name",
			Value: "quadBtwist",
		},
	}

	idToTrickMap := map[string]models.Trick{"quadFullInFrappeOut": dummyTrick}
	aliasesToTrickIds := map[string]string{"quadFullInFrappeOut": "quadFullInFrappeOut"}
	var handler = GetTrickHandler(idToTrickMap, aliasesToTrickIds)
	handler(c)

	assert.Equal(t, 404, w.Code)

	expectedErrorObj := models.TrickError{
		Message: "Oops, looks like you've requested an unknown trick! Do you think this is a mistake? Consider contributing at https://github.com/TrickingApi/trickingapi",
		Success: false,
		Data:    "Unknown trick: quadBtwist",
	}

	var got models.TrickError
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got, expectedErrorObj)
}
