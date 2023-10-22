package wenhaialgo

import "fmt"

// 每一个char排到开头 + 剩余组合的全排列
func CalcALLPermutation(str string, from, to int, result *[]string) {
	str_rune := []rune(str)
	if from > to {
		return
	}
	if from == to {
		*result = append(*result, str)
		return
	}
	for i := from; i <= to; i++ {
		str_rune[i], str_rune[from] = str_rune[from], str_rune[i]
		str = string(str_rune)
		CalcALLPermutation(str, from+1, to, result)
	}
}

func CalcALLPermutationInt(arr []int, from, to int, result *[][]int) {
	tem := make([]int, len(arr))
	copy(tem, arr)
	if from > to {
		return
	}
	if from == to {
		//*result = append(*result, *arr)
		fmt.Println(arr)
		return
	}
	for i := from; i <= to; i++ {
		//因为arr是引用类型，会被下一层的调用改动，所以每次回到该层需要恢复
		copy(arr, tem)
		arr[i], arr[from] = arr[from], arr[i]
		//fmt.Println(arr)
		CalcALLPermutationInt(arr, from+1, to, result)
	}
}

func ModifyArr(arr []int) {
	arr[0], arr[1] = arr[1], arr[0]
}
