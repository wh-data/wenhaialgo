package wenhaialgo

import (
	"fmt"
	"testing"
)

func TestSortByHeap(t *testing.T) {
	//arr := []int{8, 9, 5, 2, 1, 4, 7, 23, 34, 33, 11, 80, 90, 50, 20, 10, 40, 70, 230, 340, 330, 110, 2300, 3400, 3300}
	//HeapSortV2_forloop(arr)
	//fmt.Println(arr)
	arr := []int{8, 9, 8, 8, 5, 10, 1}
	arr1 := make([]int, len(arr))
	copy(arr1, arr)
	fmt.Println("old arr:  ", arr)
	quickSortByDec(arr, 0, len(arr)-1)
	fmt.Println("new arr:  ", arr)
	fmt.Println("old arr1:  ", arr1)
	quickSortByAsc(arr1, 0, len(arr1)-1)
	fmt.Println("new arr1:  ", arr1)
	//arr = []int{8, 9, 5, 2, 1, 4, 7}
	//HeapSort(arr, true)
	//fmt.Println(arr)
	//arr = []int{8, 9, 5, 2, 1, 4, 7}
	//HeapSort(arr, false)
	//fmt.Println(arr)
	//arr = []int{8, 9, 5, 2, 1, 4, 7}
	//HeapSortV2(arr, true)
	//fmt.Println(arr)
	//arr = []int{8, 9, 5, 2, 1, 4, 7}
	//HeapSortV2(arr, false)
	//fmt.Println(arr)
}

func TestQuickSort(t *testing.T) {
	arr := []int{8, 9, 5, 2, 1, 4, 7}
	QuickSort(arr, false)
	fmt.Println(arr)
	arr = []int{8, 9, 5, 2, 1, 4, 7}
	QuickSort(arr, true)
	fmt.Println(arr)
}

func TestQuickSortDivide(t *testing.T) {
	arr := []int{4, 5, 7}
	quickSortByDec(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func TestBubbleSortAsc(t *testing.T) {
	arr := []int{8, 9, 5, 2, 1, 4, 7}
	BubbleSort(arr, true)
	fmt.Println(arr)
	arr = []int{8, 9, 5, 2, 1, 4, 7}
	BubbleSortV2(arr, true)
	fmt.Println(arr)
	arr = []int{8, 9, 5, 2, 1, 4, 7}
	GeneralSort(arr, true)
	fmt.Println(arr)
}

func TestInsertSort(t *testing.T) {
	arr := []int{8, 9, 5, 2, 1, 4, 7}
	InsertSort(arr, true)
	fmt.Println(arr)
	InsertSort(arr, false)
	fmt.Println(arr)
}

func TestBucketSort(t *testing.T) {
	arr := []int{8, 9, 5, 2, 1, 4, 7}
	BucketSort(arr, false, 3)
	fmt.Println(arr)
	arr = []int{8, 9, 5, 2, 1, 4, 7}
	BucketSort(arr, true, 3)
	fmt.Println(arr)
}
