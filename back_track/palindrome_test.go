package back_track

import (
	"fmt"
	"testing"
	"time"
)

func TestLongestPalindrome(t *testing.T) {
	s := []string{"baab", "abcba", "aa", " | a", "a", "ttest", "asdfadsfasegewtragdasfgasdbghgaadsfasdfadsfasdfcas"}
	start := time.Now()
	for _, input := range s {
		str := longestPalindrome(input)
		fmt.Println(str)
	}
	fmt.Println(time.Now().Sub(start).Nanoseconds())
}

func TestA(t *testing.T) {
	s := "tes"
	arr := []rune(s)
	//fmt.Println(arr[1:0])
	fmt.Println(arr[1:1])
	fmt.Println(arr[1:2])
	fmt.Println(arr[1:3])
}
