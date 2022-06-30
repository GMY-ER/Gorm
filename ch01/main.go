package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	Code  sql.NullString
	Price uint
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(152.136.121.186:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	//设置全局的logger，这个logger在我们执行每个sql语句的时候会打印每一行的sql
	//sql才是最重要的，本着这个原则我尽量的给大家看到每个api背后的sql语句是什么
	//定义一个表结构-将表直接生成对应的表-migrations

	//迁移schema

	_ = db.AutoMigrate(&Product{}) //_=db.AutoMigrate(new(Product))

	//新增一条纪录
	db.Create(&Product{Code: sql.NullString{String: "D42", Valid: true}, Price: 100})
	//Read
	var product Product
	db.First(&product, 1)              //默认按整型的主键查找
	db.First(&product, "Code=?", "a1") //按照指定的字段查找
	//Update-将product的price更新为200
	db.Model(&product).Update("Price", 200)
	//Update-更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: sql.NullString{String: "", Valid: true}})
	//db.Model(&product).Updates(Product{Price: 200}) //仅更新非零字段
	//如果我们去更新一个product只设置了price：200
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": 100})
	////Delete - 删除product
	//db.Delete(&product, 1)
}
