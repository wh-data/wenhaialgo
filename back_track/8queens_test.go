package back_track

import (
	"fmt"
	"testing"
)

func TestQueensSolutions(t *testing.T) {
	var (
		isPrintIncomplete = false
	)
	QueensSolutions(4, isPrintIncomplete)
}

func TestQueenSecondPractice(t *testing.T) {
	n := 8
	rawTotal := QueenSecondPractice(n)
	fmt.Println(rawTotal)
	//unique ways
	fmt.Println(rawTotal / factorial(n))
}

func TestFactorial(t *testing.T) {
	n := 8
	fmt.Println(factorial(n))
	fmt.Println(factorialPrint(n))
}

func factorial(n int) int {
	if n <= 1 {
		fmt.Print(n, " ")
		return 1
	}
	fmt.Print(n, " ")
	return n * factorial(n-1)
}

func factorialPrint(n int) int {
	if n <= 1 {
		fmt.Print(n, " ")
		return 1
	}
	x := n * factorialPrint(n-1)
	fmt.Print(n, " ")
	return x
}

func TestXpowerFunction(t *testing.T) {
	fmt.Println(powerFunction(64, 8))
}

func powerFunction(base, power int) int {
	if power <= 1 {
		return base
	}
	return base * powerFunction(base, power-1)
}
