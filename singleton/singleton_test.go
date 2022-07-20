package singleton

import (
	"fmt"
	"testing"
)

func Test_singleton(t *testing.T) {
	s := NewSingleton()
	s["this"] = "that"

	s2 := NewSingleton()
	fmt.Println(s2["this"])
}
