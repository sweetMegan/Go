package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}
func main() {
	a := 10
	var p *int
	p = &a
	fmt.Printf("%p,%v,%p,%p\n",p,p,&p,a)
	fmt.Println(*p,"a = ",a)

	s := []int{1,2,3,4}
	var ps *[]int
	ps = &s
	fmt.Printf("s地址:%p  s=%v s指向的地址:%p",&s,s,s)
	fmt.Printf("\n ps地址：%p ps的值：%v ps指向的地址:%p ps指向地址的值:%v",&ps,ps,ps,*ps)
	s = append(s,5)
	fmt.Println("\n========")
	fmt.Printf("s地址:%p  s=%v s指向的地址:%p",&s,s,s)
	fmt.Printf("\n ps地址：%p ps的值：%v ps指向的地址:%p ps指向地址的值:%v\n",&ps,ps,ps,*ps)


	array := [4]int{6,7,8,9}
	var pa *[4]int
	pa = &array
	fmt.Printf("array的地址:%p array = %v array:%p",&array,array,array)
	fmt.Println()
	fmt.Printf("pa的值：%v",pa)
}
func swapAB2(a,b int){
	var temp int
	temp = a
	a = b
	b = temp
	fmt.Println(a,b)
}
func swapAB(a,b *int){
	temp := *a
	*a = *b
	*b = temp
	//fmt.Println("========")
	//fmt.Println(a,b)
	//var temp *int
	//fmt.Println("1、",temp,a,b)
	//temp = a
	//fmt.Println("2、",temp,a,b)
	//a = b
	//fmt.Println("3、",temp,a,b)
	//b = temp
	//fmt.Println("4、",temp,a,b)
	//fmt.Println("+++++++++")


}
func func08_02(a,b int,f func(int ,int )int)int{
	return f(a,b)
}
func func08_03()(func()(int)){
	x := 0
	f := func()int{
		x++
		return x
	}
	return f
}
func func08_04()(func()(int)){
	x := 0
	f := func()int{
		x++
		return x
	}
	return f
}
func func08_01(){
	fmt.Println(func08_03)

	f := func08_03()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	fmt.Println(f)


}
//func main() {
//	fmt.Printf("%T",func08_1)
//}
//
//func func08_1(){
//
//}
