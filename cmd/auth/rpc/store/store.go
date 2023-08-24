package store

import "cshop/pkg/store"

func FindTokenByAccountName(accountName string) (string, error) {
	return store.Get(accountName)
}

func SaveTokenByAccountName(accountName string, token string, expireTimeSecond int) error {
	return store.Set(accountName, token, expireTimeSecond)
}
