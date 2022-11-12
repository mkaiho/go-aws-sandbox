package redis

type SentinelConfig interface {
	MasterName() string
	SentinelAddrs() []string
}
