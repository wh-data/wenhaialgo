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

func TestSSpiral_Matrix(t *testing.T) {
	n := 1
	m := Spiral_Matrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(m[j][i], ",")
		}
		fmt.Println()
	}
}

func TestMaxAverage_bad_solution(t *testing.T) {
	//bad solution
	arr := []int{1, 2, 4, 5}
	k := 2
	length := len(arr)
	if length == 0 {
		return
	}
	if length < k {
		return
	}
	max := arr[0]
	for i := 0; i < length-k+1; i++ {
		tem := arr[i]
		for j := i + 1; j < i+k; j++ {
			fmt.Println("i, j, tem: ", i, j, tem)
			tem += arr[j]
		}
		fmt.Println("tem: ", tem)
		if tem > max {
			max = tem
		}
	}
	fmt.Println(float64(max) / float64(k))
}

func TestMaxAverage(t *testing.T) {
	//sliding window
	//适用于左右index都需要变化的情况，转变N^2 成N
	arr := []int{1, 2, 4, 5}
	k := 2
	length := len(arr)
	if length == 0 {
		return
	}
	if length < k {
		return
	}
	l := 0
	r := k - 1
	sum := 0
	for i := l; i <= r; i++ {
		sum += arr[i]
	}
	max := sum
	r++
	for r < length {
		fmt.Println(l, r)
		sum += arr[r]
		sum -= arr[l]
		if sum > max {
			max = sum
		}
		l++
		r++
	}
	fmt.Println(float64(max) / float64(k))
}

func TestSup(t *testing.T) {
	arr := []int{1, 3, 1, 4}
	mymap := make(map[int][]int)
	for i, v := range arr {
		if _, ok := mymap[v]; !ok {
			mymap[v] = make([]int, 2)
		}
		mymap[v][0] += 1
		if mymap[v][0] > 1 {
			fmt.Println(mymap[v][1], i)
			return
		}
		mymap[v][1] = i
	}
}

//find the longest same charactor sub string
//e.g: input aabbbaaaa--> output: a,4
func TestLongSub_sliding_window(t *testing.T) {
	str := "aabbbbbbaaaaa"
	arr := []rune(str)
	fmt.Println(string(arr))
	length := len(arr)
	max := 0
	mylen := 0
	l := 0
	r := 0
	for r < length {
		fmt.Println(r, l)
		if arr[r] == arr[l] {
			r++
		} else {
			mylen = r - l
			if mylen > max {
				max = mylen
			}

			l = r
		}
		if r == length-1 {
			mylen = r - l + 1
			if mylen > max {
				max = mylen
			}
		}
	}
	fmt.Println(max)
}
