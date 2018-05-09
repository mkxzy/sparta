package interpreter

const (
	NORMAL = iota
	//RETURN
	BREAK
	CONTINUE
)

/**
控制流接口
 */
type FlowState interface {

	SetState(s int)

	GetState() int
}
