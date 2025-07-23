package redis

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"go-starter/core/config"
	"go-starter/core/logger"

	"github.com/redis/go-redis/v9"
)

var (
	// Redis客户端实例
	redisClient *redis.Client
)

// InitRedis 初始化Redis连接
func InitRedis() error {
	conf := config.GetConfig().Redis
	if !conf.Enable {
		return nil
	}

	// 2. 创建Redis客户端
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.Database,

		// 连接池配置
		PoolSize:     10,
		MinIdleConns: 5,

		// 超时配置
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,

		// 重试配置
		MaxRetries:      3,
		MinRetryBackoff: 8 * time.Millisecond,
		MaxRetryBackoff: 512 * time.Millisecond,
	})

	// 3. 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		logger.SugaredLogger.Error("Redis连接失败", "error", err)
		return errors.New("Redis连接失败: " + err.Error())
	}

	logger.SugaredLogger.Info("Redis连接成功",
		" Addr:", conf.Host+":"+strconv.Itoa(conf.Port),
		" Database:", conf.Database)

	return nil
}

func Client() *redis.Client {
	return redisClient
}

// CloseRedis 关闭Redis连接
func CloseRedis() error {
	if redisClient != nil {
		if err := redisClient.Close(); err != nil {
			logger.SugaredLogger.Error("关闭Redis连接失败", "error", err)
			return err
		}
		logger.SugaredLogger.Info("Redis连接已关闭")
	}
	return nil
}
