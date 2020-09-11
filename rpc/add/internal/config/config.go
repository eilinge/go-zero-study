package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rpcx"
)

type Config struct {
	rpcx.RpcServerConf
	DataSource string          // 手动代码
	Table      string          // 手动代码
	Cache      cache.CacheConf // 手动代码
}
