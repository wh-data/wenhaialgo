package shift_string

func shift_string(str string, shift int) string {
	new_runes := []rune(str)
	l := len(new_runes)
	new_shift := shift % l
	for new_shift > 0 {
		str = shift_one(str)
		new_shift--
	}
	return str
}

func shift_one(str string) string {
	new_runes := []rune(str)
	l := len(new_runes)
	if l == 0 {
		return ""
	}
	tem := new_runes[0]
	for i := 0; i < l-1; i++ {
		new_runes[i] = new_runes[i+1]
	}
	new_runes[l-1] = tem
	return string(new_runes)
}

func shift_string_v2(str string, shift int) string {
	new_runes := []rune(str)
	l := len(new_runes)
	new_shift := shift % l
	po := l - 1 - new_shift //new position of last char
	cut := (l - 1) - po - 1 //cut position
	return string(append(new_runes[cut+1:], new_runes[:cut+1]...))
}

func reverse_string(str string) string {
	new_runes := []rune(str)
	l := len(new_runes)
	i := 0
	j := l - 1
	for j > i {
		new_runes[i], new_runes[j] = new_runes[j], new_runes[i]
		i++
		j--
	}
	return string(new_runes)
}

//判断一个字符串是否可以组成另一个
func ransom_magzine(ra, mag string) bool {
	mymap := make(map[rune]int, 0)
	for _, r := range ra {
		if _, ok := mymap[r]; ok {
			mymap[r]++
		} else {
			mymap[r] = 1
		}
	}
	for _, m := range mag {
		if _, ok := mymap[m]; !ok {
			return false
		}
		mymap[m]--
		if mymap[m] == 0 {
			delete(mymap, m)
		}
	}
	return true
}
