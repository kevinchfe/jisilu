package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// GetDbConfig 获取config.yaml
func GetDbConfig() (string, string) {
	viper.SetConfigFile("./configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	host := viper.GetString("Database.Host")
	port := viper.GetString("Database.Prot")
	username := viper.GetString("Database.UserName")
	password := viper.GetString("Database.Password")
	database := viper.GetString("Database.DBName")
	charset := viper.GetString("Database.Charset")
	prefix := viper.GetString("Database.TablePrefix")
	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", username, password, host, port, database, charset)

	return mysqlDSN, prefix
}

// DbEngine 数据库配置表
func DbEngine() *gorm.DB {

	mysqlDSN, prefix := GetDbConfig()

	// 初始化数据库
	db, err := gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix, // table name prefix, table for `User` would be `t_users`
			SingularTable: true,   // use singular table name, table for `User` would be `user` with this option enabled
		},
		Logger: newLogger,
	})

	if err != nil {
		fmt.Println(err)
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10)
	return db
}

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold: time.Second,  // 慢 SQL 阈值
		LogLevel:      logger.Info, // Log level
		Colorful:      false,        // 禁用彩色打印
	},
)

var DB *gorm.DB = DbEngine()
