package main

import (
	"fmt"
	"reflect"
)

func main() {
	//1、使用if语句完成：给定数字：如果为1就输出星期一，如果为2，就输出星期二，以此类推，一直到7，输出星期日。如果是其他数字，就输出"错误信息"
	//2、switch完成 1
	func1(5)
	//3、使用if语句完成：给定月份，输出该月属于哪个季节。3，4，5春季 6，7，8夏季 9,10,11秋季 12，1，2冬季
	//4、switch完成
	func2(6)
	//6、作业10，使用switch完成
	//7、模拟登录，键盘上输入用户名和密码，如果用户名是admin密码是123，或者用户名是zhangsan密码是zhangsan123，表示可以登录。否则打印登录失败
	//func3()
	//8、使用if语句完成：给定年龄，如果小于18岁，输出青少年，如果大于等于18并且小于30岁，输出青年，否则输出中老年。
	func4(100)
	//9、求给定的年份，是否是闰年
	func5(2000)
	//13、用if语句完成：从键盘输入考试分数（整数），如果成绩大于90，输出学霸实在是牛，如果大于等于80，输出学神要加油，如果大于等于70，
	// 输出学民好害羞，如果大于等于60输出学弱打酱油，如果小于60，输出学渣泪在流
	func6()
	//14、模拟计算器
	////func7()
	////16、给定三个整数，从小到大输出
	//func8()
	////19、打印'A'到'Z'的字母
	//func9()
}
func func9()  {
	//fmt.Printf("%v %v",'A','Z')
	for i := 0;i < 26 ;i++  {
		fmt.Printf("%q ",i+65)
		if  (9-i)%10 == 0{
			fmt.Println()
		}
	}
}
func func8()  {
	array := []int{1,3,5,4,9,6,2,10}
	fmt.Println(array)
	for i := 0; i < len(array); i++ {
		for j := i+1; j < len(array); j++ {
			if array[i] < array[j] {
				//fmt.Println("交换",array[i],array[j])
				temp := array[i]
				array[i] = array[j]
				array[j] = temp
			}
			fmt.Println("排序后:",array)

		}

	}

}
func func7()  {
	fmt.Println("输入第一个整数")
	a := 0
	fmt.Scanln(&a)
	fmt.Println("输入第二个整数")
	b := 0
	fmt.Scanln(&b)
	fmt.Println("输入运算符")
	method := ""
	fmt.Scanln(&method)
	result := 0
	switch method {
	case "+":
		{
			result = a + b
		}
	case "-":
		{
			result = a - b
		}
	case "*":
		{
			result = a*b
		}
	case "/":
		{
			result = a/b
		}
	default:
		{
			result = 0
			fmt.Println("这个我算不了...")
			return
		}

	}
	fmt.Println(a,method,b,"=",result)

}
func func6()  {
	fmt.Println("输入分数:")
	score := 18
	fmt.Scan(&score)
	fmt.Printf("%v:%T \n",reflect.TypeOf(score).Kind().String(),reflect.TypeOf(score).Kind().String())
	//判断输入的值的类型 ？？？？？
	if  reflect.TypeOf(score).Kind().String() == "int"{
		//fmt.Println("整型")
		if score > 90 {
			fmt.Println("学霸实在是牛")
		}else if score >= 80 {
			fmt.Println("学神要加油")
		}else if score >= 70{
			fmt.Println("学民好害羞")
		}else if score >= 60 {
			fmt.Println("输出学弱打酱油")
		}else {
			fmt.Println("学渣泪在流")
		}
	}else {
		fmt.Println("请输入一个整数")
	}

}
func func5(year int)  {
	if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
		fmt.Println("闰年")
	}else {
		fmt.Println("平年")
	}
}
func func4(age int){
	if  age < 18 && age >= 0{
		fmt.Println("青少年")
	}else if age >= 18 && age < 30 {
		fmt.Println("青年")
	}else if age < 200 && age >= 30 {
		fmt.Println("中老年")
	}else{
		fmt.Println("别闹")
	}
}
func func3()  {
	for i := 0; i < 3; i++ {
		fmt.Println("请输入用户名")
		name := ""
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		password := ""
		fmt.Scanln(&password)
		if (name == "admin" && password == "123")||(name == "zhangsan" && password == "zhangsan123") {
			fmt.Println("登录成功")
			break
		}else {
			chance := 2-i
			if chance>0 {
				fmt.Println("登录失败,还有",chance,"次机会")
			}else {
				fmt.Println("重新申请一个吧")
			}
		}
	}


}
func func2(month int) {
	switch month {
	case 3, 4, 5:
		{
			fmt.Println("春季")
		}
	case 6, 7, 8:
		{
			fmt.Println("夏季")
		}
	case 9, 10, 11:
		{
			fmt.Println("秋季")
		}
	case 12, 1, 2:
		fmt.Println("冬季")
	default:
		{
			fmt.Println("没有这个月份")
		}
	}
}
func func1(day int) {
	//方法一
	//var array = [7]string{"星期一","星期二","星期三","星期四","星期五","星期六","星期日"}
	//for i := 1; ; {
	//	fmt.Scanln(&i)
	//	if i>0 && i<=7 {
	//		fmt.Println(array[i-1])
	//	}else {
	//		fmt.Println("错误信息：输入值不合法，请输入1到7的整数")
	//		break
	//	}
	//
	//}

	//方法二
	//if day == 1 {
	//	fmt.Println("星期一")
	//}else if day == 2 {
	//	fmt.Println("星期二")
	//}else if day == 3 {
	//	fmt.Println("星期三")
	//}else if day == 4 {
	//	fmt.Println("星期四")
	//}else if day == 5 {
	//	fmt.Println("星期五")
	//}else if day == 6 {
	//	fmt.Println("星期六")
	//}else if day == 7 {
	//	fmt.Println("星期日")
	//}else {
	//	fmt.Println("错误信息：输入值不合法，请输入1到7的整数")
	//}

	//方法三
	switch day {
	case 1:
		{
			fmt.Println("星期一")
		}
	case 2:
		{
			fmt.Println("星期二")
		}
	case 3:
		{
			fmt.Println("星期三")
		}
	case 4:
		{
			fmt.Println("星期四")
		}
	case 5:
		{
			fmt.Println("星期五")
		}
	case 6:
		{
			fmt.Println("星期六")
		}
	case 7:
		{
			fmt.Println("星期日")
		}
	default:
		{
			fmt.Println("错误信息：输入值不合法，请输入1到7的整数")
		}

	}

}
