package models

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTricksModelFromJson(t *testing.T) {
	var allTricks []Trick
	content, err := os.ReadFile("test_data/test.json")
	if err != nil {
		t.Fatal("Error when opening file: ", err)
	}

	if err := json.Unmarshal(content, &allTricks); err != nil {
		t.Fatal(err)
	}

	expectedTricks := getMockTricks()

	assert.Equal(t, expectedTricks, allTricks)
}

func getMockTricks() []Trick {
	categories := []TrickCategory{VERT_KICK}
	aliases := []string{}

	prereqsA := []string{}
	nextTricksA := []string{
		"Pop 360 Shuriken",
		"Pop 720",
		"Swing 360",
		"Illusion Twist",
	}

	trickA := Trick{
		Id:            "pop360",
		Name:          "Pop 360",
		Aliases:       aliases,
		Categories:    categories,
		Prerequisites: prereqsA,
		NextTricks:    nextTricksA,
		Description:   "Test description",
	}

	prereqsB := []string{
		"Pop 360",
	}

	nextTricksB := []string{
		"Pop 720",
		"Pop 720 Double",
		"Swing 360 Shuriken",
	}

	trickB := Trick{
		Id:            "pop360Shuriken",
		Name:          "Pop 360 Shuriken",
		Aliases:       aliases,
		Categories:    categories,
		Prerequisites: prereqsB,
		NextTricks:    nextTricksB,
		Description:   "Test description 2",
	}

	result := []Trick{trickA, trickB}
	return result
}
