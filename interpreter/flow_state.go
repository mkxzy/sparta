package interpreter

const (
	NORMAL = iota
	RETURN
	BREAK
	CONTINUE
)

type FlowState interface {

	SetState(s int)

	GetState() int
}
