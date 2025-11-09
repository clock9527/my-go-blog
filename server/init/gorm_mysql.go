package init

import (
	"log"
	"my-go-blog/server/config"
	"my-go-blog/server/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化Mysql数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ByteZhou-2018](https://github.com/ByteZhou-2018)
func GormMysql() *gorm.DB {
	var dbMysql = config.DBMysql{
		GeneralDB: config.GeneralDB{
			Path:     "127.0.0.1",
			Port:     "3306",
			Dbname:   "go_blog",
			Config:   "charset=utf8mb4&parseTime=True&loc=Local",
			Username: "root",
			Password: "123456",
		},
	}
	db, err := gorm.Open(mysql.Open(dbMysql.Dsn()), &gorm.Config{})
	if err != nil {
		log.Println("server.init.GormMysql-数据库连接失败-" + err.Error())
		return nil
	}
	return db
}

func SetDB() {
	global.GVA_DB = GormMysql()
}
