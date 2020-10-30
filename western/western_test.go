package western

import (
	"fmt"
	"testing"
)

func TestScaleFormula(t *testing.T) {
	var major = []int{2, 2, 1, 2, 2, 2, 1}

	got := ScaleFormula("major")

	if len(major) != len(got) {
		fmt.Printf("Expected %v got %v", major, got)
		return
	}

	for i := range major {
		if major[i] != got[i] {
			return
		}
	}
}
