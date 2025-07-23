package database

import (
	"fmt"
	"time"

	"go-starter/core/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// MySQL 全局数据库实例
var MySQL *gorm.DB

// InitDatabase 初始化数据库连接
func InitDatabase() error {
	mysqlConfig := config.GetConfig().Mysql
	if !mysqlConfig.Enable {
		return nil
	}
	// 1. 构建数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Database,
	)

	// 2. 配置 GORM
	config := &gorm.Config{
		// 命名策略：使用单数表名
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		// 禁用默认事务
		DisableAutomaticPing: false,
		// 准备语句
		PrepareStmt: true,
	}

	// 3. 连接数据库
	var err error
	MySQL, err = gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// 4. 配置连接池
	sqlDB, err := MySQL.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(mysqlConfig.PoolSize)    // 最大空闲连接数
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxPoolSize) // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour)            // 连接最大生存时间

	// 5. 测试连接
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

// CloseDatabase 关闭数据库连接
func CloseDatabase() error {
	if MySQL != nil {
		sqlDB, err := MySQL.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
