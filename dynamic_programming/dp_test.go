package dynamic_programming

import (
	"fmt"
	"testing"
	"time"
)

func TestFragGoStair(t *testing.T) {
	fmt.Println(FragGoStair(3))
	fmt.Println(FragGoStair(4))
	fmt.Println(FragGoStair(5))
	fmt.Println(FragGoStair(6))
	fmt.Println(FragGoStairPossibility(4))
}

func TestBackpack(Z *testing.T) {
	w := []int{1, 2, 5, 6, 7}
	v := []int{1, 6, 18, 22, 28}
	art := []Article{
		{
			Weight: 1,
			Value:  1,
		},
		{
			Weight: 2,
			Value:  6,
		},
		{
			Weight: 5,
			Value:  18,
		},
		{
			Weight: 6,
			Value:  22,
		},
		{
			Weight: 7,
			Value:  28,
		},
	}
	c := 11
	bp := &Backpack{}
	fmt.Println(BackpackIssueByGreedyAlgo(c, w, v))
	fmt.Println("before algo ", bp)
	BackpackIssueByBackTrackAlgo(0, c, art, bp)
	fmt.Println("after back tracking ", bp)
}

func TestBackpackDP(Z *testing.T) {
	wt := []int{1, 2, 5, 6, 7}
	val := []int{1, 6, 18, 22, 28}
	c := 11
	BackpackIssueByDP(wt, val, c)
}

func TestBB(t *testing.T) {
	fmt.Println("td"[len("td")-1:])
	fmt.Println(string("test"[2]))
}

func TestIsPalindrome(t *testing.T) {
	strs := []string{"a", "aa", "abba"}
	for _, s := range strs {
		fmt.Println(isPalindrome(s))
	}
}

func TestGetLongestPalindrome(t *testing.T) {
	s := []string{"baab", "abcba", "aa", " | a", "a", "ttest", "asdfadsfasegewtragdasfgasdbghgaadsfasdfadsfasdfcas"}
	start := time.Now()
	for _, x := range s {
		_, p := GetPalindrome(x)
		fmt.Println(p)
	}
	fmt.Println(time.Now().Sub(start).Nanoseconds())
}

func TestNMPath(t *testing.T) {
	fmt.Println(n_m_path(4, 3))
}
