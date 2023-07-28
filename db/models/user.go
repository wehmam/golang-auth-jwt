package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(300)" json:"username"`
	Password string `gorm:"type:varchar(300) json:password"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by Indicators to `indicator` (Plural problems)
func (User) TableName() string {
	return "users"
}
