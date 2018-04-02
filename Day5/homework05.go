package main

import (
	"fmt"
	"math/rand"
	"time"

)

func main(){
	//1、某个公司采用电话传递数据，数据时四位的整数，在传递过程中是加密的，加密规则如下：每位数字都加上5，然后用和，除以10的余数代替该数字，再将第一位和第四位交换，第二位和第三位交换
	func1()
	//2、给定一个整型数组，长度为10，数字取自随机数
	var array = [10]int{}
	array = func2()
	//3、排序2
	func3(array)
	//4、创建一个切片，并向切片中存储10个数字，打印输出
	fmt.Println("华丽分割线==============4===============")
	s := make([]int,10,10)
	s = array[:]
	fmt.Println(s)
	s1 := array[:]
	fmt.Println(s1)
	//5、有n个人围成一个圈，顺序排号。从第一个人开始报数（1到3报数），凡报道3的人退出圈子，为最后留下的1个人是原来的第几号。
	func5()

}
func func5()  {
	fmt.Println("华丽分割线==============5===============")

	array := [5]int{}
	for i := range array{
		array[i] = i+1
	}
	fmt.Println(array)
	temp := 1 //需要报的数 1,2,3
	count := 0 //已出列人数 count = len(array)-1 时，所有人员已出列


	for i := 0; count != len(array) - 1 ;i++  {
		//如果 i == len(array) 开始下一轮报数
		if i == len(array) {
			i = 0
		}
		//如果 array[i] == 0 表示此人已出列，下一个人报数
		if array[i] == 0 {
			continue
		}else {
			//未出列人报数
			if temp%3 == 0 {
				temp = 1
				array[i] = 0
				fmt.Println("===",array)
				count++
			}else {
				temp++
			}

		}

	}
	fmt.Println("华丽分割线==============5===============")

}
func func3(array [10]int){
	fmt.Println("华丽分割线==============3===============")

	exchanged := true
	for i := range array{
		if !exchanged {
			break
		}
		exchanged = false
		for j := 0; j < len(array)-i-1;j++  {
			if array[j] > array[j+1] {
				array[j],array[j+1] = array[j+1],array[j]
				exchanged = true
			}
		}
	}
	fmt.Println(array)
}
func func2()([10]int){
	fmt.Println("华丽分割线==============2===============")
	array := [10]int{}
	for i := range array{
		//1、创建随机数
		rand.Seed(time.Now().UnixNano())
		value := rand.Intn(10)
		array[i] = value
	}
	fmt.Println("array : ",array)
	return  array
}
func func1()  {
	message := 2384
	//1、每位数字加5
	thousand,hundred,ten,number := getNumbers(message)
	result1 := thousand+5+hundred+5+ten+5+number+5
	//2、用和除以10的余数代替该数字
	result2 := result1%10
	result3 := [4]int{}
	result3[0],result3[1],result3[2],result3[3] = getNumbers(result2)
	fmt.Println("result3:",result3)
	//3、再将第一位和第四位交换，第二位和第三位交换
	//fmt.Println(result2,"===",result1)
	result3[0],result3[1],result3[2],result3[3] = result3[3],result3[2],result3[1],result3[0]
	fmt.Println("加密后消息:",result3)
}

func getNumbers(message int)(int,int,int,int)  {
	thousand := message/1000 //千位
	hundred := message/100%10 //百位
	ten := message%100/10 //十位
	number := message%200%10 //个位
	return thousand,hundred,ten,number
}
