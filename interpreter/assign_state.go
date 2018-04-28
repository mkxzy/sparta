package interpreter

// 状态类型
const (
	NewVariable = iota
	ListAccess
	//MapAccess
)

type AssignState struct {
	Name string //变量名称
	State int 	//赋值类型
}