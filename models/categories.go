package models

type TrickCategory string

const (
	BACKWARD   TrickCategory = "BACKWARD"
	DUB        TrickCategory = "DOUBLE"
	FLIP       TrickCategory = "FLIP"
	FORWARD    TrickCategory = "FORWARD"
	GROUNDWORK TrickCategory = "GROUNDWORK"
	INSIDE     TrickCategory = "INSIDE"
	OUTSIDE    TrickCategory = "OUTSIDE"
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

var Categories = []TrickCategory{FLIP, VERT_KICK, TWIST, PSEUDO_DUB, SING, DUB, TRIP, QUAD, GROUNDWORK, UNKNOWN, VARIATION, INSIDE, OUTSIDE, FORWARD, BACKWARD}

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
