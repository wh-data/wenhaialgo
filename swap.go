package wenhaialgo

func Swap(array []int, i, j int) {
	tem := array[i]
	array[i] = array[j]
	array[j] = tem
}
