package models

type TakeOff string

const (
	BACKSIDE  TakeOff = "Backside"
	FRONTSIDE TakeOff = "Frontside"
	NEUTRAL   TakeOff = "Neutral"
)

func (to TakeOff) String() string {
	return string(to)
}

var TakeOffs = []TakeOff{BACKSIDE, FRONTSIDE}
