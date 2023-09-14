package main

import (
	"log"
	"time"

	"github.com/go-redis/redis/v7"
)

// RedisConfig redis配置
type RedisConfig struct {
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`

	MaxRetries      int           `json:"max_retries" yaml:"max_retries"`
	MinRetryBackoff time.Duration `json:"min_retry_backoff" yaml:"min_retry_backoff"`
	MaxRetryBackoff time.Duration `json:"max_retry_backoff" yaml:"max_retry_backoff"`

	DialTimeout  time.Duration `json:"dial_timeout" yaml:"dial_timeout"`
	ReadTimeout  time.Duration `json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout" yaml:"write_timeout"`

	PoolSize           int           `json:"pool_size" yaml:"pool_size"`
	MinIdleConns       int           `json:"min_idle_conns" yaml:"min_idle_conns"`
	MaxConnAge         time.Duration `json:"max_conn_age" yaml:"max_conn_age"`
	PoolTimeout        time.Duration `json:"pool_timeout" yaml:"pool_timeout"`
	IdleTimeout        time.Duration `json:"idle_timeout" yaml:"idle_timeout"`
	IdleCheckFrequency time.Duration `json:"idle_check_frequency" yaml:"idle_check_frequency"`

	// 触发hgetall等长耗时命令转为scan命令的key数量下限
	ScanThreshold int `json:"scan_threshold" yaml:"scan_threshold"`

	// redis cluster do not support select db
	// DB int `json:"db" yaml:"db"`
}

func NewRedisClient(conf RedisConfig) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:               conf.Addr,
		Password:           conf.Password,
		MaxRetries:         conf.MaxRetries,
		MinRetryBackoff:    conf.MinRetryBackoff,
		MaxRetryBackoff:    conf.MaxRetryBackoff,
		DialTimeout:        conf.DialTimeout,
		ReadTimeout:        conf.ReadTimeout,
		WriteTimeout:       conf.WriteTimeout,
		PoolSize:           conf.PoolSize,
		MinIdleConns:       conf.MinIdleConns,
		MaxConnAge:         conf.MaxConnAge,
		PoolTimeout:        conf.PoolTimeout,
		IdleTimeout:        conf.IdleTimeout,
		IdleCheckFrequency: conf.IdleCheckFrequency,

		//DB: conf.DB,
	})

	for i := 0; i < 3; i++ {
		if err := cli.Ping().Err(); err != nil {
			log.Printf("ping redis:%s err:%v\n", conf.Addr, err)
			time.Sleep(time.Millisecond * 100)
			continue
		}
		break
	}

	return cli
}
