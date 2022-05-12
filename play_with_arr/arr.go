package play_with_arr

import "fmt"

//有序的数组
//找到所对应值所在的index
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

//从数组移除某个值
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

//最短的window，相加和>=s
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
