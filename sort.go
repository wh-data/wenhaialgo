package wenhaialgo

import (
	"fmt"
)

// time complexity O(nlogn)
// space complexity O(1), in Swap need create a new number
func HeapSort(arr []int, asc bool) {
	if asc {
		for i := 0; i < len(arr); i++ {
			//shift start from the last parent node
			//the left part is sorted arr
			lastParent := (len(arr)+i)/2 - 1
			shiftUpSmallVal(i, lastParent, arr)
		}
	} else {
		for i := 0; i < len(arr); i++ {
			lastParent := (len(arr)+i)/2 - 1
			shiftUpBigVal(i, lastParent, arr)
		}
	}
}

func HeapSortV2(arr []int, asc bool) {
	length := len(arr)
	if asc {
		for i := 0; i < len(arr); i++ {
			lastParent := length/2 - 1
			shiftUpBigValV2(lastParent, length, arr)
			arr[0], arr[length-1] = arr[length-1], arr[0] //move biggest num to last
			length--
		}
	} else {
		for i := 0; i < len(arr); i++ {
			lastParent := length/2 - 1
			shiftUpSmallValV2(lastParent, length, arr)
			arr[0], arr[length-1] = arr[length-1], arr[0] //move smallest num to last
			length--
		}
	}
}

func HeapSortV2_forloop(arr []int) {
	length := len(arr)
	for i := length; i >= 0; i-- {
		ShiftBigValV2_forloop(arr, i)
		if i > 0 {
			arr[0], arr[i-1] = arr[i-1], arr[0]
		}
	}
}

// time complexity:average O(nlogn)
// space complexity: O(1)
func QuickSort(arr []int, asc bool) {
	if asc {
		quickSortByAsc(arr, 0, len(arr)-1)
	} else {
		quickSortByDec(arr, 0, len(arr)-1)
	}
}

func quickSortByAsc(arr []int, l, r int) {
	if l < 0 {
		return
	}
	if r >= len(arr) {
		return
	}
	if l >= r {
		return
	}
	ori_l := l
	ori_r := r
	p := arr[l]
	for l < r {
		l++
		if arr[l] > p {
			for l < r {
				if arr[r] <= p {
					arr[l], arr[r] = arr[r], arr[l]
					r--
					break
				}
				r--
			}
		}
	}
	if arr[r] < p {
		arr[ori_l], arr[r] = arr[r], arr[ori_l]
	}
	quickSortByAsc(arr, ori_l, r-1)
	quickSortByAsc(arr, r, ori_r)
}

func quickSortByDec(arr []int, l, r int) {
	if l < 0 {
		return
	}
	if r >= len(arr) {
		return
	}
	if l >= r {
		return
	}
	ori_l := l
	ori_r := r
	p := arr[l]
	for l < r {
		l++
		if arr[l] < p {
			for l < r {
				if arr[r] >= p {
					arr[l], arr[r] = arr[r], arr[l]
					r--
					break
				}
				r--
			}
		}
	}
	if arr[r] >= p {
		arr[ori_l], arr[r] = arr[r], arr[ori_l]
	}
	quickSortByDec(arr, ori_l, r-1)
	quickSortByDec(arr, r, ori_r)
}

// key point: bubble sort is to adjust adjacent elements, if no change, mean all are sorted. can stop further check
// time complexity: average O(n^2), can be simpler then GeneralSort in some case
// space complexity: O(1)
func BubbleSort(arr []int, asc bool) {
	//count := 0 //for checking time complexity
	if len(arr) < 2 {
		return
	}
	length := len(arr)
	for length > 0 {
		i := 0
		isChanged := false
		for j := 1; j < length; j++ {
			//count += 1
			if asc {
				if arr[j] < arr[i] {
					Swap(arr, j, i)
					i = j
					isChanged = true
				} else {
					i++
				}
			} else {
				if arr[j] > arr[i] {
					Swap(arr, j, i)
					i = j
					isChanged = true
				} else {
					i++
				}
			}
		}
		length--
		if !isChanged {
			break
		}
	}
	//fmt.Println(count)
}

