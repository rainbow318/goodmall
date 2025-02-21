package rpc

import (
	"context"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	consulclient "github.com/kitex-contrib/config-consul/client"
	"github.com/kitex-contrib/config-consul/consul"
	"github.com/suutest/app/frontend/conf"
	frontendUtils "github.com/suutest/app/frontend/utils"
	"github.com/suutest/common/clientsuite"
	"github.com/suutest/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/suutest/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/suutest/rpc_gen/kitex_gen/order/orderservice"
	"github.com/suutest/rpc_gen/kitex_gen/product"
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
	// 给商品服务增加熔断策略
	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string { // 构建服务力度
		return circuitbreak.RPCInfo2Key(ri)
	})
	cbs.UpdateServiceCBConfig("frontend/product/GetProduct",
		circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2},
	)
	consulClient, err := consul.NewClient(consul.Options{})
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}), client.WithCircuitBreaker(cbs), client.WithFallback(
		fallback.NewFallbackPolicy( // 如果product服务down了的话，就会进入这个fallback，并展示我们在下面预定义的商品
			fallback.UnwrapHelper(
				func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
					if err == nil {
						return resp, nil
					}
					methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
					if methodName != "ListProducts" {
						return resp, err
					}
					return &product.ListProductsResp{
						Products: []*product.Product{
							{
								Price:       11999,
								Id:          3,
								Picture:     "/static/image/laptop-3.avif",
								Name:        "15寸 MacBook Air M3",
								Description: "笔记本电脑 laptop Apple/苹果 15 英寸 MacBook Air Apple M3 芯片 8 核中央处理器 10 核图形处理器 16GB 统一内存 512GB 固态硬盘",
							},
						},
					}, nil
				}),
		),
	),
		client.WithSuite(consulclient.NewSuite("product", ServiceName, consulClient)),
	)
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
