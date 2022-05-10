package dynamic_programming

//a panel with n columns and m rows
//go from left top to right bottom
//how many path can chose
func n_m_path(n, m int) int {
	if n == 0 || m == 0 {
		return 0
	}
	if n == 1 || m == 1 {
		return 1
	}
	return n_m_path(n, m-1) + n_m_path(n-1, m)
}
