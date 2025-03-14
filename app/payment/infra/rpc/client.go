package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/suutest/app/payment/conf"
	"github.com/suutest/common/clientsuite"
	"github.com/suutest/rpc_gen/kitex_gen/stock/stockservice"
)

var (
	StockClient  stockservice.Client
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
	err          error
	once         sync.Once
)

func Init() {
	once.Do(func() {
		initStockClient()
	})
}

func initStockClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	StockClient, err = stockservice.NewClient("stock", opts...)
	if err != nil {
		panic(err)
	}
}
