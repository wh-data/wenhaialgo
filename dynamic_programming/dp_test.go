package dynamic_programming

import (
	"fmt"
	"testing"
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