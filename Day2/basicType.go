package main

import (
	"fmt"
	"math"
	"strconv"
)

func main()  {
	fmt.Println(12)
	//int
	//var i int = 8
	//var i2 byte = 32
	//var i3 rune  = 64
	//fmt.Printf("i = %d 类型是%T \n",i,i)
	//fmt.Printf("i2 = %d 类型是%T \n",i2,i2)
	//fmt.Printf("i3 = %d 类型是%T \n",i3,i3)
	//var c = 'a'
	//fmt.Printf("c = %d c = %v c = %q 类型是%T \n",c,c,c,c)
	//var s string = "melon"
	//fmt.Printf("s = %s 类型是%T \n",s,s)

	//a :=100
	//b :=3.14
	//c := "hello"
	//d := `王二狗住在隔壁`
	//e := true
	//f := 'A'
	//fmt.Printf("a = %d a = %q 类型是%T \n",a,a,a)
	//fmt.Printf("b = %f  类型是%T \n",b,b)
	//fmt.Printf("c = %s  类型是%T \n",c,c)
	//fmt.Printf("d = %s  类型是%T \n",d,d)
	//fmt.Printf("e = %t  类型是%T \n",e,e)
	//fmt.Printf("f = %q  类型是%T \n",f,f)
	//
	f := 3.1405926
	fmt.Printf("%f,%.2f \n",f,f)
	fmt.Printf("%f \n",math.Ceil(f))
	fmt.Printf("%f \n",math.Floor(f))
	//保留两位四舍五入
	s :=  strconv.FormatFloat(f,'f',2,32)
	s2 :=  strconv.FormatFloat(f,'b',2,32)
	s3 :=  strconv.FormatFloat(f,'E',2,32)
	s4 :=  strconv.FormatFloat(f,'g',2,32)
	s5 :=  strconv.FormatFloat(f,'G',2,32)

	fmt.Printf("s = %s \n",s);
	fmt.Printf("s2 = %s \n",s2);
	fmt.Printf("s3 = %s \n",s3);
	fmt.Printf("s4 = %s \n",s4);
	fmt.Printf("s5 = %s \n",s5);

	//保留两位不四舍五入
	f2 := math.Floor(f*100)/100
	fmt.Printf("%.2f",f2)
	var b = 4
	 b %= 3
	 b = b%3
	 
}
