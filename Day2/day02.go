package main

import (
	"fmt"
	"strconv"
)
var teacher string

const (
	enum1 = iota
	enum2
	enum3
	enum4
	enum5 = "alibaba"
	enum6 = iota
)
const cA  = 5;
func main() {
	var age = 10
	var name = "melon"
	var class = "1802"
	fmt.Println("age:", age, "name:", name, "class:", class)
	//：=声明变量 同时声明多个的话，其中至少要有一个未声明过的变量
	school, age := "孔壹学院", 16
	fmt.Println(school, age)
	func1("韩茹")
	age = func2(age,20)
	fmt.Println("new age:",age)
	//空白舍弃标识符 用于舍弃数值
	//a,b := 10,20;
	_,b := 10,20;
	fmt.Println(b);
	func3(1,2)
	fmt.Println(func3(1,2))
	//数组
	var array[5] string
	fmt.Println("array长度:",len(array))
	for i := 0; i<len(array);i++	  {
		array[i] = strconv.Itoa(i);
	}
	fmt.Println("array:",array);
	//len()
	//
	//const cAGE = 1+cA;
	//const cAGE2 = lenTest();
	//fmt.Println(cAGE,cAGE2);
}
func lenTest() int  {
	return 1+cA
}
func func1(name string) {
	teacher = name;
	fmt.Println("teacher name:", teacher);
}
func func2(age,up int)int  {
	return up+age
}
func func3(a,b int)(int,int)  {
	return 11,22;
}
