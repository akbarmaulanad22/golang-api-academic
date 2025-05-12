package entity

type Role struct {
	Model
	Name string `gorm:"column:name"`
}
