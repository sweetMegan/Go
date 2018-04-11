package main

import (
	"Go/model"
	"fmt"
)

type Project struct {
	name string
	model.DataBase
}

func (p Project)testDataBase(d model.DataBase)  {
	d.Insert(p.name)
	d.Update(p.name)
	d.Query(p.name)
	d.Delete(p.name)
}
func main()  {
	//3、定义一个IEngine接口，3个方法：start(),speedup(),stop()，表示启动，加速，停止 定义两个结构体实现该接口：YAMAHA和HONDA
	//实现方式分别是:
	//YAMAHA :
	//启动：60，加速80，停止0
	//HONDA
	//启动：40，加速120，停止0
	//定义一个Car结构体，将IEngine作为字段，再提供一个car的方法，testIEngine()，用于测试改小汽车的发动机。也就是在testIEngine（）中调用start()，speedup(),stop()方法
	homewordFunc_03()
	//4、定义一个DataBase接口，4个方法：insert(),update(),query(),delete()
	//定义两个结构体实现该接口:Mysql和Oracle
	//定义一个Project结构体：两个字段:name表示项目名字,DataBase表示项目用的数据库，定义一个方法：testDataBase(),测试DataBase的4个方法
	homeworkFunc_04()
}
func homeworkFunc_04()  {
	fmt.Println("******04*****")
	p := Project{}
	p.name = "安辰通"
	p.testDataBase(model.Oracle{})

	p2 := Project{}
	p2.name = "衣家洗衣"
	p2.testDataBase(model.Mysql{})
}
func homewordFunc_03()  {
	y := model.YAMAHA{}
	y.Start()
	y.Speedup()
	y.Stop()

	h := model.HONDA{}
	h.Start()
	h.Speedup()
	h.Stop()
	fmt.Println("***测试改小汽车的发动机***")
	c := model.Car{}
	c.TestIEngine(h)

}