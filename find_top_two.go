package wenhaialgo

import "math"

//using one for loop
//without changing ori index
func FindTopTwoMinimumNumber(array []int) (minimum, second, indexm, indexs int) {
	minimum = array[0]
	second = math.MaxInt32 //have to use max, to make sure "else if" logic have to triggered
	indexm = 0
	indexs = 1
	for i := 1; i < len(array); i++ {
		if array[i] < minimum {
			second = minimum
			minimum = array[i]
			indexs = indexm
			indexm = i
		} else if array[i] < second {
			second = array[i]
			indexs = i
		}
	}
	return
}

func FindTopTwoMaxNumber(array []int) (max, second, indexm, indexs int) {
	if len(array) < 2 {
		return
	}
	max = array[0]
	second = math.MinInt32 //have to use min, to make sure second number is compared
	indexm = 0
	indexs = 1
	for i := 1; i < len(array); i++ {
		if array[i] > max {
			second = max
			max = array[i]
			indexs = indexm
			indexm = i
		} else if array[i] > second {
			second = array[i]
			indexs = i
		}
	}
	return
}

func FindTopTwoMinNumberV2(array []int) (min, second, indexm, indexs int) {
	if len(array) < 2 {
		return
	}
	min = array[0]
	second = math.MaxInt32 //have to use max, to make sure second number is compared
	indexm = 0
	indexs = 0
	for i := 1; i < len(array); i++ {
		if array[i] < min {
			min = array[i]
			indexm = i
		}
	}
	for i := 0; i < len(array); i++ {
		if array[i] != min && array[i] < second {
			second = array[i]
			indexs = i
		}
	}
	return
}

func FindTopTwoMaxNumberV2(array []int) (max, second, indexm, indexs int) {
	if len(array) < 2 {
		return
	}
	max = array[0]
	second = math.MinInt32 //have to use min, to make sure "else if" logic have to triggered
	indexm = 0
	indexs = 0
	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
			indexm = i
		}
	}
	for i := 0; i < len(array); i++ {
		if array[i] != max && array[i] > second {
			second = array[i]
			indexs = i
		}
	}
	return
}

func FindMaxNumberByHeap(array []int) (max int) {
	if len(array) < 0 {
		return 0
	}
	lastParent := len(array)/2 - 1
	shiftUpBigVal(0, lastParent, array)
	return array[0]
}
