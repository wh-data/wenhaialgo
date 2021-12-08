package dynamic_programming

//青蛙上楼梯，可以上一节台阶或2节台阶，到n节台阶有几种上法
func FragGoStair(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return FragGoStair(n-1) + FragGoStair(n-2)
}

//var FragGoStairPath [][]int

//青蛙上楼梯，可以上一节台阶或2节台阶，打印出所有可能性
func FragGoStairPossibility(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}
	if n == 2 {
		return [][]int{{1, 1}, {2}}
	}
	arr1 := FragGoStairPossibility(n - 1)
	arr2 := FragGoStairPossibility(n - 2)
	for i := range arr1 {
		arr1[i] = append(arr1[i], 1)
	}
	for i := range arr2 {
		arr2[i] = append(arr2[i], 2)
	}
	return append(arr2, arr1...)
}

type Article struct {
	Weight  int
	Value   int
	VwRatio float32
}
