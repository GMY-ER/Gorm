package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
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

	_ = db.AutoMigrate(&User{}) //_=db.AutoMigrate(new(Product))
	user := User{
		Name: "bobby2",
	}
	result := db.Create(&user)
	fmt.Println(user.ID)             //返回插入数据的主键
	fmt.Println(result.Error)        //返回error
	fmt.Println(result.RowsAffected) //返回插入记录的条数
	//db.Create(&User{Name: "bobby"}) //只为我的name字段赋值，其他的都是Go初试默认值
	//db.Model(&User{ID: 1}).Update("Name", "")
	//updates语句不会更新零值,但是update会更新
	//empty := ""
	//db.Model(&User{ID: 1}).Updates(User{Email: &empty})
	//解决仅更新非零值有两种
	/*
		将string 设置为 *string
		使用sql 的NULLxxx来解决
	*/
}
