package models

type LandingStance string

const (
	COMPLETE LandingStance = "Complete"
	HYPER    LandingStance = "Hyper"
	MEGA     LandingStance = "Mega"
	SEMI     LandingStance = "Semi"
)

func (ls LandingStance) String() string {
	return string(ls)
}

var LandingStances = []LandingStance{COMPLETE, HYPER, MEGA, SEMI}