func BubbleSortV2(arr []int, asc bool) {
	//count := 0 //for checking time complexity
	if len(arr) < 2 {
		return
	}
	length := len(arr)
	for length > 0 {
		i := 0
		j := 1
		isChanged := false
		for j < length {
			//count += 1
			if asc {
				if arr[j] < arr[i] {
					arr[i], arr[j] = arr[j], arr[i]
					i = j
					j++
					isChanged = true
				} else {
					i++
					j++
				}
			} else {
				if arr[j] > arr[i] {
					arr[i], arr[j] = arr[j], arr[i]
					i = j
					j++
					isChanged = true
				} else {
					i++
					j++
				}
			}
		}
		length--
		if !isChanged {
			break
		}
	}
	//fmt.Println(count)
}

func BubbleSortFloat32(arr []float32, asc bool) {
	//count := 0 //for checking time complexity
	if len(arr) < 2 {
		return
	}
	length := len(arr)
	for length > 0 {
		i := 0
		j := 1
		isChanged := false
		for j < length {
			//count += 1
			if asc {
				if arr[j] < arr[i] {
					arr[i], arr[j] = arr[j], arr[i]
					i = j
					j++
					isChanged = true
				} else {
					i++
					j++
				}
			} else {
				if arr[j] > arr[i] {
					arr[i], arr[j] = arr[j], arr[i]
					i = j
					j++
					isChanged = true
				} else {
					i++
					j++
				}
			}
		}
		length--
		if !isChanged {
			break
		}
	}
	//fmt.Println(count)
}

// time complexity: always O(n(n+1)/2)
func GeneralSort(arr []int, asc bool) {
	count := 0 //for checking time complexity
	if len(arr) < 2 {
		return
	}
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			count += 1
			if asc {
				if arr[j] < arr[i] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			} else {
				if arr[j] > arr[i] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
	}
	fmt.Println(count)
}

// key point: the left part is sorted, pick ele and compare with sorted part,
// as long as found a ele that satisfy condition, can stop further check
// time complexity O(N)~O(N(N-1)/2)
// space complexity O(1)
func InsertSort(arr []int, asc bool) {
	if asc {
		insertSortAsc(arr)
	} else {
		insertSortDec(arr)
	}
}

func insertSortAsc(arr []int) {
	times := 0 //for debug
	for i := 0; i < len(arr)-1; i++ {
		for count := i + 1; count > 0; count-- {
			times += 1
			if arr[count] < arr[count-1] {
				arr[count], arr[count-1] = arr[count-1], arr[count]
			} else {
				break
			}
		}
	}
	fmt.Println(times)
}

func insertSortDec(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for count := i + 1; count > 0; count-- {
			if arr[count] > arr[count-1] {
				arr[count], arr[count-1] = arr[count-1], arr[count]
			} else {
				break
			}
		}
	}
}

// time complexity O(N+N*(logN-logM))
// space complexity O(max - min ), max and min are the number in the array
func BucketSort(arr []int, asc bool, bucketSize int) {
	if bucketSize < 1 {
		bucketSize = 10
	}
	max := arr[0]
	min := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	bucketCount := (max-min)/bucketSize + 1
	buckets := make([][]int, bucketCount)
	for _, v := range arr {
		index := (v - min) / bucketSize
		buckets[index] = append(buckets[index], v)
	}
	for _, bucket := range buckets {
		QuickSort(bucket, asc)
	}
	arr = append(arr[:0])
	if asc {
		for _, bucket := range buckets {
			arr = append(arr, bucket...)
		}
	} else {
		index := len(buckets) - 1
		for index >= 0 {
			arr = append(arr, buckets[index]...)
			index--
		}
	}
}
