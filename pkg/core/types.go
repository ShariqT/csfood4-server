package core

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ModelID string
}

func (m *Model) GenerateID() string {

	ID, err := uuid.NewUUID()
	if err != nil {
		return "failed-to-generate-id"
	}
	m.ModelID = ID.String()
	return m.ModelID
}

type Categoryable struct {
	Name string
}

type Product struct {
	Model
	Price             float64
	AvailableQuantity int
	Name              string
	Active            bool
	Description       string
	Photo             string
	Tags              []*Tag
	Vendor            *Vendor
	Variations        []*Variations
	Categories        []*ProductCategory
}

type Tag struct {
	Categoryable
}

type VendorCategory struct {
	Categoryable
}

type Market struct {
	Model
	Name                             string
	Photo                            string
	WeekdayOpen                      Weekday
	HoursOpen                        string
	HoursForPickup                   string
	OrderCutoffDay                   Weekday
	MinimumPrice                     int
	InstructionsForWic               string
	InstructionsForVendorPickup      string
	InstructionsForAggregationPickup string
	Address                          string
	Location                         string
	ParkingInfo                      string
	WicInfo                          string
}

func NewMarketFromData(data map[string]interface{}) *Market {
	item := &Market{}
	item.ModelID = fmt.Sprintf("%v", data["ID"])
	item.Name = fmt.Sprintf("%v", data["Name"])
	item.Photo = fmt.Sprintf("%v", data["Photo"])
	return item

}

type Vendor struct {
	Name        string
	Photo       string
	Markets     []*Market
	Categories  []*VendorCategory
	Description string
	User        *User
}

type Variations struct {
	Model
	Name                string
	Price               float64
	Ingredients         Ingredients
	MaximumQuantity     int
	IsUnitOrPound       UnitMeasurement
	UnitTypeDescription string
}

type ProductCategory struct {
	Categoryable
}

type Ingredients string

type Order struct {
	Model
	OrderedBy           *User
	Market              *Market
	SaleItems           map[*Variations]float64
	TotalSalePrice      float64
	DeliveryAddress     string
	ZipCode             int
	IsOriginal          bool
	StripePaymentId     string
	TotalOrderPrice     float64
	UseSavedCard        bool
	StripeTransactionId string
	MaxCancellationDate time.Time
	CancellationReason  string
	StripeSourceId      string
	DeliveryMethod      string
	Vendor              *Vendor
	StripeRefundId      string
	OrderStatus         OrderProcessingStatus
	OrderKey            string
}

type User struct {
	Username            string
	Password            string
	Email               string
	Status              UserStatus
	StripeCustomerId    string
	ReadTOS             bool
	Phone               string
	StripePayMethodId   string
	StripeAccountStatus string
}

type OrderMessage struct {
	Message  string
	DateSent time.Time
	Order    *Order
	To       *User
	From     *User
}
