package dynamic_programming

import (
	"fmt"
	"sort"
)

//青蛙上楼梯，可以上一节台阶或2节台阶，到n节台阶有几种上法
func FragGoStair(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return FragGoStair(n-1) + FragGoStair(n-2)
}

//var FragGoStairPath [][]int

//青蛙上楼梯，可以上一节台阶或2节台阶，打印出所有可能性
func FragGoStairPossibility(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}
	if n == 2 {
		return [][]int{{1, 1}, {2}}
	}
	arr1 := FragGoStairPossibility(n - 1)
	arr2 := FragGoStairPossibility(n - 2)
	for i := range arr1 {
		arr1[i] = append(arr1[i], 1)
	}
	for i := range arr2 {
		arr2[i] = append(arr2[i], 2)
	}
	return append(arr2, arr1...)
}

type Article struct {
	Weight  int
	Value   int
	VwRatio float32
}

//Greedy: is not always the best solution
//e.g w := []int{1, 2, 5, 6, 7}, v := []int{1, 6, 18, 22, 28}, c := 11
func BackpackIssueByGreedyAlgo(capacity int, weight, value []int) (chosenWeight, chosenValue []int, totalWeight, totalValue int) {
	if len(value) != len(weight) {
		return nil, nil, 0, 0
	}
	//cal Vi/Wi
	var articleArr []Article
	for i := 0; i < len(weight); i++ {
		articleArr = append(articleArr, Article{
			Weight:  weight[i],
			Value:   value[i],
			VwRatio: float32(value[i]) / float32(weight[i]),
		})
	}
	//sort
	sort.Slice(articleArr[:], func(i, j int) bool {
		return articleArr[i].VwRatio > articleArr[j].VwRatio
	})
	//greedy algo
	for i := 0; i < len(weight); i++ {
		chosenWeight = append(chosenWeight, articleArr[i].Weight)
		chosenValue = append(chosenValue, articleArr[i].Value)
		totalWeight += articleArr[i].Weight
		totalValue += articleArr[i].Value
		if totalWeight > capacity {
			chosenWeight = chosenWeight[:len(chosenWeight)-1]
			chosenValue = chosenValue[:len(chosenValue)-1]
			totalWeight -= articleArr[i].Weight
			totalValue -= articleArr[i].Value
			break
		}
	}
	return
}

//back tracking
var resTotalWeight, resTotalValue, curTotalWeight, curTotalValue int
var resWeightList, resValueList, curWeightList, curValueList []int
var checkedMap = make(map[int]bool, 0)

var count int

//todo: fix bug
func BackpackIssueByBackTrackAlgo(articleIndex, capacity int, articles []Article) {
	if articleIndex >= len(articles) {
		if curTotalValue > resTotalValue {
			resTotalValue = curTotalValue
			resTotalWeight = curTotalWeight
			resWeightList = curWeightList
			resValueList = curValueList
		}
		fmt.Println("debugg: added all")
		return
	}
	for i, art := range articles {
		//do not repeat same index in each path
		if checkedMap[i] {
			continue
		}
		curTotalWeight += art.Weight
		curTotalValue += art.Value
		curWeightList = append(curWeightList, art.Weight)
		curValueList = append(curValueList, art.Value)
		checkedMap[i] = true
		//exceed capacity, return
		if curTotalWeight > capacity {
			//check whether the cur value is the optimised one
			if curTotalValue-art.Value > resTotalValue {
				resTotalValue = curTotalValue - art.Weight
				resTotalWeight = curTotalWeight - art.Value
				resWeightList = curWeightList[:len(curWeightList)-1]
				resValueList = curValueList[:len(curValueList)-1]
				fmt.Println("debugg: resWeightList", resWeightList, "resValueList", resValueList)
			}
			//roll back data, but still same level
			curTotalWeight -= art.Weight
			curTotalValue -= art.Value
			curWeightList = curWeightList[:len(curWeightList)-1]
			curValueList = curValueList[:len(curValueList)-1]
			checkedMap[i] = false
			//if all ele are checked, return to upper level
			if i == len(articles) {
				articleIndex--
				return
			}
			continue
		}
		BackpackIssueByBackTrackAlgo(articleIndex+1, capacity, articles)
		//roll back to upper level
		curTotalWeight -= art.Weight
		curTotalValue -= art.Value
		curWeightList = curWeightList[:len(curWeightList)-1]
		curValueList = curValueList[:len(curValueList)-1]
		checkedMap[i] = false
		articleIndex--
	}
}

//dynamic programming
//refer https://www.jianshu.com/p/c289cd8ae0ed
//todo: check the algo
func BackpackIssueByDP(nums [][]int, total int) int {
	dp := make([]int, total+1)
	for i := 0; i < len(nums); i++ {
		for j := nums[i][0]; j <= total; j++ {
			if dp[j-nums[i][0]+nums[i][1]] > dp[j] {
				dp[j] = dp[j-nums[i][0]+nums[i][1]]
			}
		}
	}
	return dp[total]
}
