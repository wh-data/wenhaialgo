package play_with_arr

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestARR(t *testing.T) {
	arr := []int{1, 2}
	fmt.Println(FindIndex(arr, 2))
	fmt.Println(remove_ele(arr, 2))
}

func TestAadd_on(t *testing.T) {
	arr := []int{1, 2, 3, 4, 8}
	s := 9
	fmt.Println(add_on(arr, s))
}

func TestSSpiral_Matrix(t *testing.T) {
	n := 1
	m := Spiral_Matrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(m[j][i], ",")
		}
		fmt.Println()
	}
}

func TestMaxAverage_bad_solution(t *testing.T) {
	//bad solution
	arr := []int{1, 2, 4, 5}
	k := 2
	length := len(arr)
	if length == 0 {
		return
	}
	if length < k {
		return
	}
	max := arr[0]
	for i := 0; i < length-k+1; i++ {
		tem := arr[i]
		for j := i + 1; j < i+k; j++ {
			fmt.Println("i, j, tem: ", i, j, tem)
			tem += arr[j]
		}
		fmt.Println("tem: ", tem)
		if tem > max {
			max = tem
		}
	}
	fmt.Println(float64(max) / float64(k))
}

func TestMaxAverage(t *testing.T) {
	//sliding window
	//适用于左右index都需要变化的情况，转变N^2 成N
	arr := []int{1, 2, 4, 5}
	k := 2
	length := len(arr)
	if length == 0 {
		return
	}
	if length < k {
		return
	}
	l := 0
	r := k - 1
	sum := 0
	for i := l; i <= r; i++ {
		sum += arr[i]
	}
	max := sum
	r++
	for r < length {
		fmt.Println(l, r)
		sum += arr[r]
		sum -= arr[l]
		if sum > max {
			max = sum
		}
		l++
		r++
	}
	fmt.Println(float64(max) / float64(k))
}

func TestSup(t *testing.T) {
	arr := []int{1, 3, 1, 4}
	mymap := make(map[int][]int)
	for i, v := range arr {
		if _, ok := mymap[v]; !ok {
			mymap[v] = make([]int, 2)
		}
		mymap[v][0] += 1
		if mymap[v][0] > 1 {
			fmt.Println(mymap[v][1], i)
			return
		}
		mymap[v][1] = i
	}
}

//find the longest same charactor sub string
//e.g: input aabbbaaaa--> output: a,4
func TestLongSub_sliding_window(t *testing.T) {
	str := "aabbbbbbaaaaa"
	arr := []rune(str)
	fmt.Println(string(arr))
	length := len(arr)
	max := 0
	mylen := 0
	l := 0
	r := 0
	for r < length {
		fmt.Println(r, l)
		if arr[r] == arr[l] {
			r++
		} else {
			mylen = r - l
			if mylen > max {
				max = mylen
			}

			l = r
		}
		if r == length-1 {
			mylen = r - l + 1
			if mylen > max {
				max = mylen
			}
		}
	}
	fmt.Println(max)
}

//find sum of two numbs
func twoSum(nums []int, target int) []int {
	mymap := make(map[int]int, 0)
	length := len(nums)
	if length < 2 {
		return nil
	}
	for i := 0; i < length; i++ {
		if idx, ok := mymap[target-nums[i]]; ok {
			return []int{idx, i}
		}
		mymap[nums[i]] = i
	}
	return nil
}

func Test_findMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	//fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(findMedian(nums1, nums2))
}

func findMedian(nums1 []int, nums2 []int) float64 {
	combined := append(nums1, nums2...)
	sort.Ints(combined[:])

	var (
		median     float64
		idx1, idx2 int
	)
	if len(combined)%2 != 0 {
		idx1 = int(float64(len(combined)-1) * float64(0.5))
		median = float64(combined[idx1])
	} else {
		idx1 = int(float64(len(combined)-1) * float64(0.5))
		idx2 = int(float64(len(combined)) * float64(0.5))
		median = (float64(combined[idx1]) + float64(combined[idx2])) / float64(2)
	}

	return median
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	p1, p2 := 0, 0
	newnums := make([]int, 0)

	if len(nums1) == 0 {
		newnums = nums2
	} else if len(nums2) == 0 {
		newnums = nums1
	} else {
		v1 := nums1[0]
		v2 := nums2[0]
		tem := int(-1e6 - 1)
		for p1 < len(nums1) || p2 < len(nums2) {
			fmt.Println("p1: ", p1)
			fmt.Println("p2: ", p2)
			if p1 < len(nums1) {
				v1 = nums1[p1]
				if v1 == tem {
					p1++
					continue
				}
			} else {
				v1 = 1e6 + 1
			}
			if p2 < len(nums2) {
				v2 = nums2[p2]
				if v2 == tem {
					p2++
					continue
				}
			} else {
				v2 = 1e6 + 1
			}
			if v1 < v2 {
				newnums = append(newnums, v1)
				tem = v1
				if p1 < len(nums1) {
					p1++
				}
			} else if v2 < v1 {
				newnums = append(newnums, v2)
				tem = v2
				if p2 < len(nums2) {
					p2++
				}
			} else { //v1==v2
				newnums = append(newnums, v2)
				tem = v2
				if p1 < len(nums1) {
					p1++
				}
				if p2 < len(nums2) {
					p2++
				}
			}
		}
	}
	length := len(newnums)
	if length%2 == 1 {
		return float64(newnums[length/2])
	} else {
		return (float64(newnums[length/2-1]) + float64(newnums[length/2])) / 2
	}
}

