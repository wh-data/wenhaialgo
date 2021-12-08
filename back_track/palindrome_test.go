package back_track

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	s := []string{"baab", "abcba", "aa", " | a", "a", "ttest"}
	for _, input := range s {
		str := longestPalindrome(input)
		fmt.Println(str)
	}
}

func TestA(t *testing.T) {
	s := "tes"
	arr := []rune(s)
	//fmt.Println(arr[1:0])
	fmt.Println(arr[1:1])
	fmt.Println(arr[1:2])
	fmt.Println(arr[1:3])
}
