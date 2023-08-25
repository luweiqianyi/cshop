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
)

func TestRegister(t *testing.T) {
	url := fmt.Sprintf("http://%v:%v%v", auth_api_host, port, register)

	resp, err := resty.New().R().
		SetBody(map[string]interface{}{accountNameKey: "leebai", passwordKey: "123456"}).
		Post(url)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("resp: %v\n", resp)
}

func TestUnRegister(t *testing.T) {
	url := fmt.Sprintf("http://%v:%v%v", auth_api_host, port, unregister)

	resp, err := resty.New().R().
		SetQueryParams(map[string]string{accountNameKey: "leebai"}).
		Get(url)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("resp: %v\n", resp)
}

func TestLogin(t *testing.T) {
	url := fmt.Sprintf("http://%v:%v%v", auth_api_host, port, login)

	resp, err := resty.New().R().
		SetBody(map[string]interface{}{accountNameKey: "leebai", passwordKey: "123456"}).
		Post(url)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("resp: %v\n", resp)
}

func TestLogout(t *testing.T) {
	url := fmt.Sprintf("http://%v:%v%v", auth_api_host, port, logout)

	resp, err := resty.New().R().
		SetQueryParams(map[string]string{accountNameKey: "leebai"}).
		Get(url)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("resp: %v\n", resp)
}
