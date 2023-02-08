package routes

import (
	"net/http/httptest"

	"github.com/TrickingApi/trickingapi/models"
	"github.com/gin-gonic/gin"
)

func SetUpTests() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	return c, w
}

func CreateMockTrick() models.Trick {
	categories := []models.TrickCategory{"QUAD", "PSEUDO_DOUBLE_FLIP", "FLIP", "TWIST"}
	prereqs := []string{"Quad Full", "Frappe", "Triple Full In Frappe"}
	nextTricks := []string{"Quad Full In Frappe Kyro"}
	aliases := []string{}

	dummyTrick := models.Trick{
		Id:            "quadFullInFrappeOut",
		Name:          "Quad Full In Frappe",
		Aliases:       aliases,
		Categories:    categories,
		Prerequisites: prereqs,
		NextTricks:    nextTricks,
		Description:   "Impossible trick but ya never really know lmao",
	}
	return dummyTrick
}

func CreateMockTransition() models.Transition {
	aliases := []string{}
	examples := []string{"Triple buterflytwist (vanish) cheat 1620"}
	dummyTransition := models.Transition{
		Id:          "vanish",
		Name:        "Vanish",
		Description: "Vanishing vanish vanishes",
		Aliases:     aliases,
		Examples:    examples,
	}
	return dummyTransition
}
