package models

type Trick struct {
	Id	string `json:"id"`
	Name string `json:"name"`
	Categories []TrickCategory `json:"categories"`
	DifficultyRank int `json:"difficultRank"`
	Prerequisites []string `json:"prerequisites"`
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

/*func (tc TrickCategory) String() string {
	switch tc {
		case FLIP:
			return "Flip"
		case VERT_KICK:
			return "Vert Kick"
		case TWIST:
			return "Twist"
		case PSEUDO_DUB:
			return "Pseudo Double Flip"
		case SING:
			return "Single"
		case DUB:
			return "Double"
		case TRIP:
			return "Triple"
		case QUAD:
			return "Quad"
	}
	return "unknown"
}*/