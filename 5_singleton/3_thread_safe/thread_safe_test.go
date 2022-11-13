package thread_safe

import (
	"fmt"
	"testing"
)

func TestThreadSafe(t *testing.T) {
	singleton := GetInstance()
	fmt.Println(singleton.GetDescription())
}
