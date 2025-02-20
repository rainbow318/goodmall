package email

import (
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"github.com/suutest/app/email/infra/mq"
	"github.com/suutest/app/email/infra/notify"
	"github.com/suutest/rpc_gen/kitex_gen/email"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	// 先订阅一个主题

	sub, err := mq.Nc.Subscribe("email", func(m *nats.Msg) { // 定义一个消费方法
		var req email.EmailReq
		fmt.Println("here0")
		fmt.Printf("%+v\n", m.Data)
		// 消息格式是protobuf 所以要先反序列化收到的消息
		err := proto.Unmarshal(m.Data, &req)
		fmt.Println("here1")
		if err != nil {
			klog.Error(err)
		}
		// 如果反序列化成功，就发送邮件
		noopEmail := notify.NewNoopEmail()
		_ = noopEmail.Send(&req)
		fmt.Println("here2")
	})
	if err != nil { // 如果订阅失败：
		panic(err)
	}
	// 在服务关闭时要取消订阅
	server.RegisterShutdownHook(func() {
		sub.Unsubscribe()
		mq.Nc.Close()
	})
}
