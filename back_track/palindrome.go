package back_track

import "fmt"

var myData = make(map[int]int)

func longestPalindrome(s string) string {
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		getLongestPalindrome(arr, i, 0)
	}
	str := ""
	length := 0
	for i, j := range myData {
		if j-i+1 > length {
			fmt.Println("debuggg : i : j+1 ", i, j+1)
			str = string(arr[i : j+1])
			length = j - i + 1
			fmt.Println("str: ", str)
		}
	}
	return str
}

func getLongestPalindrome(arr []rune, index, length int) {
	fmt.Println("debuggg index: ", index, "length: ", length)
	if index+length > len(arr)-1 {
		return
	}
	if arr[index] == arr[index+length] {
		myData[index] = index + length
		getLongestPalindrome(arr, index, length+1)
	}
	if index-length < 0 {
		return
	}
	if arr[index-length] == arr[index+length] {
		myData[index-length] = index + length
		getLongestPalindrome(arr, index, length+1)
	}
}
