package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"testing"
)

const (
	auth_api_host = "127.0.0.1"
	port          = "9003"
	register      = "/register"
	unregister    = "/unregister"
	login         = "/login"
	logout        = "/logout"

	accountNameKey = "accountName"
	passwordKey    = "password"

	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoie1wiYWNjb3VudE5hbWVcIjpcImxlZWJhaVwifSIsImV4cCI6MTY5NDA4MDA0Nn0.uzPqBT-de3nb1OhKDWqO5XUzGYdvwgKl4qCj4SCLKJQ"
)

func TestLogin(t *testing.T) {
	url := fmt.Sprintf("http://%v:%v%v", auth_api_host, port, login)

	resp, err := resty.New().R().
		SetHeader("Access-Token", token).
		SetBody(map[string]interface{}{accountNameKey: "leebai", passwordKey: "123456"}).
		Post(url)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("resp: %v\n", resp)
}
