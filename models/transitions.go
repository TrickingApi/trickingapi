package models

type TransitionId string

const (
	BACKSWING    TransitionId = "Backswing"
	BOUND        TransitionId = "Bound"
	CARRYTHROUGH TransitionId = "Carrythrough"
	FRONTSWING   TransitionId = "Frontswing"
	MISSLEG      TransitionId = "MISSleg"
	POP          TransitionId = "Pop"
	PUNCH        TransitionId = "Punch"
	RAPID        TransitionId = "Rapid"
	REDIRECT     TransitionId = "Redirect"
	REVERSAL     TransitionId = "Reversal"
	REVERSE_POP  TransitionId = "Reverse_Pop"
	SWINGTHROUGH TransitionId = "Swingthrough"
	SKIP         TransitionId = "Skip"
	VANISH       TransitionId = "Vanish"
	WRAPTHROUGH  TransitionId = "Wrapthrough"
)

func (tsId TransitionId) String() string {
	return string(tsId)
}

var TransitionIds = []TransitionId{BACKSWING, SWINGTHROUGH, FRONTSWING, VANISH, WRAPTHROUGH, SKIP, CARRYTHROUGH, MISSLEG, REVERSAL, POP, PUNCH, REVERSE_POP, REDIRECT, BOUND, RAPID}

type Transition struct {
	Id          TransitionId `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Aliases     []string     `json:"aliases"`
	Examples    []string     `json:"examples"`
}
