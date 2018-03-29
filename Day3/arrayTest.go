package main

import "fmt"

func main()  {
	NSlog("=======数组初始化======")
	func1()
	func2()
	func3()
	func4()
}
func func4()  {
	var slice = []int{1,2,3,4,5,6,7,8}
	fmt.Println(len(slice),cap(slice))
	slice = append(slice,10)
	fmt.Println(slice,len(slice),cap(slice))
	var slice2 = make([]int,len(slice))
	//拷贝slice到slice2 slice2的长度不能小于slice
	//var slice2 = []int{11}
	copy(slice2,slice)
	fmt.Println(slice2,slice)
}
func func3()  {
	var array = []int{1,2,35,5,2,5,2,56,7}
	//fmt.Println(array)
	NSlogArray(array);
}

func func2()  {
	var array = [5]int{1,2,3,4,6}
	fmt.Println(array)
}
func func1()  {
	var array [5] float32
	for i := 0; i < len(array); i++ {
		array[i] = float32(i)
	}
	fmt.Println(array)
}
func NSlog(str string)  {
	fmt.Println(str)
}
func NSlogArray(array []int ){
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}