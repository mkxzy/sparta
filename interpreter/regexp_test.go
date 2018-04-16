package interpreter

import (
	"testing"
	"regexp"
)

func TestReplace(t *testing.T)  {
	r := regexp.MustCompile(`\\t`)
	s := "ninin\\t"
	s = r.ReplaceAllLiteralString(s, "")
	t.Log(s)
}
