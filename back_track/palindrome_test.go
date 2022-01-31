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

func TestGetLongestPalindromeByDP(t *testing.T) {
	s := "abaab"
	getLongestPalindromeByDP(s)
}

func TestGetLongestPalindromeByCentralExpending(t *testing.T) {
	s := "abbac"
	fmt.Println(getLongestPalindromeByCentralExpending(s))
}

func TestXgetLongestPalindromeByManacher(t *testing.T) {
	s := "abbac"
	fmt.Println(getLongestPalindromeByManacher(s))
}

func TestXgetLongestPalindromeByBackTrack(t *testing.T) {
	s := "abbac"
	getLongestPalindromeByBruteForce(s)
}
