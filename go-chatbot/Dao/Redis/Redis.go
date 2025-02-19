package Redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client

// 初始化Redis客户端并连接到Redis服务器
func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = rdb.Close()
}
