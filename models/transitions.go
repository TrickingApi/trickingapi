package models

type Transition string

const (
	BACKSWING    Transition = "Backswing"
	BOUND        Transition = "Bound"
	CARRYTHROUGH Transition = "Carrythrough"
	FRONTSWING   Transition = "Frontswing"
	MISSLEG      Transition = "MISSleg"
	POP          Transition = "Pop"
	PUNCH        Transition = "Punch"
	RAPID        Transition = "Rapid"
	REDIRECT     Transition = "Redirect"
	REVERSAL     Transition = "Reversal"
	REVERSE_POP  Transition = "Reverse_Pop"
	SWINGTHROUGH Transition = "Swingthrough"
	SKIP         Transition = "Skip"
	VANISH       Transition = "Vanish"
	WRAPTHROUGH  Transition = "Wrapthrough"
)

func (ls Transition) String() string {
	return string(ls)
}

var Transitions = []Transition{BACKSWING, SWINGTHROUGH, FRONTSWING, VANISH, WRAPTHROUGH, SKIP, CARRYTHROUGH, MISSLEG, REVERSAL, POP, PUNCH, REVERSE_POP, REDIRECT, BOUND, RAPID}
