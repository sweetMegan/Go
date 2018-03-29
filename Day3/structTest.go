package main

import "fmt"
type Book struct {
	title string
	author string
	price float32

}
type Library struct {
	name string
	manager string
	book Book
}
func main()  {
	fmt.Println("\n 结构体")
	var book1 Book
	book1.title = "西游记"
	book1.author = "吴承恩"
	book1.price = 100.8;
	fmt.Println(book1)
	var library1 Library
	library1.name = "国家图书馆"
	library1.manager = "管理员"
	library1.book = book1
	fmt.Println(library1)
	scanfTest();

}
func scanfTest()  {
	fmt.Println("输入用户名，密码")
	name := "name"
	password := "password"
	fmt.Scanln(&name,&password)
	fmt.Println("用户名:",name,"密码:",password)
	//a := "a"
	////b := "b"
	//c := 1
	//fmt.Scanf("%s ,%d",&a,&c)
	//fmt.Println(a,c);
}