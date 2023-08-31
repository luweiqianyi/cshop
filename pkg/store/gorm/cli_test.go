package gorm

import (
	"fmt"
	"testing"
)

func TestMustNewMySQLClient(t *testing.T) {
	cfg := Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/go_zero_demo?charset=utf8mb4&parseTime=true",
	}
	cli := MustNewMySQLClient(cfg)
	fmt.Printf("cli: %#v\ndb: %#v\n", cli, cli.GetDB())

}
