package wenhaialgo

import (
	"fmt"
	"math"
	"sort"
)

func FindMedianBySort(array []int) (median float64){
	if len(array)<1{
		return -1
	}
	sort.Ints(array)
	l := len(array)
	if l%2==0{
		return 	float64(array[(l+1)/2-1] + array[(l+1)/2])/2
	}else {
		return  float64(array[(l+1)/2-1])
	}
}

func FindMedianByHeap(array []int) (median float64){
	if len(array) ==0{
		return math.MinInt64
	}
	if len(array) == 1{
		return float64(array[0])
	}
	if len(array) == 2{
		return float64((array[0]+array[1]) /2)
	}
	myheap := NewHeap()
	length := len(array)
	loop := length/2 + 1
	myheap.Elements = array[0:loop]
	myheap.MinHeapBuild()
	for i := loop; i < len(array); i++{
		if array[i] > myheap.Elements[0]{
			myheap.Elements[0] = array[i]
			myheap.MinHeapAdjustRoot()
		}
	}
	if length%2==0{
		second := myheap.Elements[1]
		if  second > myheap.Elements[2]{
			second = myheap.Elements[2]
		}
		return 	float64(myheap.Elements[0] + second)/2
	}else {
		return  float64(myheap.Elements[0])
	}
}


func FindMedianByForLoop(array []int) (median float64){
	if len(array)<1{
		return -1
	}
	i := (len(array)+1) /2 + 1
	//step 1: get minimum and second minimum in first half
	newarr := array[0:i]
	max, second, indexmax, indexsec := FindTopTwoMaxNumber(newarr)
	fmt.Println(max, second, indexmax, indexsec)
	//step 2: get the smaller half elements in array
	for k := i; k<len(array);k++{
		if array[k] >= max{
			continue
		}else if array[k]>second{
			max = array[k]
			newarr[indexmax] = max
		}else {
			max = second
			newarr[indexmax] = max
			newarr[indexsec] = array[k]
			max, second, indexmax, indexsec  = FindTopTwoMaxNumber(newarr)
		}
	}
	//get the median
	fmt.Println("len(array)%2 = ",len(array)%2)
	fmt.Println("minimum, second ",max, second)
	if len(array)%2==0{
		median= float64(max + second)/2
	}else {
		median = float64(max)
	}
	return
}