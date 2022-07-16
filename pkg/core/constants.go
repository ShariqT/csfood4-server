package core

type UnitMeasurement int

type Weekday int

const (
	BY_UNIT  UnitMeasurement = iota
	BY_POUND UnitMeasurement = iota
)

const (
	MONDAY    Weekday = 1
	TUESDAY   Weekday = 2
	WEDNESDAY Weekday = 3
	THURSDAY  Weekday = 4
	FRIDAY    Weekday = 5
	SATURDAY  Weekday = 6
	SUNDAY    Weekday = 7
)

type UserStatus int

const (
	VENDOR   UserStatus = iota
	CUSTOMER UserStatus = iota
)

type OrderProcessingStatus int

const (
	ORDER_APPROVED  OrderProcessingStatus = iota
	ORDER_CANCELLED OrderProcessingStatus = iota
	ORDER_REFUNDED  OrderProcessingStatus = iota
)
