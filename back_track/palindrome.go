package back_track

import "fmt"

var myData = make(map[int]int)

func longestPalindrome(s string) string {
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		getLongestPalindrome(arr, i, -1, 0)
	}
	str := ""
	length := 0
	for i, j := range myData {
		if j-i+1 > length {
			//fmt.Println("debuggg : i : j+1 ", i, j+1)
			str = string(arr[i : j+1])
			length = j - i + 1
			//fmt.Println("str: ", str)
		}
	}
	return str
}

func getLongestPalindrome(arr []rune, index, indexRight, length int) {
	//fmt.Println("debuggg index: ", index, "length: ", length)
	if index+length > len(arr)-1 {
		return
	}
	if arr[index] == arr[index+length] {
		myData[index] = index + length
		indexRight = index + length
		getLongestPalindrome(arr, index, indexRight, length+1)
	}
	if index-length < 0 {
		return
	}
	if arr[index-length] == arr[index+length] {
		myData[index-length] = index + length
		getLongestPalindrome(arr, index, indexRight, length+1)
	}
	if indexRight > 0 {
		if arr[index-length] == arr[indexRight+length] {
			myData[index-length] = indexRight + length
			getLongestPalindrome(arr, index, indexRight, length+1)
		}
	}
}

func getLongestPalindromeByDP(s string) {
	if len(s) < 2 {
		return
	}
	//init
	dp := make([][]bool, len(s))
	for i, _ := range dp {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(dp); i++ {
		dp[i][i] = true
	}
	//have to put j to the outer for loop, purpose is to make dp[i+1][j-1] ready before use
	//先遍历列
	for j := 0; j < len(dp); j++ {
		for i := 0; i < j; i++ {
			dp[i][j] = s[j] == s[i] && (j-i < 2 || dp[i+1][j-1])
			if dp[i][j] {
			}
		}
	}
	//print
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp); j++ {
			fmt.Print(dp[i][j], " ")
		}
		fmt.Println()
	}
	return
}

//if use this iteration, time complexity will be n^3
func getDP(s string, dp [][]bool, i, j int) bool {
	if j <= i {
		return true
	}
	dp[i][j] = s[j] == s[i] && getDP(s, dp, i+1, j-1)
	return dp[i][j] == true
}

func getLongestPalindromeByCentralExpending(s string) (int, string) {
	radius := 0
	radiusDouble := 0
	central := ""
	//for 1 central
	for i, c := range s {
		for j := 1; j < len(s); j++ {
			left := i - j
			right := i + j
			if left < 0 || right >= len(s) {
				break
			}
			if s[left] == s[right] {
				if j > radius {
					radius = j
					central = string(c)
				}
				continue
			} else {
				break
			}
		}
	}
	//for 2 central chars
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			continue
		}
		c2 := s[i : i+2]
		fmt.Println(s[i:2])
		fmt.Println(c2)
		for j := 1; j < len(s); j++ {
			left := i - j
			right := i + 1 + j
			if left < 0 || right >= len(s) {
				break
			}
			if s[left] == s[right] {
				if j > radiusDouble {

					radiusDouble = j
					central = c2

				}
				continue
			} else {
				break
			}
		}
	}
	radiusFinal := 0
	if radius == 0 && radiusDouble == 0 {
		radiusFinal = 1
	} else {
		if radius*2+1 > radiusDouble*2+2 {
			radiusFinal = radius*2 + 1
		} else {
			radiusFinal = radiusDouble*2 + 2
		}
	}
	return radiusFinal, central
}
