package gorm

import (
	"cshop/pkg/store/gorm/model"
	"fmt"
	"testing"
)

const (
	DSN = "root:123456@tcp(127.0.0.1:3306)/go_zero_demo?charset=utf8mb4&parseTime=true"
)

func TestAdd(t *testing.T) {
	cfg := Config{
		DSN: DSN,
	}
	cli := MustNewMySQLClient(cfg)

	db := cli.GetDB()

	err := AddOneRecord(
		db,
		model.UserAccount{
			AccountName: "huangzhong",
			Password:    "123456",
		})
	fmt.Printf("AddOneRecord err: %v\n", err)
}

func TestDel(t *testing.T) {
	cfg := Config{
		DSN: DSN,
	}
	cli := MustNewMySQLClient(cfg)

	db := cli.GetDB()

	err := DelOneRecord(
		db,
		&model.UserAccount{},
		fmt.Sprintf("%v=?", model.AccountNameColumn),
		"huangzhong")
	fmt.Printf("DelOneRecord err: %v\n", err)
}

func TestQuery(t *testing.T) {
	cfg := Config{
		DSN: DSN,
	}
	cli := MustNewMySQLClient(cfg)

	db := cli.GetDB()

	userAccount := &model.UserAccount{}
	err := QueryOneRecord(
		db,
		userAccount,
		fmt.Sprintf("%v=?", model.AccountNameColumn),
		"huangzhong")
	fmt.Printf("QueryOneRecord err: %v, userAccount: %#v\n", err, userAccount)
}

func TestUpdate(t *testing.T) {
	cfg := Config{
		DSN: DSN,
	}
	cli := MustNewMySQLClient(cfg)

	db := cli.GetDB()

	err := UpdateOneColumn(
		db,
		&model.UserAccount{},
		model.PasswordColumn,
		"666888",
		fmt.Sprintf("%v=?", model.AccountNameColumn),
		"huangzhong")
	fmt.Printf("UpdateOneColumn err: %v\n", err)
}

func TestUpdateWhole(t *testing.T) {
	cfg := Config{
		DSN: DSN,
	}
	cli := MustNewMySQLClient(cfg)

	db := cli.GetDB()

	err := UpdateOneRecord(
		db,
		model.UserAccount{
			AccountName: "huangzhong",
			Password:    "654321",
		})
	fmt.Printf("UpdateOneRecord err: %v\n", err)
}

func TestMakeSureDBTableExist(t *testing.T) {
	// 执行本函数前确保TbUserAccountTest不存在，调用sql: "drop table TbUserAccountTest;"先删除原来存在的表

	cfg := Config{
		DSN: DSN,
	}
	cli := MustNewMySQLClient(cfg)

	db := cli.GetDB()

	err := MakeSureDBTableExist(db, model.UserAccountTest{})
	fmt.Printf("MakeSureDBTableExist err: %v\n", err)
}
