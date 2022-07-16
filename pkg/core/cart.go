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

func (t *Cart) AddOrderItem(item *Variations, quantity float64) {
	t.order.SaleItems[item] = quantity
}

func (t *Cart) RemoveOrderItem(item *Variations) {
	delete(t.order.SaleItems, item)
}

func (t *Cart) EditOrderItem(item *Variations) {

}

func NewCart(order *Order) *Cart {
	return &Cart{order: order}
}
