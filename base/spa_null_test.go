package base

import "testing"

func TestNull(t *testing.T)  {
	n := Null()
	if !IsNull(n){
		t.Error("失败")
	}
}
