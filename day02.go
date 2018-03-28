package main

import "fmt"
var teacher string

const (
	enum1 = iota
	enum2
	enum3
	enum4
	enum5 = "alibaba"
	enum6 = iota
)

func main() {
	var age = 10
	var name = "melon"
	var class = "1802"
	fmt.Println("age:", age, "name:", name, "class:", class)
	school, age := "孔壹学院", 16
	fmt.Println(school, age)
	func1("韩茹")
	age = func2(age,20)
	fmt.Println("new age:",age)
}
func func1(name string) {
	teacher = name;
	fmt.Println("teacher name:", teacher);
}
func func2(age,up int)int  {
	return up+age
}