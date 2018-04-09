package main
import "fmt"
type Employee struct {
	id string
	name string
	age int
	salary float32
}
type User struct {
	username string
	password string
}
type Wheel struct {
	radius_Max float32 //外半径
	radius_Min float32 //内半径
}
type Car struct {
	brand string
	wheels []Wheel
}
type Car2 struct {
	Car
	color string
	speed float32
}
func (c Car2)move(){
	fmt.Println(c.brand,"启动")
}
func (c Car2)stop(){
	fmt.Println(c.brand,"停车")
}
type Bike struct {
	Car2
	power string
}
type Racing struct {
	Car2
	capability float32
}
type Person09 struct {
	name string
	age int
}
type Student struct {
	Person09
	score float32 //成绩
}
type Worker struct {
	Person09
	salary float32 //工资
}
func (stu Student)printInfo(){
	fmt.Printf("学生%s，年龄：%d,成绩:%.1f\n",stu.name,stu.age,stu.score)
}
func (w Worker)printInfo(){
	fmt.Printf("工人%s,年龄:%d,工资:%.2f\n",w.name,w.age,w.salary)
}


func main()  {
	var e1 Employee
	e1.name = "程序员"
	e1.id = "00001"
	e1.age = 20
	e1.salary = 9000.0
	fmt.Println(e1)

	user1 := User{"河流","password"}
	fmt.Println(user1)


	wheels := []Wheel{}
	for i := 0;i<4;i++{
		wheel := Wheel{45,35}
		wheels = append(wheels,wheel)
	}
	car := Car{"东方红",wheels}
	fmt.Println(car)

	bike := Bike{Car2{Car{"洋车",[]Wheel{Wheel{45,35},Wheel{45,35}}},"黑色",45.0},"电动"}
	bike.move()
	bike.stop()

	racing := Racing{Car2{Car{"跑狼",[]Wheel{Wheel{45,35},Wheel{45,35}}},"黑色",45.0},2.5}
	racing.move()
	racing.stop()

	stu := Student{Person09{"小明",18},100}
	stu.printInfo()
	w := Worker{}
	w.name = "奥特曼"
	w.age = 40
	w.salary = 5000.99
	w.printInfo()
}
