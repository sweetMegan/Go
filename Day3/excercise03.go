package main

import (
	"fmt"
	"math"
)

func main() {
	//给定一个年龄，如果大于22岁，输出可以娶媳妇啦
	age := 23
	func1(age)
	//给定一个数值 输出它的绝对值
	num := -100
	func2(num)
	//给定一个成绩 如果大于等于60，输出及格
	record := 50
	func3(record)
	switch record {
	case 50:
		{
			fmt.Println(50)

		}
		//fallthroug用于穿透swtich
		fallthrough
	case 60:
		{
			fmt.Println(60)
			fmt.Println("----======")
		}
	default:
		fmt.Println("default")

	}
}
func func1(age int) {
	if age > 22 {
		fmt.Println("可以娶媳妇啦")
	} else {
		fmt.Println("还早呢")
	}
}
func func2(num int) {
	fmt.Println("num 绝对值是:", math.Abs(float64(num)))

}
func func3(record int) {
	if record >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

}
