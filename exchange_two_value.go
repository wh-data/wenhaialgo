package wenhaialgo

func exchange_int_v1(a, b int) (int, int) {
	//a,b= b,a, go support direct multiple variable assignment
	return b, a
}

func exchange_int_v2(a, b int) (int, int) {
	//e.g a=1, b=2
	a = a + b //a=3
	b = a - b //b=2
	a = a - b //a=1
	return a, b
}
