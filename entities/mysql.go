package entities

type Credentials struct {
	CredentialId uint   `gorm:"primary_key"`
	AccountId    uint   `gorm:"column:account_id"`
	PasswordHash string `gorm:"column:password_hash"`
	Salt         string `gorm:"column:salt"`
}
