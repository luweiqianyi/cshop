package entity

type Order struct {
	OrderID           string
	OrderCreatorID    string
	CreateTimestampMs int64

	OrderDetailInfo OrderDetailInfo
}

type OrderDetailInfo struct {
}
