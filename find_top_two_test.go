package wenhaialgo

import (
	"fmt"
	"testing"
)

func TestFindTopTwoMaxNumberByHeap(t *testing.T) {
	arr := []int{222, 2, 34, 6, 2, 45}
	//fmt.Println(FindTopTwoMaxNumber(arr))
	arr = []int{0, 1, 2, 3}
	fmt.Println(FindTopTwoMaxNumberV2(arr))

}
