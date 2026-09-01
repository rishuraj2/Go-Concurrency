package sharedstate

type State int

const (
	ZERO_BEFORE_ODD State = iota
	ODD
	ZERO_BEFORE_EVEN
	EVEN
)
