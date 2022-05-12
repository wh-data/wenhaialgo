package play_with_arr

import (
	"fmt"
	"testing"
)

func TestARR(t *testing.T) {
	arr := []int{1, 2}
	fmt.Println(FindIndex(arr, 2))
	fmt.Println(remove_ele(arr, 2))
}

func TestAadd_on(t *testing.T) {
	arr := []int{1, 2, 3, 4, 8}
	s := 9
	fmt.Println(add_on(arr, s))
}
