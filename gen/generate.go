package gen

import (
	"clickhouse2gorm/db"
	"clickhouse2gorm/util"
	"fmt"
	"gorm.io/gorm"
	"os"
	"strings"
)

type CHGenConf struct {
	Dsn       string //数据库配置
	WritePath string //生成文件路径
	Stdout    bool   //true时只输出至标准输出
	Overwrite bool   //true则覆盖原文件
}

func (g CHGenConf) isValid() error {
	if g.Dsn == "" {
		return fmt.Errorf("dsn is empty")
	}

	if g.Stdout == false {
		if g.WritePath == "" {
			return fmt.Errorf("when Stdout is false WritePath can't be empty")
		}

		if util.PathExists(g.WritePath) == false {
			os.MkdirAll(g.WritePath, 0766)
		}
	}

	return nil
}

func GenerateOne(conf CHGenConf, dbName string, tblName string) (err error) {
	err = conf.isValid()
	if err != nil {
		fmt.Printf("GenerateOne for db:%s table:%s err:%s\n", dbName, tblName, err)
		return err
	}

	db, err := db.InitCH(conf.Dsn)
	if err != nil {
		return fmt.Errorf("ch init err:%s", err)
	}

	err = doGenerateOne(db, conf, dbName, tblName)
	if err != nil {
		fmt.Printf("GenerateOne for table %s err:%s\n", tblName, err)
	}

	return nil
}

func doGenerateOne(db *gorm.DB, conf CHGenConf, dbName string, tblName string) (err error) {
	info, err := GetTableInfo(db, dbName, tblName)
	builder := strings.Builder{}
	builder.WriteString("package model")
	builder.WriteString("\n")
	builder.WriteString("\n")
	builder.WriteString("type ")
	builder.WriteString(util.StrCamel(tblName) + " ")
	builder.WriteString("struct {")
	builder.WriteString("\n")
	for _, columnInfo := range info.ColumnsInfo {
		builder.WriteString("\t")
		pName := util.StrCamel(columnInfo.Field)
		builder.WriteString(pName + " ")
		builder.WriteString(columnInfo.Type + " ")
		builder.WriteString(fmt.Sprintf("`json:\"%s\"`", util.StrFirstToLower(pName)))
		builder.WriteString("\n")
	}
	builder.WriteString("}")

	if conf.Stdout {
		fmt.Println(builder.String())
		return nil
	}

	//输出至文件
	path, err := mkDir(conf.WritePath, "")
	if err != nil {
		return fmt.Errorf("mkdir %s err:%s", path, err)
	}

	fName := mkFileName(path, info.TableName)
	if util.PathExists(fName) && conf.Overwrite == false {
		return fmt.Errorf("file :%s is exists", fName)
	}

	os.WriteFile(fName, []byte(builder.String()), 0666)

	fmt.Printf("Generate Table:%s to %s success!\n", info.TableName, fName)

	return nil
}

//mkFileName 拼接路径和文件名
func mkFileName(path string, name string) string {
	path = strings.TrimRight(path, "/")
	return fmt.Sprintf("%s/%s.go", path, name)
}

//mkDir 创建并返回文件所在目录
func mkDir(base string, tblDir string) (string, error) {
	fullDir := fmt.Sprintf("%s/%s", base, tblDir)

	err := os.MkdirAll(fullDir, 0766)
	return fullDir, err
}
