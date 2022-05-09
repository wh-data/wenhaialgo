package wenhaialgo

import (
	"fmt"
	"math"
)

type MyHeap struct {
	Elements []int
	Size     int
	Inited   bool
}

func NewHeap() *MyHeap {
	return &MyHeap{Inited: true}
}

func (h *MyHeap) MinHeadShiftUp(index int) {
	if float64(index-1)/2 >= 0 {
		if h.Elements[(index-1)/2] > h.Elements[index] {
			Swap(h.Elements, (index-1)/2, index)
			h.MinHeadShiftUp((index - 1) / 2)
		}
	}
}

//recursion to all leaf node
func (h *MyHeap) MinHeapShiftDown(i int) {
	//if no node
	if 2*i+1 >= h.Size {
		return
	}
	//only left node
	if 2*i+2 == h.Size {
		if h.Elements[i] > h.Elements[2*i+1] {
			Swap(h.Elements, i, 2*i+1)
		}
		return
	}
	//if have both left and right node
	if 2*i+2 < h.Size {
		//left node smaller
		if h.Elements[2*i+1] <= h.Elements[2*i+2] {
			if h.Elements[i] > h.Elements[2*i+1] {
				Swap(h.Elements, i, 2*i+1)
				h.MinHeapShiftDown(2*i + 1)
			}
		} else { //right node bigger
			if h.Elements[i] > h.Elements[2*i+2] {
				Swap(h.Elements, i, 2*i+2)
				h.MinHeapShiftDown(2*i + 2)
			}
		}
	}
}

func (h *MyHeap) MaxHeadShiftUp(index int) {
	if float64(index-1)/2 >= 0 {
		if h.Elements[(index-1)/2] < h.Elements[index] {
			Swap(h.Elements, (index-1)/2, index)
			h.MaxHeadShiftUp((index - 1) / 2)
		}
	}
}

//recursion to all leaf node
func (h *MyHeap) MaxHeapShiftDown(i int) {
	//if no node
	if 2*i+1 >= h.Size {
		return
	}
	//only left node
	if 2*i+2 == h.Size {
		if h.Elements[i] < h.Elements[2*i+1] {
			Swap(h.Elements, i, 2*i+1)
		}
		return
	}
	//if have both left and right node
	if 2*i+2 < h.Size {
		//left node bigger
		if h.Elements[2*i+1] >= h.Elements[2*i+2] {
			if h.Elements[i] < h.Elements[2*i+1] {
				Swap(h.Elements, i, 2*i+1)
				h.MaxHeapShiftDown(2*i + 1)
			}
		} else { //right node bigger
			if h.Elements[i] < h.Elements[2*i+2] {
				Swap(h.Elements, i, 2*i+2)
				h.MaxHeapShiftDown(2*i + 2)
			}
		}
	}
}

func (h *MyHeap) MinHeapInsert(ele int) {
	//i is the index of array
	i := h.Size
	h.Elements = append(h.Elements, ele)
	h.Size++
	h.MinHeadShiftUp(i)
	return
}

func (h *MyHeap) MaxHeapInsert(ele int) {
	//i is the index of array
	i := h.Size
	h.Elements = append(h.Elements, ele)
	h.Size++
	h.MaxHeadShiftUp(i)
	return
}

//time complexity O(n), each level in the heap need different time
//space complexity O(1), in Swap need create a new number
func (h *MyHeap) MinHeapBuild() (err error) {
	if len(h.Elements) < 1 {
		return
	}
	h.Size = len(h.Elements)
	//start swap form the last parent node
	index := h.Size/2 - 1
	for index >= 0 {
		h.MinHeapShiftDown(index)
		index--
	}
	return
}

func (h *MyHeap) MaxHeapBuild() (err error) {
	if len(h.Elements) < 1 {
		return
	}
	h.Size = len(h.Elements)
	//start swap form the last parent node
	index := h.Size/2 - 1
	for index >= 0 {
		h.MaxHeapShiftDown(index)
		index--
	}
	return
}

func (h *MyHeap) MinHeapDeleteMin() (min int, err error) {
	if len(h.Elements) < 1 {
		return math.MinInt64, fmt.Errorf("no ele in the heap")
	}
	min = h.Elements[0]
	Swap(h.Elements, 0, h.Size-1)
	h.Size--
	h.Elements = h.Elements[0:h.Size]
	h.MinHeapAdjustRoot()
	return
}

