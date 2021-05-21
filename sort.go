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
