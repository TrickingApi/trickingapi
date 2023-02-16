package routes

import (
	"encoding/json"
	"testing"

	"github.com/TrickingApi/trickingapi/models"
	"github.com/stretchr/testify/assert"
)

func TestGetAllLandingStanceIdsHandler(t *testing.T) {
	c, w := SetUpTests()

	var handler = GetAllLandingStanceIdsHandler()
	handler(c)
	assert.Equal(t, 200, w.Code)

	expectedResult := models.LandingStances

	var got []models.LandingStanceId
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	assert.ElementsMatch(t, got, expectedResult)
}
