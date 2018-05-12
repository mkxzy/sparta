package interpreter

import (
	"testing"
	"fmt"
	"path/filepath"
)

func TestFilePath(t *testing.T)  {
	filePath, _ := filepath.Abs("../example/for_test.spa")
	fmt.Println(filePath)
	fmt.Println(filepath.Dir(filePath))
}
