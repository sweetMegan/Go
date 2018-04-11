package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float32
	y float32
	z float32
}
func (p Point)getDistance(target Point){
	distance := math.Sqrt(math.Pow(float64(p.x - target.x),2) + math.Pow(float64(p.y - target.y),2) + math.Pow(float64(p.z - target.z),2))
	fmt.Println(p,"点距离",target,"点:",distance)
	testFunc()
}
func main()  {
	p1 := Point{}
	p1.x = 1
	p1.y = 2
	p1.z = 3
	p2 := Point{x:1,y:1,z:1}
	p1.getDistance(p2)


}
func testFunc(){
	fmt.Println("test")
}