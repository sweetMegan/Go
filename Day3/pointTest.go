package main

import "fmt"

func main() {
	fmt.Println("=======指针=======")
	var a = 10
	fmt.Println("a的地址:", &a)

	var ip *int
	ip = &a
	fmt.Println(ip, *ip)

	func2()
}
func func2() {
	fmt.Println("指针数组")
	//定义数组时，数组长度要么不写，要写就是常量
	a := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] = %d \n", i, a[i])
		if i == 2 {
			break
		}
	}
	//数组指针长度 可以是变量
	var ptr [len(a)]*int
	for i := 0; i < len(a); i++ {
		ptr[i] = &a[i]
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] = %d \n", i, *ptr[i])
	}
}
