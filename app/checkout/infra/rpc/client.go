package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/suutest/app/checkout/conf"
	"github.com/suutest/common/clientsuite"
	"github.com/suutest/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/suutest/rpc_gen/kitex_gen/order/orderservice"
	"github.com/suutest/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/suutest/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/suutest/rpc_gen/kitex_gen/stock/stockservice"
)

var (
	CartClient    cartservice.Client
	PaymentClient paymentservice.Client
	ProductClient productcatalogservice.Client
	OrderClient   orderservice.Client
	StockClient   stockservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func Init() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
		initStockClient()
	})
}

func initCartClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		panic(err)
	}
}

func initPaymentClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		panic(err)
	}
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		panic(err)
	}
}

func initOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		panic(err)
	}
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
