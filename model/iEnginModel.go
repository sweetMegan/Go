package model

import "fmt"

type IEngine interface {
	Start() //启动
	Speedup() //加速
	Stop()//停止
}
type YAMAHA struct {

}

func (y YAMAHA)Start()  {
	fmt.Print("YAHAMA启动：60")

}
func (y YAMAHA)Speedup(){
	fmt.Println("YAMAHA加速:80")
}
func (y YAMAHA)Stop(){
	fmt.Println("YAMAHA停止：0")
}

type HONDA struct {

}
func (h HONDA)Start(){
	fmt.Println("HONDA启动：40")
}
func (h HONDA)Speedup(){
	fmt.Println("HONDA提速:120")
}
func (h HONDA)Stop()  {
	fmt.Println("HONDA停止:0")

}
