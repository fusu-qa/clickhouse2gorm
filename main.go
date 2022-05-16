package main

import (
	"clickhouse2gorm/gen"
	"fmt"
)

func main() {
	ip := "192.168.100.200"
	port := "9000"
	dbName := "db"
	username := "userName"
	password := "password"
	dsn := fmt.Sprintf("tcp://%s:%s?database=%s&username=%s&password=%s&read_timeout=10&write_timeout=20", ip, port, dbName, username, password)
	//生成指定单表
	tblName := "tableName"
	err := gen.GenerateOne(gen.CHGenConf{
		Dsn:       dsn,
		WritePath: "./model",
		Stdout:    false,
		Overwrite: true,
	}, dbName, tblName)
	if err != nil {
		return
	}
}
