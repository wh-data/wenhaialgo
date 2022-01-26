package wenhaialgo

import (
	"fmt"
	"testing"
)

func TestDeduplicate(t *testing.T) {
	x := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	l := DeduplicateByForLoop(x)
	fmt.Println(l, x)
	x = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	l = DeduplicateByPaceDiff(x)
	fmt.Println(l, x)
}
