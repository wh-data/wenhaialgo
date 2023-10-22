package wenhaialgo

import (
	"fmt"
	"testing"
)

// 每一个char排到开头 + 剩余组合的全排列
func Test_CalcALLPermutation(t *testing.T) {
	str := "abc"
	result := make([]string, 0)
	CalcALLPermutation(str, 0, len([]rune(str))-1, &result)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

func Test_CalcALLPermutationInt(t *testing.T) {
	arr := []int{1, 2, 3}
	result := make([][]int, 0)
	CalcALLPermutationInt(arr, 0, len(arr)-1, &result)
	//for i := 0; i < len(result); i++ {
	//	fmt.Println(result[i])
	//}

}

func Test_ModifyArr(t *testing.T) {
	arr := []int{1, 2}
	ModifyArr(arr)
	fmt.Println(arr)
}
