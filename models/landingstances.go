package models

type LandingStanceId string

const (
	COMPLETE LandingStanceId = "Complete"
	HYPER    LandingStanceId = "Hyper"
	MEGA     LandingStanceId = "Mega"
	SEMI     LandingStanceId = "Semi"
)

func (ls LandingStanceId) String() string {
	return string(ls)
}

var LandingStances = []LandingStanceId{COMPLETE, HYPER, MEGA, SEMI}

type LandingStance struct {
	Id          LandingStanceId `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Examples    []string        `json:"examples"`
}
