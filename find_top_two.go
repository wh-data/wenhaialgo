package wenhaialgo

func FindTopTwoMinimumNumber(array []int)(minimum, second, indexm, indexs int){
	minimum = array[0]
	second = array[1]
	indexm= 0
	indexs = 1
	for i := 1; i<len(array); i++ {
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

func FindTopTwoMaxNumber(array []int)(max, second, indexm, indexs int){
	max = array[0]
	second = array[1]
	indexm= 0
	indexs = 1
	for i := 1; i<len(array); i++ {
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
