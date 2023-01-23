package models

type Transition string

const (
	BACK_SWING    Transition = "Back_Swing"
	SWING         Transition = "Swing"
	FRONT_SWING   Transition = "Front_Swing"
	VANISH        Transition = "Vanish"
	WRAP          Transition = "Wrap"
	SKIP          Transition = "Skip"
	CARRY_THROUGH Transition = "Carry_Through"
	MISLEG        Transition = "Misleg"
	REVERSAL      Transition = "Reversal"
	POP           Transition = "Pop"
	PUNCH         Transition = "Punch"
)

func (ls Transition) String() string {
	return string(ls)
}

var Transitions = []Transition{BACK_SWING, SWING, FRONT_SWING, VANISH, WRAP, SKIP, CARRY_THROUGH, MISLEG, REVERSAL, POP, PUNCH}
