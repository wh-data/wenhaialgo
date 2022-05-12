package merge_two_sorted_arr

//precondition: array are sorted
func merge(arr1, arr2 []int) []int {
	if len(arr1) == 0 {
		return deduplicate(arr2)
	}
	if len(arr2) == 0 {
		return deduplicate(arr1)
	}
	//two pointer to traverse
	i, j := 0, 0
	l1 := len(arr1)
	l2 := len(arr2)
	arr3 := []int{}
	if arr1[0] > arr2[0] {
		arr3 = append(arr3, arr2[0])
	} else {
		arr3 = append(arr3, arr1[0])
	}
	for i < l1 && j < l2 {
		if arr1[i] > arr2[j] {
			if arr3[len(arr3)-1] != arr2[j] {
				arr3 = append(arr3, arr2[j])
			}
			j++
			continue
		}
		if arr1[i] < arr2[j] {
			if arr3[len(arr3)-1] != arr1[i] {
				arr3 = append(arr3, arr2[i])
			}
			i++
			continue
		}
		if arr1[i] == arr2[j] {
			if arr3[len(arr3)-1] != arr1[i] {
				arr3 = append(arr3, arr2[j])
				j++
				continue
			}
			i++
			j++
		}
	}
	//check the left part of the arr
	if i < l1 {
		for _, a := range arr1 {
			if arr3[len(arr3)-1] != a {
				arr3 = append(arr3, a)
			}
		}
	}
	if j < l2 {
		for _, a := range arr2 {
			if arr3[len(arr3)-1] != a {
				arr3 = append(arr3, a)
			}
		}
	}
	return arr3
}

func deduplicate(arr []int) []int {
	newarr := []int{}
	for _, a := range arr {
		if len(newarr) == 0 {
			newarr = append(newarr, a)
		} else if newarr[len(newarr)-1] != a {
			newarr = append(newarr, a)
		}
	}
	return newarr
}
