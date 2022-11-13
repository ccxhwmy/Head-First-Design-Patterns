package classic

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	singleton := GetInstance()
	fmt.Println(singleton.getDescription())
}
