package model

const (
	AccountNameColumn = "accountName"
	PasswordColumn    = "password"
)

type UserAccount struct {
	AccountName string `gorm:"primarykey;column:accountName;not null"`
	Password    string `gorm:"column:password"`
}

// TableName customize table name
func (u *UserAccount) TableName() string {
	return "TbUserAccount"
}
