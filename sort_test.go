package wenhaialgo

import (
	"fmt"
	"testing"
)

func TestSortByHeap(t *testing.T) {
	arr := []int{8, 9, 5, 2, 1, 4, 7}
	HeapSort(arr, false)
	fmt.Println(arr)
	HeapSort(arr, true)
	fmt.Println(arr)
	arr = []int{8, 9, 5, 2, 1, 4, 7}
	HeapSort(arr, true)
	fmt.Println(arr)
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
	quickSortDivideAsc(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func TestBubbleSortAsc(t *testing.T) {
	arr := []int{8, 9, 5, 2, 1, 4, 7}
	BubbleSort(arr, true)
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
