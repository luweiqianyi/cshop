package store

import (
	"cshop/pkg/store/redis"
)

func FindTokenByAccountName(accountName string) (string, error) {
	return redis.Get(accountName)
}

func SaveTokenByAccountName(accountName string, token string, expireTimeSecond int) error {
	return redis.Set(accountName, token, expireTimeSecond)
}

func DeleteTokenByAccountName(accountName string) error {
	return redis.Delete(accountName)
}
