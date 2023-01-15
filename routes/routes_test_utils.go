package routes

import (
	"net/http/httptest"

	"github.com/TrickingApi/trickingapi/models"
	"github.com/gin-gonic/gin"
)

func SetUpTests() (*gin.Context, *httptest.ResponseRecorder, models.Trick) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	dummyTrick := createMockData()

	return c, w, dummyTrick
}

func createMockData() models.Trick {
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
