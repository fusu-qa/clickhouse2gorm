package gen

import (
	"clickhouse2gorm/util"
	"fmt"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

type Column struct {
	ColumnName string `json:"name" gorm:"column:name"`
	ColumnType string `json:"type" gorm:"column:type"`
}

type Columns []Column

//change2Info Column转换成ColumnInfo
func (cs Columns) change2Info() ([]ColumnInfo, error) {
	num := len(cs)
	csInfo := make([]ColumnInfo, num)

	for idx, one := range cs {
		if err := csInfo[idx].set(one); err != nil {
			return nil, err
		}
	}

	return csInfo, nil
}

//getTableColumns 获取表的列和类型
func (cs *Columns) getTableColumns(db *gorm.DB, dbName string, tblName string) error {
	sql := fmt.Sprintf("select name,type from system.columns where database='%s' and table='%s'", dbName, tblName)
	res := db.Raw(sql).Scan(cs)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

type ColumnInfo struct {
	Field string
	Type  string
}

func (c *ColumnInfo) set(src Column) (err error) {
	c.setField(src.ColumnName)
	if err := c.setType(src.ColumnType); err != nil {
		return err
	}
	return nil
}

func (c *ColumnInfo) setField(name string) {
	tmp := strings.Split(name, "_")
	for _, v := range tmp {
		if v == "id" {
			c.Field += "ID"
		} else {
			c.Field += util.StrFirstToUpper(v)
		}
	}
}

//setType sql中的type与go的类型对应
func (c *ColumnInfo) setType(typeName string) error {
	//精确匹配
	if v, ok := TypeCHDicMp[typeName]; ok {
		c.Type = v
		return nil
	}

	// Fuzzy Regular Matching.模糊正则匹配
	for _, l := range TypeCHMatchList {
		if ok, _ := regexp.MatchString(l.Key, typeName); ok {
			c.Type = l.Value
			return nil
		}
	}

	return fmt.Errorf("no type for src typeName:%s", typeName)
}

// GetTableColumnsInfo 获取DB.tblName对应的表模板信息
func GetTableColumnsInfo(db *gorm.DB, dbName string, tblName string) ([]ColumnInfo, error) {
	var columns Columns
	if err := columns.getTableColumns(db, dbName, tblName); err != nil {
		return nil, err
	}

	columnsInfo, err := columns.change2Info()
	if err != nil {
		return nil, err
	}

	return columnsInfo, nil
}
