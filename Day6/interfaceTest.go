package main

import "fmt"

type Phone interface {
	call()
	reject()
}
type NokiaPhone struct {
}
type Iphone struct {
}

func (nokia NokiaPhone) call() {
	fmt.Println("I am Nokia,I can call you ")
}
func (nokia NokiaPhone) reject()  {
	fmt.Println("I am Nokia,I am busy")
}
func (iphone Iphone) call() {
	fmt.Println("I am Iphone ,I can call you")
}
func (iphone Iphone) reject()  {
	fmt.Println("I am Iphone,I am busy too")
}
func func1(i int) (int) {
	return i
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone.reject()

	phone = new(Iphone)
	phone.call()
	phone.reject()

}
//func call() {
//
//}
//func reject()  {
//
//}