package sql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

type MySQL struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

// ConnectionPools 设置数据库连接池的参数，并返回 gorm.DB 对象
// sql: 需要连接的 MySQL 实例对象
// MaxLink: 连接池中最大连接数，即数据库允许的最大活动连接数
// MaxIdle: 连接池中最大空闲连接数，即保持空闲的最大连接数
// MaxTime: 连接的最大生命周期，超过此时间后，连接会被关闭并重建
func ConnectionPools(sql *MySQL, MaxLink int, MaxIdle int, MaxTime time.Duration) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", sql.User, sql.Password, sql.Host, sql.Port, sql.Database)
	var once sync.Once
	once.Do(func() {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("SQL连接出错(There was an error with the SQL connection)", err)
		}
		//配置连接池参数
		SQL, err := db.DB()

		if err != nil {
			log.Println("无法获取服务器实例(failed to get DB instance)", err)
		}
		//设置最大连接数
		SQL.SetMaxOpenConns(MaxLink)
		//设置最大空闲数
		SQL.SetMaxIdleConns(MaxIdle)
		//设置最大生存时间
		SQL.SetConnMaxLifetime(MaxTime)
	})
	log.Println("MySQL连接成功")
	return db, nil
}
