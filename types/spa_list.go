package types

import "bytes"

type SPAList struct {
	arr []SPAValue
}

func NewList(cap int) *SPAList {
	return &SPAList{
		arr: make([]SPAValue, 0, cap),
	}
}

func (list *SPAList) Append(v SPAValue)  {
	list.arr = append(list.arr, v)
}

func (list *SPAList) Set(index SPAInteger, v SPAValue)  {
	list.arr[index] = v
}

func (list *SPAList) Get(index SPAInteger) SPAValue {
	return list.arr[index]
}

func(s *SPAList) IsTrue() bool  {
	return len(s.arr) > 0
}

func (s *SPAList) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i := 0; i < len(s.arr); i++{
		if i > 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(s.arr[i].String())
	}
	buffer.WriteString("]")
	return buffer.String()
}
