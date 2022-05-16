# 1. 功能
根据clickhouse数据库表生成gorm需要的struct。支持指定单表生成

比如有如下数据表：
```clickhouse
CREATE TABLE app_count
(
    `cluster_id` Int32,
    `cluster_name` String,
    `running_num` Int32,
    `pending_num` Int32,
    `failed_num` Int32,
    `collected_at` Int64,
    `partition` String
)
    ENGINE = MergeTree
    PARTITION BY partition
    ORDER BY (cluster_name,
              collected_at)
    SETTINGS index_granularity = 8192;
```

clickhouse2gorm可以在指定的目录下生成 model/app_count.go，内容如下：

```
package model

type AppCount struct {
	ClusterID int32 `json:"clusterID"`
	ClusterName string `json:"clusterName"`
	RunningNum int32 `json:"runningNum"`
	PendingNum int32 `json:"pendingNum"`
	FailedNum int32 `json:"failedNum"`
	CollectedAt int64 `json:"collectedAt"`
	Partition string `json:"partition"`
}
```

# 2. 使用
## 2.1 指定单表生成文件
```go
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
```

gen.GenConf的说明如下：
- Dsn：数据库配置
- WritePath：指定文件写入的目录。不存在会自动创建
- Stdout：是否输出至标准输出。如果Stdout为true，则生成的struct不会写入文件。
- Overwrite：当原文件存在时，是否进行覆盖。true为覆盖。

# 4. 生成规则
## 4.1 表名
表名对应的大驼峰命名做为struct名，表名做为文件名。
比如表名为demo_test, 则：
- struct名：DemoTest
- 文件名：demo_test.go
- 包名：model


## 4.2 struct字段名
表字段对应的大驼峰命名做为struct字段名。比如is_admin对应IsAdmin。但id字段除名，id固定对应ID。
gorm对应struct字段名首字母小写



