package interpreter

import "github.com/mkxzy/sparta/base"

type SPAState struct {
	vars map[string]base.SPAValue
}

func NewState() *SPAState {
	state := &SPAState{
		vars: make(map[string]base.SPAValue),
	}
	return state
}

func(v *SPAState) PutVar(name string, value base.SPAValue)  {
	v.vars[name] = value
}

func(v *SPAState) GetVar(name string) base.SPAValue  {
	return v.vars[name]
}
