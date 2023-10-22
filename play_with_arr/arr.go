package play_with_arr

import (
	"fmt"
)

// 1,-1,2
func ttt(arr []int) int {
	return 0
}

func longestConsecutive(nums []int) int {
	mp := make(map[int]int, 0)
	for _, v := range nums {
		if _, ok := mp[v]; !ok {
			left := 0
			right := 0
			if _, ok := mp[v-1]; ok {
				left = mp[v-1]
			}
			if _, ok := mp[v+1]; ok {
				right = mp[v+1]
			}
			mp[v] = 1 + left + right
			mp[v-left] = mp[v] //边界也重新设值
			mp[v+right] = mp[v]
			fmt.Println(v-left, mp[v-left])
			fmt.Println(v+right, mp[v+right])
			fmt.Println("===============")
		}

	}
	longest := 0
	for _, v := range mp {
		if v > longest {
			longest = v
		}
	}

	return longest
}

// 有序的数组
// 找到所对应值所在的index
func FindIndex(arr []int, val int) int {
	lt := len(arr)
	return findIndex(arr, val, 0, lt-1)
}

func findIndex(arr []int, val int, l, r int) int {
	if r < 0 {
		return -1
	}
	if l > r {
		return -1
	}
	po := (r + l) / 2
	if arr[po] == val {
		return po
	}
	if arr[po] > val {
		return findIndex(arr, val, l, po)
	}
	if arr[po] < val {
		return findIndex(arr, val, po+1, r)
	}
	return -1
}

// 从数组移除某个值
func remove_ele(arr []int, val int) int {
	fast, slow := 0, -1
	l := len(arr)
	for l > 0 {
		if arr[fast] == val {
			fast++
		} else {
			slow++
			arr[slow] = arr[fast]
			fast++
		}
		l--
	}
	return slow + 1
}

// 最短的window，相加和>=s
func add_on(arr []int, s int) int {
	fast, slow := 0, 0
	res := len(arr) + 1
	sum := 0
	for fast < len(arr) {
		if arr[fast] >= s {
			return 1
		}
		sum += arr[fast]
		if sum >= s {
			if res > fast-slow+1 {
				res = fast - slow + 1
			}
			for slow <= fast {
				fmt.Println("Sadfd: ", slow, fast)
				sum -= arr[slow]
				slow++
				if sum >= s {
					if res > fast-slow+1 {
						res = fast - slow + 1
					}
				} else {
					break
				}
			}
		}
		fast++
	}
	if res == len(arr)+1 {
		return 0
	}
	return res
}

// n is positive integer
func Spiral_Matrix(n int) [][]int {
	matrix := make([][]int, n)
	for i, _ := range matrix {
		matrix[i] = make([]int, n)
	}
	x0, y0 := 0, 0
	xn, yn := n-1, n-1
	num := 1
	for num <= n*n {
		for i := x0; i <= xn; i++ {
			matrix[i][y0] = num
			num++
		}
		for j := y0 + 1; j <= yn; j++ {
			matrix[xn][j] = num
			num++
		}
		for i := xn - 1; i >= x0; i-- {
			matrix[i][yn] = num
			num++
		}
		for j := yn - 1; j > y0; j-- {
			matrix[x0][j] = num
			num++
		}
		x0++
		xn--
		y0++
		yn--
	}
	return matrix
}
