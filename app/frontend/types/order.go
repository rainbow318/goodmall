package types

// 订单商品
type OrderItem struct {
	ProductName string
	Picture     string
	Qty         uint32
	Cost        float32
}

// 订单
type Order struct {
	OrderId     string
	CreatedDate string
	Cost        float32
	Items       []OrderItem
}
