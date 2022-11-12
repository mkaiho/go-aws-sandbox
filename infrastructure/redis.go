package infrastructure

import (
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/kelseyhightower/envconfig"
	redisadapter "github.com/mkaiho/go-aws-sandbox/adapter/redis"
)

// Redis Sentinel Configuration
var _ redisadapter.SentinelConfig = (*RedisSentinelConfig)(nil)

func LoadRedisSentinelConfig() (*RedisSentinelConfig, error) {
	var c redisSentinelEnvConfig
	if err := envconfig.Process("REDIS", &c); err != nil {
		return nil, err
	}
	return &RedisSentinelConfig{
		masterName:    c.MasterName,
		sentinelAddrs: strings.Split(c.SentinelAddrs, ","),
	}, nil
}

type RedisSentinelConfig struct {
	masterName    string
	sentinelAddrs []string
}

func (c *RedisSentinelConfig) MasterName() string {
	return c.masterName
}

func (c *RedisSentinelConfig) SentinelAddrs() []string {
	return c.sentinelAddrs
}

type redisSentinelEnvConfig struct {
	MasterName    string `envconfig:"MASTER_NAME" required:"true"`
	SentinelAddrs string `envconfig:"SENTINEL_ADDRS" required:"true"`
}

func OpenRedisSentinel(conf redisadapter.SentinelConfig) (*redis.Client, error) {
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    conf.MasterName(),
		SentinelAddrs: conf.SentinelAddrs(),
	})
	return client, nil
}
