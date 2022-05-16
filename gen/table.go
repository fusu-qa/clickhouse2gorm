package gen

import (
	"github.com/qmhball/db2gorm/util"
	"gorm.io/gorm"
)

// Tables 单个原始表名集合
type Tables []string

// GetTables 获取DB下所有的表名
func (t *Tables) GetTables(orm *gorm.DB) error {
	res := orm.Raw("show tables").Scan(t)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// TableInfo 单个表在生成struct时所需的全部信息
type TableInfo struct {
	TableName   string //原始表名
	StructName  string //驼峰表名
	ColumnsInfo []ColumnInfo
}

func GetTableInfo(db *gorm.DB, dbName string, tblName string) (TableInfo, error) {
	var i TableInfo
	i.TableName = tblName
	i.StructName = util.StrCamel(tblName)
	info, err := GetTableColumnsInfo(db, dbName, tblName)
	if err != nil {
		return i, err
	}

	i.ColumnsInfo = info

	return i, nil
}
