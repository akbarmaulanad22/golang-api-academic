package entity

type GradeComponent struct {
	Entity
	Name   string `gorm:"column:name"`
	Weight int    `gorm:"column:weight"`
}
