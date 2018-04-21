package vm

var operands [3]SPAValue
var sp = -1

func PushValue(v SPAValue)  {
	sp++
	operands[sp] = v
}

func PopValue() SPAValue  {
	v := operands[sp]
	sp--
	return v
}

func PeekValue() SPAValue {
	if sp > -1 {
		return operands[sp]
	}
	return Null()
}
