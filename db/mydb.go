package db

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

// InitCH 初始化Clickhouse
func InitCH(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
