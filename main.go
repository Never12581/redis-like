package main

import (
	"redis-like/config"
	"redis-like/protocol"
	"redis-like/storage"
)

func main() {
	// 环境变量配置
	config.EnvConfigInstance()
	// 存储引擎初始化
	storage.StorageInstance()
	// 网络初始化
	protocol.Start()
}
