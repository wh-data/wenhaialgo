package wenhaialgo

import (
	"fmt"
	"testing"
)

func TestFindMedianBySort(t *testing.T) {
	arr := []int{9, 5, 2, 1, 6, 7}
	median := FindMedianBySort(arr)
	fmt.Println(median)
}

func TestFindMedianByForLoop(t *testing.T) {
	arr := []int{9, 5, 2, 1, 6, 7}
	median := FindMedianByForLoop(arr)
	fmt.Println(median)
}

func TestFindMedianByHeap(t *testing.T) {
	arr := []int{9, 5, 2, 1, 6, 7}
	median := FindMedianByHeap(arr)
	fmt.Println(median)
}

func TestFindTopTwoMinimumNumber(t *testing.T) {
	//arr := []int{ 9,5, 2,1,6,7}
	arr := []int{-3, -2, -1}
	one, two, ione, itwo := FindTopTwoMinimumNumber(arr)
	fmt.Println(one, two, ione, itwo)
}

func TestFindTopTwoMaxNumber(t *testing.T) {
	//arr := []int{9, 5, 2, 1, 6, 7}
	arr := []int{-1, -2, -3}
	one, two, ione, itwo := FindTopTwoMaxNumber(arr)
	fmt.Println(one, two, ione, itwo)
}

func TestMyHeap_InitMinHeap(t *testing.T) {
	h := NewHeap()
	fmt.Println(h)
}

func TestMyHeap_Insert(t *testing.T) {
	arr := []int{9, 5, 2, 1, 6, 7}
	h := NewHeap()
	for _, v := range arr {
		h.MinHeapInsert(v)
		fmt.Println(h.Elements)
	}

}

func TestMyHeap_Delete(t *testing.T) {
	arr := []int{9, 5, 2, 1, 6, 7}
	var err error
	h := NewHeap()
	for _, v := range arr {
		h.MinHeapInsert(v)
	}
	fmt.Println(h.Elements, err)
	min, err := h.MinHeapDeleteMin()
	fmt.Println(min, err)
	fmt.Println(h.Elements)
}

func TestMyHeap_MinHeapBuild(t *testing.T) {
	arr := []int{8, 9, 5, 2, 1, 4, 7}
	var err error
	h := NewHeap()
	h.Elements = append(h.Elements, arr...)
	fmt.Println(h.Elements)
	err = h.MinHeapBuild()
	fmt.Println(h.Elements, err)
}
