package hbase

import (
	"git.apache.org/thrift.git/lib/go/thrift"

	"time"
)

//p 连接池
var p Pool

//InitHbase 初始化hbase
func InitHbase(host, port string, pool int) error {
	var err error
	hostPort := host + ":" + port
	config := new(PoolConfig)
	config.InitialCap = pool / 2
	config.MaxCap = pool
	config.IdleTimeout = 5 * time.Minute
	config.Factory = func() (*THBaseServiceClient, error) {
		trans, err := thrift.NewTSocket(hostPort)
		if err != nil {
			return nil, err
		}

		trans.SetTimeout(1 * time.Minute)
		client := NewTHBaseServiceClientFactory(trans, thrift.NewTBinaryProtocolFactoryDefault())
		if err := trans.Open(); err != nil {
			return nil, err
		}
		return client, nil
	}
	config.Close = func(client *THBaseServiceClient) error {
		return client.Transport.Close()
	}
	p, err = NewChannelPool(config)
	if err != nil {
		return err
	}
	return nil
}

//OpenClient 打开客户端
func OpenClient() (*THBaseServiceClient, error) {
	return p.Get()
}

//CloseClient 关闭客户端
func CloseClient(client *THBaseServiceClient) error {
	return p.Put(client)
}
