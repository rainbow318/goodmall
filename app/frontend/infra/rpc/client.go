package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/suutest/app/frontend/conf"
	frontendUtils "github.com/suutest/app/frontend/utils"
	"github.com/suutest/common/clientsuite"
	"github.com/suutest/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/suutest/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/suutest/rpc_gen/kitex_gen/order/orderservice"
	"github.com/suutest/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/suutest/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient     userservice.Client // 用来存放user服务的RPC客户端的一个实例
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once // 保证它只会被初始化一次
	ServiceName    = frontendUtils.ServiceName
	RegistryAddr   = conf.GetConf().Hertz.RegistryAddr
	err            error
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

// 具体的初始化用户服务的函数逻辑写在这里
func initUserClient() {
	UserClient, err = userservice.NewClient("user", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	})) // 使用rpc_gen模块提供的代码来生成一个PRC客户端
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}
