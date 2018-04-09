package main

import (
	"fmt"
	"math"
)

func main() {
	//pro := getPro(5)
	//fmt.Println(pro)
	//a, b := 1, 2
	//fmt.Println(a, ",", b, "中最大的是：", getMax(a, b))
	//var r float64 = 2
	//area, _ := getArea(r)
	//fmt.Println("半径为r的圆，面积是:", area)
	fmt.Println(fibonacci(6))
}
func fibonacci(n int)int{
	if n == 2 || n == 1{
		return 1
	}else {
		return fibonacci(n-1)+fibonacci(n-2)
	}

}
func getPro(n int) int {
	if n == 1 {
		return 1
	}else {
		return getPro(n-1)*n
	}

	//pro := 1
	//for i := 1; i <= n; i++ {
	//	pro *= i
	//}
	//fmt.Println(n, "的阶乘是:", pro)
	//return pro
}
func getMax(a, b int) (int) {
	if a > b {
		return a
	} else {
		return b
	}
}
func getArea(r float64) (area float64, peri float64) {
	fmt.Println(math.Pow(r, 2.0))
	area = math.Pi * math.Pow(r, 2.0)
	peri = math.Pi * r * 2
	return
}
//func getArea2(r float64) (float64, float64) {
//	area := math.Pi * math.Pow(r, 2.0)
//	return area
//}
