package vm

type Instruction int
type ConstantPool []interface{}

const (
	MAX_CODE = 65536
)

var code = make([]Instruction, MAX_CODE) //代码存储器
var ip = 0 //指令计数器
var fp = 0 //帧寄存器
var constantPool ConstantPool //常量池

func cpu()  {
}