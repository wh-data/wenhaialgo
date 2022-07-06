package play_with_list

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	result := removeNthFromEnd(head, 3)
	fmt.Println(result.Val)
	next := result.Next
	for next != nil {
		fmt.Println(next.Val)
		next = next.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tem_n := new(int)
	*tem_n = n - 1
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Next.Next == nil {
		if n == 1 {
			head.Next = nil
			return head
		} else if n == 2 {
			return head.Next
		}
	}
	next := head.Next
	recur(next, next.Next, tem_n)
	fmt.Println("*tem_n: ", *tem_n)
	if *tem_n == 0 {
		tail := head.Next.Next
		head.Next = tail
		return head
	} else if *tem_n == 1 {
		return head.Next
	}
	return head
}

func recur(pre, cur *ListNode, n *int) {
	if cur.Next != nil {
		pre_m := cur
		cur_m := cur.Next
		recur(pre_m, cur_m, n)
		if *n == 0 {
			pre.Next = cur.Next
			*n--
		} else {
			*n--
		}
	} else {
		if *n == 0 {
			pre.Next = nil
			*n--
		} else {
			*n--
		}
	}
}

func TestMergeTwoLists(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	list2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	res := mergeTwoLists(list1, list2)
	for res != nil {
		fmt.Print(res.Val, " ")
		res = res.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	cur1 := list1
	cur2 := list2
	res := &ListNode{Val: 0}
	tem := res
	if cur1 == nil {
		return cur2
	}
	if cur2 == nil {
		return cur1
	}

	for cur1 != nil || cur2 != nil {
		if cur1 != nil && cur2 != nil {
			if cur1.Val < cur2.Val {
				tem.Next = &ListNode{Val: cur1.Val}
				tem = tem.Next
				cur1 = cur1.Next
			} else {
				tem.Next = &ListNode{Val: cur2.Val}
				tem = tem.Next
				cur2 = cur2.Next
			}
		} else if cur1 == nil {
			tem.Next = cur2
			tem = tem.Next
			cur2 = nil
		} else {
			tem.Next = cur1
			tem = tem.Next
			cur1 = nil
		}

	}
	return res.Next
}

func TestGenerateParenthesis(t *testing.T) {
	n := 3
	res := generateParenthesis(n)
	for _, v := range res {
		fmt.Print(v, " | ")
	}
}

func generateParenthesis(n int) []string {
	res := recur_again(n)
	//deduplicate
	final := deduplicate(res)
	return final
}

func recur_again(n int) (strs []string) {
	if n == 1 {
		return []string{"()"}
	}
	//pattern 1
	hashMap1 := make(map[string]struct{})
	str1 := new(string)
	recurParentesis_pattern1(1, n, str1, hashMap1)
	//pattern 2
	strs = append(strs, *str1)
	for i := 1; i < n; i++ {
		res_le := recur_again(i)
		res_ri := recur_again(n - i)
		for j := 0; j < len(res_le); j++ {
			for k := 0; k < len(res_ri); k++ {
				strs = append(strs, res_le[j]+res_ri[k])
			}
		}
	}
	//pattern 3
	tem := 1
	for tem < n-1 {
		for i := 1; i <= n-tem-1; i++ {
			res_le := recur_again(i)
			res_ri := recur_again(n - tem - i)
			for j := 0; j < len(res_le); j++ {
				for k := 0; k < len(res_ri); k++ {
					strs = append(strs, strings.Repeat("(", tem)+res_le[j]+res_ri[k]+strings.Repeat(")", tem))
				}
			}
		}
		tem++
	}

	return
}

func deduplicate(input []string) []string {
	hashMap := make(map[string]struct{})
	final := make([]string, 0)
	for _, v := range input {
		if _, ok := hashMap[v]; !ok {
			hashMap[v] = struct{}{}
			final = append(final, v)
		}
	}
	return final
}

func recurParentesis_pattern2(n int) (str string) {
	for i := 1; i <= n; i++ {
		str += "()"
	}
	return
}

func recurParentesis_pattern1(step, n int, str *string, hashMap map[string]struct{}) {
	if step > n {
		return
	}
	//每个括号进来的可能性：（ ）
	//左括号：len（str）+ 1
	//右括号：位置由左括号唯一确定 + 1
	//for i := 0; i < 2; i++ {
	//	ori := *str
	//	if i == 0 {
	//		*str = *str + "("
	//	} else {
	//		*str = *str + ")"
	//	}
	//	recurParentesis(step+1, n, res, str, hashMap)
	//	*str = ori
	//}
	*str = *str + "("
	recurParentesis_pattern1(step+1, n, str, hashMap)
	*str = *str + ")"

}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	ori := lists[0]
	for i := 1; i < len(lists); i++ {
		ori = merge2Lists(ori, lists[i])
	}
	return ori
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	res := &ListNode{
		Val:  0,
		Next: nil,
	}
	tem := res
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				tem.Next = &ListNode{
					Val:  l1.Val,
					Next: nil,
				}
				tem = tem.Next
				l1 = l1.Next
			} else {
				tem.Next = &ListNode{
					Val:  l2.Val,
					Next: nil,
				}
				tem = tem.Next
				l2 = l2.Next
			}
		} else if l1 == nil {
			tem.Next = l2
			return res.Next
		} else {
			tem.Next = l1
			return res.Next
		}
	}
	return res.Next
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail := head.Next.Next
	res := head.Next
	res.Next = head
	head.Next = tail
	tem := res.Next
	//cur, next
	cur := res.Next.Next
	var next *ListNode
	if cur != nil {
		next = cur.Next
	}

	for cur != nil && next != nil {
		tail := next.Next
		tem.Next = next
		cur.Next = tail
		next.Next = cur
		tem = cur
		//reset cur
		cur = tail
		if cur != nil {
			next = cur.Next
		} else {
			next = nil
		}
	}
	return res
}

//k<=list length
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	arr := make([]*ListNode, 0)
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		tem := cur
		arr = append(arr, tem)
		cur = cur.Next
	}
	tail := arr[k-1].Next
	arr[0].Next = tail
	for i := 1; i < k; i++ {
		arr[i].Next = arr[i-1]
	}
	tem := tail
	res := reverseKGroup(tem, k)
	arr[0].Next = res
	return arr[k-1]
}

func TestReverse(t *testing.T) {
	a := []int{3, 2, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
	b := []int{3, 2, 5}
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	fmt.Println(b)
}
