package models

type TrickCategory string

const (
	FLIP        TrickCategory = "FLIP"
	VERT_KICK   TrickCategory = "VERT_KICK"
	TWIST       TrickCategory = "TWIST"
	PSEUDO_DUB  TrickCategory = "PSEUDO_DOUBLE_FLIP"
	SING        TrickCategory = "SINGLE"
	DUB         TrickCategory = "DOUBLE"
	TRIP        TrickCategory = "TRIPLE"
	QUAD        TrickCategory = "QUAD"
	GROUND_WORK TrickCategory = "GROUND_WORK"
	UNKNOWN     TrickCategory = "UNKNOWN"
)

func (tc TrickCategory) String() string {
	return string(tc)
}

var Categories = []TrickCategory{FLIP, VERT_KICK, TWIST, PSEUDO_DUB, SING, DUB, TRIP, QUAD, GROUND_WORK, UNKNOWN}
