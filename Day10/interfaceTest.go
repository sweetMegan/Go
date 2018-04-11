package main

import (
	"fmt"
	"math"
)
type Shape interface {
	area() float64
	perimeter() float64
}
type Triangle struct {
	a,b,c float64
}

func (tri Triangle)perimeter()float64  {
	return tri.a+tri.b+tri.c
}
func (tri Triangle)area()float64  {
	p := tri.perimeter()/2.0
	return math.Sqrt(p*(p-tri.a) * (p-tri.b) * (p-tri.c))

}
type Rectangle struct {
	a,b float64
}
func (rect Rectangle)perimeter()float64{
	return 2*(rect.a+rect.b)
}
func (rect Rectangle)area()float64{
	return rect.a*rect.b
}
type Circle struct {
	r float64
}
func (c Circle)perimeter()float64{
	return math.Pi*c.r*2
}
func (c Circle)area()float64  {
	return math.Pi*math.Pow(c.r,2)
}

func main()  {
	tri := Triangle{4,3,5}
	fmt.Printf("三角形周长:%.2f\t 面积:%.2f",tri.perimeter(),tri.area())

	rect := Rectangle{2,3}
	fmt.Printf("\n长方形周长:%.2f\t 面积:%.2f\n",rect.perimeter(),rect.area())

	cy := Circle{5}
	fmt.Printf("圆周长:%.2f\t 面积:%.2f\n",cy.perimeter(),cy.area())

	fmt.Printf("三角形面积:%.2f\t 长方形面积:%.2f\t 圆形面积:%.2f\t",testArea(tri),testArea(rect),testArea(cy))
	array := [3]int{1,2,3}
	fmt.Println(array)
}
func testArea(s Shape)float64{
	return s.area()
}