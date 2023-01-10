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
	categories := []models.TrickCategory{"Quad", "Pseudo Double Flip", "Flip", "Twist"}
	prereqs := []string{"Quad Full", "Frappe", "Triple Full In Frappe"}
	nextTricks := []string{"Quad Full In Frappe Kyro"}

	dummyTrick := models.Trick{
		Id:            "quadFullInFrappeOut",
		Name:          "Quad Full In Frappe",
		Categories:    categories,
		Prerequisites: prereqs,
		NextTricks:    nextTricks,
		Description:   "Impossible trick but ya never really know lmao",
	}
	return dummyTrick
}
