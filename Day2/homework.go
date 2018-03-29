package main

import "fmt"

func main() {
	//1、每种基本类型的变量，各声明5个变量，并打印变量的数值，以及类型
	func1();
	//2、声明几个常量
	func2();
	//3、给定两个整数变量，求两个数的和、差、乘积、求商、求余数
	func3();
	//4、给定两个整数变量，比较>，<，>=，<=，==，!=
	func4();
	//5、求res1和res2的值
	func5();
	//6、求a和b的结果
	func6();
	//9、交换2个变量的值
	 func9();
	 //10、定义一个四位数的整数，分贝获取各个位数的值
	 func10();
}
func func10()  {
	fmt.Println("\n华丽分割线=======10=========")
	a := 1976;
	fmt.Printf("千位 :%d 百位:%d 十位:%d 个位:%d",a/1000,(a/100)%10,(a%100)/10,a%10);
}
func swapAB(a int,b int)(int,int)  {
	var temp int = 0;
	temp = a;
	a = b;
	b = temp;
	return a,b;
}
func func9()  {
	fmt.Println("\n华丽分割线=======9=========")

	a := 3;
	b := 4;
	a,b = swapAB(a,b);
	fmt.Println(a,b);
}
func func6()  {
	fmt.Println("\n华丽分割线=======6=========")
	a := 2;
	b := 3;
	a *= b;
	fmt.Println("a = ",a);
	b %= 5;
	fmt.Println("b = ",b);
	a <<= 2;
	fmt.Println("a = ",a);
	b &= a;
	fmt.Println("b = ",b);
}
func func5(){
	fmt.Println("\n华丽分割线=======5==========")
	a := 4;
	b := 3;
	res1 := a<b && b/2 == 1 && a%3 != 0;
	res2 := (a+b)*3 < a<<2 || (a-b) >0;
	fmt.Println(res1,res2);
}
func func4(){
	var a int = 3;
	var b float32 = 3.14;
	fmt.Printf("a>b = %t a<b = %t a>=b %t a<=b %t a!=b %t",a>int(b),float32(a)<b,float32(a)>=b,float32(a)<=b,float32(a)!=b);
}
func func1()  {
	//整型
	//var i int;
	//var i2 int8; byte
	//var i3 int16;rune
	//var i4 int32;
	//var i5 int64;
	//浮点型
	//var f float32;
	//var f2 float64;
	//字符串
	//var str string;
	//布尔型
	//var flag bool;
}
func func2(){
	//const name  = "name";
	//const age int = 32;
	//const (
	//	class = "1802"
	//	school = "孔壹学院"
	//	studentId = iota
	//	studentId2 = iota
	//)
	//fmt.Println(name,age,class,school,studentId,studentId2);

}
func func3(){
	a := 3;
	b := 7;
	sum := a+b;
	sub := a-b;
	pro := a*b;
	div := a/b;
	mod := a%b;
	fmt.Println("和",sum,"差",sub,"积",pro,"商",div,"取余",mod);
	fmt.Println("^",4^12);
}

