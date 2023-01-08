package models

type Trick struct {
	Id	string `json:"id"`
	Name string `json:"name"`
	Categories []TrickCategory `json:"categories"`
	DifficultyRank int `json:"difficultRank"`
	Prerequisites []string `json:"prereqs"`
	NextTricks []string `json:"nextTricks"`
	Description string `json:"description"`
}

type TrickCategory string

const (
	FLIP TrickCategory = "Flip"
  VERT_KICK TrickCategory = "Vert Kick"
  TWIST TrickCategory = "Twist"
  PSEUDO_DUB TrickCategory = "Pseudo Double Flip"
  SING TrickCategory = "Single"
  DUB TrickCategory = "Double"
  TRIP TrickCategory = "Triple"
  QUAD TrickCategory = "Quad"
)

func (tc TrickCategory)String() string {
	return string(tc)
}