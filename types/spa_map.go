package types

import "fmt"

type SPAMap struct {
	mp map[SPAValue]SPAValue
}

func NewMap() *SPAMap {
	return &SPAMap{
		mp: make(map[SPAValue]SPAValue),
	}
}

func(v *SPAMap) String() string {
	return fmt.Sprintf("%v", v.mp)
}

func (v *SPAMap) Put(key, value SPAValue)  {
	v.mp[key] = value
}

func (v *SPAMap) Get(key SPAValue) SPAValue {
	value := v.mp[key]
	if value == nil {
		return Null()
	}
	return value
}

func(s SPAMap) IsTrue() bool  {
	return len(s.mp) > 0
}
