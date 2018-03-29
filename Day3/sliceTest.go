package main

import "fmt"

func main()  {
	fmt.Println("========初始化切片=======")
	func1()
	func2()
	func3()
}
func func1()  {
	var slice1  = make([]int,3)
	slice1 = []int{1,3,4}
	fmt.Println(slice1)
}
func func2(){
	slice := []int{3,6,8,5,7,3,7,27}
	fmt.Println(slice)
}
func func3()  {
	array :=[]int{1,3,6,5,4,5,6,7,3,10,33,22,5,5}
	/**
	*[startIndex:endIndex]
	*array中从下标startIndex到下标endIndex-1个元素
	*/
	//slice := array[:]
	slice := array[0:4]
	//slice := array[:]
	fmt.Println(slice)
	fmt.Println("slice长度:",len(slice),"slice容量:",cap(slice))
}