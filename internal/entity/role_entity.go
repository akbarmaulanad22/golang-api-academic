package entity

type Role struct {
	Entity
	Name string `gorm:"column:name"`
}
