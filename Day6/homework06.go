package main

import (
	"fmt"
	"strings"
)

func main() {
	//1、	str := "hello hello hello world" 统计字符串的长度，统计"llo"出现的次数 统计"wor"出现的位置
	func1()
	//2、统计该字符串中每个字母出现的次数"ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"
	func2()
	//3、将字符串"aa:zhangsan@163.com!bb:lisi@sina.com!cc:wangwu@126.com"，按照指定的key-value规则存入map中
	func3()
}
func func3() {
	fmt.Println("华丽的分割线=================3===============")
	str := "aa:zhangsan@163.com!bb:lisi@sina.com!cc:wangwu@126.com"
	source := strings.Split(str, "!")
	fmt.Println(source)
	result := make(map[string]string)
	for i := range source {
		temp := strings.Split(source[i], ":")
		if len(temp) == 2 {
			result[temp[0]] = temp[1]
		}
	}
	fmt.Println(result)
	for key, value := range result {
		fmt.Println(key, ":", value)
	}
}
func func1() {
	fmt.Println("华丽的分割线=================1===============")
	str := "hello hello hello world"
	//统计字符串的长度
	length := len(str)
	fmt.Println("str的长度是:", length)
	//统计"llo"出现的次数
	count := strings.Count(str, "llo")
	fmt.Println("llo出现:", count, "次")
	//统计"wor"出现的位置
	location := strings.Index(str, "wor")
	fmt.Println("wor在：", location)
}
func func2() {
	fmt.Println("华丽的分割线=================2===============")
	str := "ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"
	for len(str) != 0 {
		targetC := str[0:1]
		count := strings.Count(str, targetC)
		str = strings.Replace(str, targetC, "", -1)
		fmt.Println(targetC, "出现:", count, "次")
	}
}
