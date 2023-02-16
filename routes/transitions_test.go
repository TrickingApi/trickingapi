package routes

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/TrickingApi/trickingapi/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTransitionIdsHandler(t *testing.T) {
	c, w := SetUpTests()

	var handler = GetAllTransitionIdsHandler()
	handler(c)
	assert.Equal(t, 200, w.Code)

	expectedResult := models.TransitionIds

	var got []models.TransitionId
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	assert.ElementsMatch(t, got, expectedResult)
}

func TestGetAllTransitionsHandler(t *testing.T) {
	c, w := SetUpTests()
	dummyTransition := CreateMockTransition()

	transitionMap := createMockTransitionMap(dummyTransition)

	var handler = GetAllTransitionsHandler(transitionMap)
	handler(c)

	expectedSlice := []models.Transition{dummyTransition}

	assert.Equal(t, 200, w.Code)

	var got []models.Transition
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	assert.ElementsMatch(t, got, expectedSlice)
}

func TestGetTransitionHandler(t *testing.T) {
	c, w := SetUpTests()
	dummyTransition := CreateMockTransition()

	transitionMap := createMockTransitionMap(dummyTransition)

	c.Params = []gin.Param{
		{
			Key:   "id",
			Value: "vanish",
		},
	}

	var handler = GetTransitionHandler(transitionMap)
	handler(c)

	assert.Equal(t, 200, w.Code)

	var got models.Transition
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, got, dummyTransition)
}

func createMockTransitionMap(dummyTransition models.Transition) map[string]models.Transition {
	transitionMap := make(map[string]models.Transition)

	transitionMap[strings.ToUpper(dummyTransition.Id.String())] = dummyTransition
	return transitionMap
}
