package main

import (
	"fmt"
	"encoding/xml"
)
type Person struct {
	name string
	age int
	sex string
	address Address
	books []Book
}
type Book struct {
	name string
	price float32
}

type Address struct {
	city,province string
}
type Animal struct {
	name string
	age int
}
type Cat struct {
	Animal
	class string
}
type Dog struct {
	Animal
	color string
}
func main(){
	books := make([]Book,0)
	book1 := Book{"奥特曼",10.0}
	books = append(books,book1)
	book2 := Book{}
	book2.name = "黑猫警长"
	book2.price = 12.5
	books = append(books,book2)

	p1 := Person{"zhq" ,1,"male",Address{"武汉","湖北"},books}
	fmt.Println("姓名:",p1.name,"年龄:",p1.age,"性别:",p1.sex,"地址:",p1.address,"书籍1:",p1.books[0],"书籍2:",p1.books[1].name)

	cat := Cat{Animal{"kitty",2},"波斯猫"}
	fmt.Println(cat)
	dog := Dog{Animal{"旺财",1},"黄色"}
	fmt.Println(dog)
	dog.name = "八公"
	fmt.Println(dog)

	//struct {
	//	name string
	//	age int
	//}{"zhq",20}

}
