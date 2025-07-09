package sql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//func (User) TableName() string {
//	return "user"
//}

var DB *gorm.DB

// ConnectMysql 连接数据库
func ConnectMysql() {
	fmt.Println("尝试连接数据库当中。。。")

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:00-Halloworld-00@tcp(rm-bp18lj5pz4ymf76ja7o.mysql.rds.aliyuncs.com:3306)/cloud_platform?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})

	if err != nil {
		fmt.Println("====连接失败====", err.Error())
	} else {
		fmt.Println("====连接成功====", DB)
	}
}
