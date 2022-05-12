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
