package models

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

func (User) TableName() string {
	return "users"
}