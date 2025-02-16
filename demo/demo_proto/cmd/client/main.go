package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/suuyh/demo/demo_proto/kitex_gen/pbapi"
	"github.com/suuyh/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/suuyh/demo/demo_proto/middleware"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500") // 创建一个服务发现解析器，连接到Consul服务注册中心（地址127.0.0.1:8500)
	if err != nil {
		log.Fatal(err)
	}
	c, err := echoservice.NewClient("demo_proto", client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithMiddleware(middleware.Middleware),
	) // 创建客户端
	if err != nil {
		log.Fatal(err)
	}
	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto client") // 使用metainfo构造一个带元信息的上下文
	res, err := c.Echo(ctx, &pbapi.Request{Message: "error"})                                     // 发起RPC调用，调用服务端的Echo方法
	var bizErr *kerrors.GRPCBizStatusError
	if err != nil {
		ok := errors.As(err, &bizErr)
		if ok {
			fmt.Printf("%#v", bizErr)
		}
		log.Fatal(err)
	}
	fmt.Printf("%v", res)
}
