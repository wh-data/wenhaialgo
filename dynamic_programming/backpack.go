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

/*画表格分析会简单很多

        1	2	3	4	5	6	7	8	9	10	11	12	13	14	15
{5, 12}	0	0	0	0	12	12	12	12	12	12	12	12	12	12	12
{4, 3}	0	0	0	3	12	12	12	12	15	15	15	15	15	15	15
{7, 10}	0	0	0	3	12	12	12	12	15	15	15	22	22	22	22
{2, 3}	0	3	3	3	12	12	15	15	15	15	18	22	22	25	25
{6, 6}	0	3	3	3	12	12	15	15	15	15	18	22	22	25	25
————————————————
版权声明：本文为CSDN博主「雪zi」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/LemonGirls/article/details/83378980
*/

//precondition: w length equal v length
func back_pack_v3(w, v []int, c int) int {
	n := len(w)
	if n != len(v) || n == 0 {
		return 0
	}
	//init 2 dimension arr with 0
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	//init first line
	for j := w[0]; j < n; j++ {
		dp[0][j] = v[0]
	}
	//可顺序放，概念是如果放比较某一个容量情况下放能增值还是不放（放之前的case）值比较大
	for i := 1; i < n; i++ {
		for j := 0; j < c; j++ {
			//dp[i-1][j-w[i]]这个是理解的关键，表示当放第i个物品时，j-w[i]的空间时候的最大值加上第i个物品的值
			//例如总空间是15，当我放质量为7的物品时候，最大的价值是：（15-7=8）的空间的时候的最大价值加上本物品的价值
			dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-1][j-w[i]]+v[i])))
		}
	}
	return dp[n-1][c]
}
