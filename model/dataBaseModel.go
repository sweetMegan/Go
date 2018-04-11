package model

import "fmt"

type DataBase interface {
	Insert(target string)
	Update(target string)
	Query(target string)
	Delete(target string)
}
type Mysql struct {

}

func (sql Mysql)Insert(target string)  {
	fmt.Println("向",target," 插入sql数据")
}
func (sql Mysql)Update(target string){
	fmt.Println("sql 更新数据")
}
func (sql Mysql)Query(target string)  {
	fmt.Println("sql 查询数据")
}
func (sql Mysql)Delete(target string){
	fmt.Println("sql 删除数据")
}
type Oracle struct {

}
func (ora Oracle)Insert(target string)  {
	fmt.Println("向",target,"ora 插入数据")
}
func (ora Oracle)Update(target string){
	fmt.Println("ora 更新数据")
}
func (ora Oracle)Query(target string)  {
	fmt.Println("ora 查询数据")
}
func (ora Oracle)Delete(target string){
	fmt.Println("ora 删除数据")
}
