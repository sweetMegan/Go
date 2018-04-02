package main

import (
	"fmt"
	"time"
)

func main() {
	//arrayF := [...]float32{1.0,2.0}
	//fmt.Println(arrayF)
	//arrayF2 := []float32{3.0,4.4}
	//fmt.Println(arrayF2)
	//fmt.Printf("%T\n",arrayF2)
	////copy(arrayF2,arrayF)
	//var s = make([]float32,2)
	//copy(s,arrayF2)
	//fmt.Println(s)
	//arrayF2[0] = 1.5
	//fmt.Println("s= =",s)
	//fmt.Println("arrF2",arrayF2)
	//
	//var arrStr  = [...]string{"aliliba","en","o"}
	//for i := 0; i < len(arrStr); i++ {
	//	fmt.Print(arrStr[i],"\t")
	//}
	//
	//var arr1  = [10]int{5,4,3,7,10,2,9,8,4,10}
	//sum := 0
	//for _,value := range arr1	{
	//	sum += value;
	//}
	//fmt.Println(" \n sum =  ",sum)


	//var s = make([]int,5)
	//for i := range s{
	//	s[i] = i
	//	fmt.Print(s[i],"\t")
	//}
	//fmt.Println()
	//var s2 = make([]int,5)
	//for i := range s{
	//	s2[i] = i+10
	//	fmt.Print(s2[i],"\t")
	//}
	//fmt.Println()
	////copy(s,s2)
	//s = s2
	//fmt.Println(s)
	//s2[0] = 100
	//fmt.Println(s)

	//var array = [...]int{1,2,3,4,6,7,78,3}
	//for i := 0; i < len(array); i++ {
	//	for j := 0; j < len(array)-i-1 ;j++  {
	//		if array[j] > array[j+1] {
	//			array[j],array[j+1] = array[j+1],array[j]
	//		}
	//	}
	//}
	//fmt.Println(array)

	//array := [3][3]int{
	//	{1,2,3},
	//	{4,5,6},
	//	{7,8,9}}
	//for _,value := range array{
	//	fmt.Println(value[2])
	//}

	var s = make([]int,8,8)
	fmt.Println(s)
	s2 := make([]int,8,8)
	fmt.Println(s2)
	array := [...]int{1,2,3,4,5,6,7,8,9}
	//s2 := make([]int,9,9)
	s3 := array[2:9]
	fmt.Println(s3)
	s4 := array[:]
	fmt.Println(s4)
	s5 := array[0:9]
	fmt.Println(s5)

	year := time.Now().Year()
	fmt.Println(year)
}
