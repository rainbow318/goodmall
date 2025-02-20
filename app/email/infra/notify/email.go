package notify

import (
	"fmt"

	"github.com/kr/pretty"
	"github.com/suutest/rpc_gen/kitex_gen/email"
)

type NoopEmail struct{}

// 模拟邮件发送
func (e *NoopEmail) Send(req *email.EmailReq) error {
	// 这里的处理我们只是简单打印，而不是真的发邮件
	pretty.Printf("%v\n", req)
	fmt.Println("print")
	return nil
}

func NewNoopEmail() NoopEmail {
	return NoopEmail{}
}
