package types

type OrderItem struct {
	ProductName string
	Picture     string
	Cost        float32
	Qty         uint32
}

type Order struct {
	OrderId     string
	CreatedDate string
	Cost        float32
	Items       []OrderItem
}
