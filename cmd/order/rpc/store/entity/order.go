package entity

// 用户下单后，订单初始状态就为“待支付”; 用户取消支付或者订单超时，订单状态变为“已取消”; 用户完成支付，订单状态变为“已完成”
// 如果有退款流程： TODO think about it
const (
	OrderCanceled     = 1 // 已取消
	OrderWaitForPay   = 2 // 待支付
	OrderComplemented = 3 // 已完成
	OrderRefunded     = 4 // 已退款
)

type Order struct {
	OrderID           string
	OrderCreatorID    string
	CreateTimestampMs int64
	OrderStatus       int

	OrderDetailInfo OrderDetailInfo
}

type OrderDetailInfo struct {
}
