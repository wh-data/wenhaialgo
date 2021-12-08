package dynamic_programming

import (
	"fmt"
	"math"
	"sort"
)

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

type Backpack struct {
	ResTotalWeight int
	ResTotalValue  int
	ResWeightList  []int
	ResValueList   []int
	CurTotalWeight int
	CurTotalValue  int
	CurWeightList  []int
	CurValueList   []int
}

var checkedMap = make(map[int]bool, 0)

var count int

func BackpackIssueByBackTrackAlgo(articleIndex, capacity int, articles []Article, bp *Backpack) {
	if articleIndex >= len(articles) {
		if bp.CurTotalValue > bp.ResTotalValue {
			bp.ResTotalValue = bp.CurTotalValue
			bp.ResTotalWeight = bp.CurTotalWeight
			bp.ResWeightList = bp.CurWeightList
			bp.ResValueList = bp.CurValueList
		}
		fmt.Println("debugg: added all")
		return
	}
	for i, art := range articles {
		if !checkedMap[i] && bp.CurTotalWeight+art.Weight <= capacity {
			bp.CurTotalWeight += art.Weight
			bp.CurTotalValue += art.Value
			bp.CurWeightList = append(bp.CurWeightList, art.Weight)
			bp.CurValueList = append(bp.CurValueList, art.Value)
			checkedMap[i] = true
			//cal optimised res
			if bp.CurTotalValue > bp.ResTotalValue {
				bp.ResTotalValue = bp.CurTotalValue
				bp.ResTotalWeight = bp.CurTotalWeight
				//have to do deep copy, otherwise res will updated when cur updated
				bp.ResWeightList = make([]int, len(bp.CurWeightList))
				bp.ResValueList = make([]int, len(bp.CurValueList))
				copy(bp.ResWeightList, bp.CurWeightList)
				copy(bp.ResValueList, bp.CurValueList)
			}
			BackpackIssueByBackTrackAlgo(articleIndex+1, capacity, articles, bp)
			//roll back to upper level
			bp.CurTotalWeight -= art.Weight
			bp.CurTotalValue -= art.Value
			bp.CurWeightList = bp.CurWeightList[:len(bp.CurWeightList)-1]
			bp.CurValueList = bp.CurValueList[:len(bp.CurValueList)-1]
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
