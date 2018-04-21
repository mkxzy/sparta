package vm

import "testing"

func TestNull(t *testing.T)  {
	n := Null()
	if !IsNull(n){
		t.Error("测试错误")
	} else{
		t.Log("正确")
	}
}

func TestSpaNull_IsTrue(t *testing.T) {
	n := Null()
	if n.IsTrue(){
		t.Error("错误")
	}else {
		t.Log("正确")
	}
}
