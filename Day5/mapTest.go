package main

import (
	"fmt"
)

func main() {
	var map1 map[string]string
	/*
	创建集合
	 */
	 map1 = make(map[string]string)
	map1["France"] = "Paris"
	map1["England"] = "London"
	map1["China"] = "Beijing"
	for country := range map1{
		fmt.Println(country,"\t",map1[country],"\t")
	}
	/*
	删除元素
	 */
	 delete(map1,"France")
	 fmt.Println(map1)

	 var s = make([]int,0)
	 s = []int{1,2,6,7,7,33,5,3,7}
	for value := range s {
		fmt.Println(value)
	}

	//
	//var a = []int{1,2,34,5,6,7,78,8}
	//for value := range a {
	//	fmt.Println(value)
	//}
	//s = append(s,100)
	////copy(s,a)
	//fmt.Println(s)
	//
	//s2 := make([]int,len(s))
	// copy(s2,s)
	//fmt.Println(s2)


}

