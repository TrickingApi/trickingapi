package models

type TrickCategory string

const (
	DUB        TrickCategory = "DOUBLE"
	FLIP       TrickCategory = "FLIP"
	GROUNDWORK TrickCategory = "GROUNDWORK"
	PSEUDO_DUB TrickCategory = "PSEUDO_DOUBLE_FLIP"
	QUAD       TrickCategory = "QUAD"
	SING       TrickCategory = "SINGLE"
	TRIP       TrickCategory = "TRIPLE"
	TWIST      TrickCategory = "TWIST"
	UNKNOWN    TrickCategory = "UNKNOWN"
	VARIATION  TrickCategory = "VARIATION"
	VERT_KICK  TrickCategory = "VERT_KICK"
)

func (tc TrickCategory) String() string {
	return string(tc)
}

var Categories = []TrickCategory{FLIP, VERT_KICK, TWIST, PSEUDO_DUB, SING, DUB, TRIP, QUAD, GROUNDWORK, UNKNOWN}

func IsCategory(str string) bool {
	return GetCategoryFromString(str) != ""
}

func GetCategoryFromString(str string) TrickCategory {
	for _, category := range Categories {
		if str == string(category) {
			return category
		}
	}
	return ""
}