//The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//And then read line by line: "PAHNAPLSIIGYIR"
//
//Write the code that will take a string and make this conversion given a number of rows:
//
//string convert(string s, int numRows);
//
//
//Example 1:
//
//Input: s = "PAYPALISHIRING", numRows = 3
//Output: "PAHNAPLSIIGYIR"
//Example 2:
//
//Input: s = "PAYPALISHIRING", numRows = 4
//Output: "PINALSIGYAHRPI"
//Explanation:
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
//Example 3:
//
//Input: s = "A", numRows = 1
//Output: "A"
//
//
//Constraints:
//
//1 <= s.length <= 1000
//s consists of English letters (lower-case and upper-case), ',' and '.'.
//1 <= numRows <= 1000

func Test_zigzag_convert(t *testing.T) {
	fmt.Println(zigzag_convert("ab", 1))
}
func zigzag_convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	metrix := make([][]rune, numRows)
	chars := []rune(s)

	for i, _ := range metrix {
		metrix[i] = make([]rune, 0)
	}

	f := true
	i := 0

	for _, v := range chars {
		fmt.Println(string(v), i)
		metrix[i] = append(metrix[i], v)
		flag(&i, f)
		if i == numRows {
			i -= 2
			f = false
		}
		if i == -1 {
			i += 2
			f = true
		}
	}

	res := ""
	for _, row := range metrix {
		res += string(row)
	}
	return res
}
func flag(i *int, f bool) {
	if f {
		*i++
	} else {
		*i--
	}

}

/*
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).



Example 1:

Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".


Constraints:

1 <= s.length <= 20
1 <= p.length <= 30
s contains only lowercase English letters.
p contains only lowercase English letters, '.', and '*'.
It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
*/

func TestIsMatch(t *testing.T) {
	s := "mississippi"
	p := "mis*is*ip*."
	fmt.Println(isMatch(s, p))
}

//todo: the algo is not finished, direction is correct
func isMatch(s string, p string) bool {
	arr_s := []rune(s)
	arr_p := []rune(p)
	if len(arr_p) == 0 {
		return len(arr_s) == 0
	}

	dp := make([][]bool, len(arr_s)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, len(arr_p)+1)
	}

	dp[0][0] = true
	for j := 1; j <= len(arr_p); j++ {
		dp[0][j] = dp[0][j-1] && string(arr_p[j-1]) == "*"
	}
	for i := 1; i <= len(arr_s); i++ {
		dp[i][0] = false
	}
	for j := 1; j <= len(arr_p); j++ {
		for i := 1; i <= len(arr_s); i++ {
			if arr_s[i-1] == arr_p[j-1] || string(arr_p[j-1]) == "." {
				dp[i][j] = dp[i-1][j-1]
			} else if string(arr_p[j-1]) == "*" {
				if i == 1 {
					dp[i][j] = i == 1 && string(arr_s[i-1]) == string(arr_p[j-2]) || string(arr_p[j-2]) == "."
				} else {
					dp[i][j] = dp[i-1][j-1] ||
						(dp[i-1][j] && (arr_s[i-1] == arr_s[i-2] || string(arr_p[j-2]) == "."))
				}
			}
		}
	}
	for i := 0; i <= len(arr_s); i++ {
		for j := 0; j <= len(arr_p); j++ {
			if dp[i][j] {
				fmt.Print("\033[32m", dp[i][j], " ")
			} else {
				fmt.Print("\033[31m", dp[i][j], " ")
			}
		}
		fmt.Println()
	}
	return dp[len(arr_s)][len(arr_p)]
}

/*Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral.



Example 1:

Input: num = 3
Output: "III"
Explanation: 3 is represented as 3 ones.
Example 2:

Input: num = 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.
Example 3:

Input: num = 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.


Constraints:

1 <= num <= 3999

*/
//1 <= num <= 3999
func TestIntToRoman(t *testing.T) {
	num := 3000
	fmt.Println(intToRoman(num))
	roman := "VI"
	fmt.Println(romanToInt(roman))
}

func intToRoman(num int) string {
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	alb := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res := ""
	for i, v := range alb {
		res += strings.Repeat(roman[i], num/v)
		num %= v
	}
	return res
}

func romanToInt(s string) int {
	romanMap := map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90, "L": 50,
		"XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}
	arr_ori := []rune(s)
	l, r := 0, 0
	res := 0
	for r < len(arr_ori) {
		if r+2 <= len(arr_ori) {
			if v, ok := romanMap[string(arr_ori[l:r+2])]; ok {
				res += v
				l += 2
				r += 2
				continue
			}
		}
		if v, ok := romanMap[string(arr_ori[l:r+1])]; ok {
			res += v
			l++
			r++
		}
	}
	return res
}

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"abg", "c", "adfs"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return strs[0]
	}
	p := 1
	base := strs[0]
	res := ""
	for p < length {
		res = ""
		rune_b := []rune(base)
		rune_c := []rune(strs[p])
		len_b := len(rune_b)
		len_c := len(rune_c)
		l, r := 0, 0
		for l < len_b && r < len_c {
			if rune_b[l] == rune_c[r] {
				res += string(rune_b[l])
				l++
				r++
				continue
			}
			break
		}
		base = res
		p++
	}
	return res
}
