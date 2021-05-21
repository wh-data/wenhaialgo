package wenhaialgo

//time complexity O(nlogn)
//space complexity O(1), in Swap need create a new number
func HeapSort(arr []int, asc bool) {
	myheap := NewHeap()
	myheap.Elements = arr
	if len(myheap.Elements) < 1 {
		return
	}
	if asc {
		myheap.MaxHeapBuild()
	} else {
		myheap.MinHeapBuild()
	}
	for myheap.Size > 2 {
		Swap(myheap.Elements, 0, myheap.Size-1)
		myheap.Size--
		if asc {
			myheap.MaxHeapAdjustRoot()
		} else {
			myheap.MinHeapAdjustRoot()
		}
	}
	arr = myheap.Elements
}

//time complexity:average O(nlogn)
//space complexity: O(1)
func QuickSort(arr []int, asc bool) {
	if asc {
		quickSortByAsc(arr, 0, len(arr)-1)
	} else {
		quickSortByDec(arr, 0, len(arr)-1)
	}
}

func quickSortByAsc(arr []int, l, r int) {
	pivot := quickSortDivideAsc(arr, l, r)
	//left side
	if pivot > l {
		quickSortByAsc(arr, l, pivot-1)
	}
	//right side
	if pivot < r {
		quickSortByAsc(arr, pivot+1, r)
	}
}

func quickSortByDec(arr []int, l, r int) {
	pivot := quickSortDivideDec(arr, l, r)
	//left side
	if pivot > l {
		quickSortByDec(arr, l, pivot-1)
	}
	//right side
	if pivot < r {
		quickSortByDec(arr, pivot+1, r)
	}
}

func quickSortDivideAsc(arr []int, l, r int) int {
	tem := l
	l++
	for l <= r {
		if arr[l] >= arr[tem] {
			for l <= r {
				if arr[r] < arr[tem] {
					Swap(arr, l, r)
					l++
					r--
					break
				} else {
					r--
				}
			}
		} else {
			l++
		}
	}
	//r is the smaller one, r can smaller than l
	Swap(arr, tem, r)
	return r
}

func quickSortDivideDec(arr []int, l, r int) int {
	tem := l
	l++
	for l <= r {
		if arr[l] <= arr[tem] {
			for l <= r {
				if arr[r] > arr[tem] {
					Swap(arr, l, r)
					r--
					l++
					break
				} else {
					r--
				}
			}
		} else {
			l++
		}
	}
	Swap(arr, tem, r)
	return r
}

//time complexity: average O(n^2)
//space complexity: O(1)
func BubbleSort(arr []int, asc bool) {
	if len(arr) < 2 {
		return
	}
	length := len(arr)
	for length > 0 {
		i := 0
		for j := 1; j < length; j++ {
			if asc {
				if arr[j] < arr[i] {
					Swap(arr, j, i)
					i = j
				} else {
					i++
				}
			} else {
				if arr[j] > arr[i] {
					Swap(arr, j, i)
					i = j
				} else {
					i++
				}
			}
		}
		length--
	}
}

//time complexity O(N)~O(N^2)
//space complexity O(1)
func InsertSort(arr []int, asc bool) {
	if asc {
		insertSortAsc(arr)
	} else {
		insertSortDec(arr)
	}
}

func insertSortAsc(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		count := i + 1
		for count > 0 {
			if arr[count] < arr[count-1] {
				Swap(arr, count, count-1)
				count--
			} else {
				break
			}
		}
	}
}

func insertSortDec(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		count := i + 1
		for count > 0 {
			if arr[count] > arr[count-1] {
				Swap(arr, count, count-1)
				count--
			} else {
				break
			}
		}
	}
}

//time complexity O(N+N*(logN-logM))
//space complexity O(max - min ), max and min are the number in the array
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
