package shift_string

import (
	"fmt"
	"testing"
)

func TestShift(t *testing.T) {
	str := "s新加坡人"
	shift := 4
	fmt.Println(shift_string(str, shift))
	fmt.Println(shift_string_v2(str, shift))
	fmt.Println(reverse_string(str))
}
