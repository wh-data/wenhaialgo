package dynamic_programming

import (
	"fmt"
	"math"
	"sort"
)

// Greedy: is not always the best solution
// e.g w := []int{1, 2, 5, 6, 7}, v := []int{1, 6, 18, 22, 28}, c := 11
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

func BackpackIssueByDP(wt, val []int, capacity int) (highestVal int) {
	//二维数组解释
	/*
		i表示前i个物品,对应背包容量从0到capacity的最优解
		j表示当背包容量为j时候，对应从0到n个物品的最优解
	*/

	//对前i个item开始分析
	/*
		有3中情况：
		1.这个物品重量大于背包的最大容量，则这个物品永远不可能放进背包，则最优解就是不放这个物品
		2.当可以放进去的时候，有两种情况：选择其中最优解释最大的
			2.1. 放进去后前i个的最优解变小，
			2.2. 放进去后前i个的最优解变大
	*/

	//实例解释
	/*
		假设capacity = j；尝试前i个物品的最优解：
		1. wt[i]>capacity: 第i个物品永远不可能放进去，最优解就是继续保留j的容量，存i-1个物品的最优解；
		2. wt[i]<=capacity:前i个物品的最优解是两种可能中值较大的一个
			2.1. 放的最优解：val[i] + 容量减少wt[i]后对前i-1个物品的最优解
			2.2. 不放的最优解：保留容量j，对前i-1个物品的最优解
	*/

	itemCount := len(wt)
	//exceptedItems = []int{}
	dp := make([][]int, itemCount+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}
	newwt, newval := []int{}, []int{}
	newwt = append(newwt, 0)
	newwt = append(newwt, wt...)
	newval = append(newval, 0)
	newval = append(newval, val...)
	//初始化：当前0个物品（没有放物品）时候，无论背包容量多少，价值都是0
	for i := 0; i <= len(wt); i++ {
		dp[i][0] = 0
	}
	//初始化：当背包容量为0，无论到前几个物品，价值都是0
	for j := 0; j <= capacity; j++ {
		dp[0][j] = 0
	}
	for i := 1; i <= itemCount; i++ {
		for j := 1; j <= capacity; j++ {
			if newwt[i] > j { //如果当前物品重量大于当前背包的容量，最优解肯定是不放该物品
				dp[i][j] = dp[i-1][j]
				//exceptedItems = append(exceptedItems,i)
			} else {
				l := float64(dp[i-1][j])
				r := float64(dp[i-1][j-newwt[i]] + newval[i])
				dp[i][j] = int(math.Max(l, r))
			}
		}
	}
	return dp[itemCount][capacity]
}
