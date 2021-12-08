package dynamic_programming

//1. build a matrix to contain all palindrome
//2. compare and find out the longest palindrome
func GetPalindrome(str string) (dp [][]bool, palindrome string) {
	dp = make([][]bool, len(str))
	longest := 0
	for i := range dp {
		dp[i] = make([]bool, len(str))
	}
	for i := range str {
		for j := i; j < len(str); j++ {
			if isPalindrome(str[i : j+1]) {
				dp[i][j] = true
				if j+1-i > longest {
					longest = j + 1 - i
					palindrome = str[i : j+1]
				}
			}
		}
	}
	return
}

//dp concept: for every palindrome, remove head and tail, it is still palindrome.
func isPalindrome(str string) bool {
	if len(str) == 1 || len(str) == 0 {
		return true
	}
	if str[0:1] != str[len(str)-1:] {
		return false
	}
	return isPalindrome(str[1 : len(str)-1])
}
