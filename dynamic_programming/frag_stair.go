package dynamic_programming

import (
	"fmt"
	"math"
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
var resWeightList = make([]int, 0)
var resValueList = make([]int, 0)
var curWeightList = make([]int, 0)
var curValueList = make([]int, 0)
var checkedMap = make(map[int]bool, 0)

var count int

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
		if !checkedMap[i] && curTotalWeight+art.Weight <= capacity {
			curTotalWeight += art.Weight
			curTotalValue += art.Value
			curWeightList = append(curWeightList, art.Weight)
			curValueList = append(curValueList, art.Value)
			checkedMap[i] = true
			//cal optimised res
			if curTotalValue > resTotalValue {
				resTotalValue = curTotalValue
				resTotalWeight = curTotalWeight
				//have to do deep copy, otherwise res will updated when cur updated
				resWeightList = make([]int, len(curWeightList))
				resValueList = make([]int, len(curValueList))
				copy(resWeightList, curWeightList)
				copy(resValueList, curValueList)
			}
			BackpackIssueByBackTrackAlgo(articleIndex+1, capacity, articles)
			//roll back to upper level
			curTotalWeight -= art.Weight
			curTotalValue -= art.Value
			curWeightList = curWeightList[:len(curWeightList)-1]
			curValueList = curValueList[:len(curValueList)-1]
			checkedMap[i] = false
		}
	}
}

//dynamic programming
//refer https://www.jianshu.com/p/c289cd8ae0ed
//todo: check the algo
func BackpackIssueByDPV2(nums [][]int, total int) int {
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

//todo: bug fix
func BackpackIssueByDP(wt, val []int, size int) int {
	dp := make([][]int, len(wt))
	for i := range dp {
		dp[i] = make([]int, size)
	}
	for i := 0; i < len(wt); i++ {
		dp[i][0] = 0
	}
	for j := 0; j < size; j++ {
		dp[0][j] = 0
	}
	for i := 1; i <= len(wt); i++ {
		for j := 1; j <= size; j++ {
			if wt[i] > size {
				dp[i][size] = dp[i-1][size]
			} else {
				dp[i][size] = int(math.Max(float64(dp[i-1][size]), float64(dp[i-1][size-wt[i]]+val[i])))
			}
		}
	}
	return dp[len(wt)][size]
}
