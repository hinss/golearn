package proxy

import "fmt"

type DataSource interface {
	
	ExecuteSql()
}

type MysqlDataSource struct {
	
}

func (MysqlDataSource) ExecuteSql() {
	fmt.Println("执行mysql的sql的执行!")
}

type ProxyDataSource struct{
	proxy MysqlDataSource

}

func (p ProxyDataSource) ExecuteSql() {
	fmt.Println("开启事务。。。")
	p.proxy.ExecuteSql()
	fmt.Println("事务提交")
}

