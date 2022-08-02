package core

type Cart struct {
	order *Order
}

type Fee map[string]float64

type CartTotals struct {
	Subtotal float64
	Fees     []*Fee
	Total    float64
}
type FeeCalculator interface {
	Calculate()
}

func (t *Cart) AddOrderItem(item Modelable, quantity float64) {
	i := item.(*Variations)
	t.order.SaleItems[i] = quantity
}

func (t *Cart) RemoveOrderItem(item Modelable) {
	i := item.(*Variations)
	delete(t.order.SaleItems, i)
}

func (t *Cart) EditOrderItem(item Modelable, quantity float64) {
	i := item.(*Variations)
	t.order.SaleItems[i] = quantity
}

func NewCart(order *Order) *Cart {
	return &Cart{order: order}
}
