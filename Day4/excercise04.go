package main

import "fmt"

func main() {
	//1、打印100以内所有偶数
	//func1()
	//2、打印1-100以内，能被3整除，不能被5整除的数，每行打印5个
	//func2()
	//3、求1-100内，奇数的和
	func3()
	//4、打印2-100内的素数
	func4()
	//5、百元白鸡 一百元买100只鸡，公鸡5元/只，母鸡3元/只 小鸡1元/3只
	func5()
}
func func5() {
	//5g+3m+x/3 = 100 元
	//g+m+x = 100只
	//money := 100
	total := 100
	totalG := 0
	moneyG := 0
	totalM := 0
	moneyM := 0
	totalX := 0
	moneyX := 0
	for totalG = 1; totalG <= 20; totalG++ {
		moneyG = 5 * totalG
		for totalM = 1; totalM <= 33; totalM++ {
			moneyM = 3 * totalM
			for totalX = 1; totalX <= total-totalG-totalM; totalX++ {
				if totalX%3 == 0 {
					moneyX = totalX / 3
					if moneyX+moneyM+moneyG == total && totalX+totalM+totalG == total {
						fmt.Printf("\n嗯，知道了!公鸡：%d，母鸡：%d,小鸡仔:%d", totalG, totalM, totalX)
						//return
					}
				}
			}
		}
	}
	fmt.Println("\nTMD 算不出来")
}
func func4() {
	for i := 2; i <= 100; i++ {
		count := 0
		for j := 1; j <= i/2; j++ {
			//fmt.Println(j,"%",i,j%i)
			if i%j == 0 {
				count++
			}
		}
		if count == 1 {
			fmt.Print(i, "\t")
		}
	}
}
func func3() {
	sum := 0
	for i := 0; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println("sum = ", sum)

}
func func2() {
	count := 0
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Print(i, "\t")
			if (4-count)%5 == 0 {
				fmt.Println()
			}
			count++
		}

	}
}
func func1() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Print(i, "\t")
		}
		if (9-i)%10 == 0 {
			fmt.Println()
		}
	}
}
