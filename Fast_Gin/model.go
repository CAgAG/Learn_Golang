package main

import (
	"Gin_go/models"
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func NewDB() (*gorm.DB, error) {
	dsn := "test_root:123456@tcp(127.0.0.1:3306)/gin_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gin_test_",
			SingularTable: true,
		},
		// 打印对应的 sql 语句
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{}) // 自动同步

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(time.Hour) // 设置连接存活时间
	return db, nil
}

// ========================================
// redis
type RedisClient struct {
	DB *redis.Client
}

func NewRedis() (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{DB: rdb}, nil
}

func (rc *RedisClient) Get(key string) (any, error) {
	return rc.DB.Get(context.Background(), key).Result()
}

func (rc *RedisClient) Set_And_Keep(key string, value any, keep_time time.Duration) error {
	return rc.DB.Set(context.Background(), key, value, keep_time).Err()
}

func (rc *RedisClient) Set(key string, value any) error {
	return rc.Set_And_Keep(key, value, 1*time.Hour) // 默认保持 1小时
}

func (rc *RedisClient) Del(key ...string) error {
	if len(key) == 0 {
		return nil
	}
	return rc.DB.Del(context.Background(), key...).Err()
}

func (rc *RedisClient) Exist(key string) (bool, error) {
	exists, err := rc.DB.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}

	if exists != 0 {
		return true, nil
	} else {
		return false, nil
	}
}
