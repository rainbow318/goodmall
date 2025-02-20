package consumer

import "github.com/suutest/app/email/biz/consumer/email"

// 统一初始化所有consumer
func Init() {
	email.ConsumerInit()
}
