package main

import (
	"fmt"
	"strings"
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

	 map1["Japan"] = "Tokio"
	 fmt.Println(map1)
	 map1["Japan"] = "小日本儿"
	 fmt.Println(map1)

	 map2 := map[int]string{1:"HR",2:"coder",4:"四",5:"五",3:"三"}
	for i := 1; i < len(map2); i++ {
		fmt.Println(map2[i])
	}
	delete(map2,1)
	fmt.Println(map2)


	//s := []int{}
	//s = append(s,2,3,4)
	//s[1] = 2
	s := []string{"helloworld"}
	fmt.Printf("%v",s)
}

