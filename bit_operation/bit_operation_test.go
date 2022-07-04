package bit_operation

import (
	"fmt"
	"testing"
)

//计算一个32位整数换算成二进制后1的个数，包括负数
/*
负数的二进制换算
1。 原来的数：001
2。 反码：1......110
3。 补码：反码+1： 1......110
4。 负数等于补码
*/

func TestGetNumberOf1(t *testing.T) {
	n := int32(-4)
	fmt.Println(getNumberOf1(n))
}

//n &= (n – 1)总能清除最右边的一个1
//因为减1后，最后一个1的位置肯定变为0，所以用且总是去掉最后一个1
func getNumberOf1(n int32) int {
	count := 0
	for n != 0 {
		count++
		n = n & (n - 1)
	}
	return count
}

func TestDivideUsingBit(t *testing.T) {
	a := -2
	b := -1
	fmt.Println(divideUsingBit(a, b))
}

func TestPlusUsingBit(t *testing.T) {
	a := 2
	b := -3
	fmt.Println(plusUsingBit(a, b))
	fmt.Println(substractUsingBit(a, b))
	fmt.Println(multipleUsingBit(a, b))
}

//用位运算实现整数除法
//时间复杂度log2（dividend-divisor）
//this is only for positive number
func divideUsingBit(dividend, divisor int) int {
	if dividend == divisor {
		return 1
	}
	if dividend < divisor {
		return 0
	}
	if divisor == 0 {
		return 0
	}
	res := 1
	tem_divisor := divisor
	//fmt.Println(tem_dividend, tem_divisor)
	for tem_divisor < dividend {
		tem_divisor = tem_divisor << 1
		res = res << 1
	}
	res = res + divideUsingBit(dividend-tem_divisor, divisor)
	return res
}

//refer: https://rgb-24bit.github.io/blog/2019/bitop.html
func plusUsingBit(a, b int) int {
	sum := a
	carry := b
	for carry != 0 {
		tem := sum
		sum = tem ^ carry //a^b
		carry = (tem & carry) << 1
	}
	return sum
}

func substractUsingBit(a, b int) int {
	sum := a
	carry := -b
	for carry != 0 {
		tem := sum
		sum = tem ^ carry //a^b
		carry = (tem & carry) << 1
	}
	return sum
}

func multipleUsingBit(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	res := a
	if b > 0 {
		b--
		for b > 0 {
			res += a
			b--
		}
	} else {
		tem := -b
		tem--
		for tem > 0 {
			res += a
			tem--
		}
		res = -res
	}
	return res
}
