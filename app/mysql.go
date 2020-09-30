package app

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
	"sync"
	"log"
)

var (
	gor *gorm.DB
	once sync.Once
)

func LoadDB() *gorm.DB {
	once.Do(func() {
		dbConnectBaseStr := "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local"
		dbConnect := fmt.Sprintf(dbConnectBaseStr,
			Config.Get("db_user"),
			Config.Get("db_pass"),
			Config.Get("db_host"),
			Config.Get("db_port"),
			Config.Get("db_name"),
			Config.Get("db_charset"))
		//log.Printf("mysql connect %s",dbConnect)
		db, err := gorm.Open("mysql", dbConnect)
		log.Println(db)
		if err != nil {
			fmt.Println(err)
		}
		db.LogMode(true)
		//db.LogMode(true)
		//db.SetLogger(Log) // TODO 指定日志组件
		//fmt.Println("连接成功。")
		db.DB().Ping()
		// SetMaxIdleCons 设置连接池中的最大闲置连接数。
		db.DB().SetMaxIdleConns(10)
		// SetMaxOpenCons 设置数据库的最大连接数量。
		db.DB().SetMaxOpenConns(50)
		// SetConnMaxLifetiment 设置连接的最大可复用时间。
		//db.DB().SetConnMaxLifetime(time.Hour)
		db.DB().SetConnMaxLifetime(time.Second * 10) // 连接过期时间; 如果不设置 连接会一直不释放
		//defer db.Close()
		gor = db
	})
	return  gor
}