func (h *MyHeap) MaxHeapDeleteMax() (max int, err error) {
	if len(h.Elements) < 1 {
		return math.MaxInt64, fmt.Errorf("no ele in the heap")
	}
	max = h.Elements[0]
	Swap(h.Elements, 0, h.Size-1)
	h.Size--
	h.Elements = h.Elements[0:h.Size]
	h.MaxHeapAdjustRoot()
	return
}

//MinHeapAdjustRoot is used only after root node is changed, other  node still have heap rules
//the scenario is in heap rank, after exchange last element with root
//when to update the size when to shift down
func (h *MyHeap) MinHeapAdjustRoot() {
	if len(h.Elements) < 2 {
		return
	}
	h.MinHeapShiftDown(0)
	return
}

//MinHeapAdjustRoot is used only after root node is changed, other  node still have heap rules
//the scenario is in heap rank, after exchange last element with root
//when to update the size when to shift down
func (h *MyHeap) MaxHeapAdjustRoot() {
	if len(h.Elements) < 2 {
		return
	}
	h.MaxHeapShiftDown(0)
	return
}

//more general shift up
//start from the last parent node:len(array)/2 - 1
func shiftUpBigVal(root, parent int, array []int) {
	if parent < root {
		return
	}
	leftKid := 2*parent + (1 - root)
	rightKid := 2*parent + (2 - root)
	//shift up left
	if leftKid < len(array) {
		if array[leftKid] > array[parent] {
			array[parent], array[leftKid] = array[leftKid], array[parent]
		}
	}
	//shift up right
	if rightKid < len(array) {
		if array[rightKid] > array[parent] {
			array[parent], array[rightKid] = array[rightKid], array[parent]
		}
	}
	parent--
	shiftUpBigVal(root, parent, array)
}
func shiftUpBigValV2(parent, length int, array []int) {
	if parent < 0 {
		return
	}
	leftKid := 2*parent + 1
	rightKid := 2*parent + 2
	//shift up left
	if leftKid < length {
		if array[leftKid] > array[parent] {
			array[parent], array[leftKid] = array[leftKid], array[parent]
		}
	}
	//shift up right
	if rightKid < length {
		if array[rightKid] > array[parent] {
			array[parent], array[rightKid] = array[rightKid], array[parent]
		}
	}
	parent--
	shiftUpBigValV2(parent, length, array)
}

func shiftUpSmallVal(root, parent int, array []int) {
	if parent < root {
		return
	}
	leftKid := 2*parent + (1 - root)
	rightKid := 2*parent + (2 - root)
	if leftKid < len(array) {
		if array[leftKid] < array[parent] {
			array[parent], array[leftKid] = array[leftKid], array[parent]
		}
	}
	if rightKid < len(array) {
		if array[rightKid] < array[parent] {
			array[parent], array[rightKid] = array[rightKid], array[parent]
		}
	}
	parent--
	shiftUpSmallVal(root, parent, array)
}

func shiftUpSmallValV2(parent, length int, array []int) {
	if parent < 0 {
		return
	}
	leftKid := 2*parent + 1
	rightKid := 2*parent + 2
	//shift up left
	if leftKid < length {
		if array[leftKid] < array[parent] {
			array[parent], array[leftKid] = array[leftKid], array[parent]
		}
	}
	//shift up right
	if rightKid < length {
		if array[rightKid] < array[parent] {
			array[parent], array[rightKid] = array[rightKid], array[parent]
		}
	}
	parent--
	shiftUpSmallValV2(parent, length, array)
}

//time complexity is log(n), log(n)是数的层数，重新调整数时候只需要调整有变动的那个分支
func ShiftBigValV2_forloop(arr []int, length int) {
	depth := length/2 - 1
	for i := depth; i >= 0; i-- {
		leftC := 2*i + 1
		rightC := 2*i + 2
		max := i
		if leftC <= length-1 && arr[leftC] > arr[max] {
			max = leftC
		}
		if rightC <= length-1 && arr[rightC] > arr[max] {
			max = rightC
		}
		if max != i {
			arr[i], arr[max] = arr[max], arr[i]
		}
	}
}
