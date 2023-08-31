package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"testing"
)

const (
	api_host = "127.0.0.1"
	port     = "9005"

	accountName = "leebai"

	ProductName          = "productName"
	ProductNumber        = "productNumber"
	PayMethod            = "payMethod"
	DeliveryMethod       = "deliveryMethod"
	ExpectedDeliveryTime = "expectedDeliveryTime"
)

func TestSendCreateOrderRequest(t *testing.T) {
	url := fmt.Sprintf("http://%v:%v/commitOrder", api_host, port)

	resp, err := resty.New().R().
		SetHeader("AccountName", accountName).
		SetBody(map[string]interface{}{
			ProductName:          "iphone13 pro",
			ProductNumber:        2,
			PayMethod:            1,
			DeliveryMethod:       1,
			ExpectedDeliveryTime: "2023-01-02,09:00-21:00",
		}).Post(url)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("resp: %v\n", resp)
}
