package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email   string
	Street  string
	City    string
	State   string
	Country string
	Phone   string
	ZipCode string
}

type Order struct { // 定义订单的表结构
	gorm.Model
	OrderId      string      `gorm:"type:varchar(100);uniqueIndex"`
	UserId       uint32      `gorm:"type:int(11)"`
	UserCurrenty string      `gorm:"type:varchar(10)"`
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

func (Order) TableName() string {
	return "order"
}

func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) ([]*Order, error) {
	var orders []*Order
	err := db.WithContext(ctx).Where("user_id=?", userId).Preload("OrderItems").Find(&orders).Error // 预加载这里写的是字段名
	if err != nil {
		return nil, err
	}
	return orders, nil
}
