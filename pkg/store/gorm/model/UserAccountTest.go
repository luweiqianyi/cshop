package model

type UserAccountTest struct {
	AccountName string `gorm:"primarykey;column:accountName;not null"`
	Password    string `gorm:"column:password"`
}

// TableName customize table name
func (u *UserAccountTest) TableName() string {
	return "TbUserAccountTest"
}
