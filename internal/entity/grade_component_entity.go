package entity

type GradeComponent struct {
	Entity
	Name   string  `gorm:"column:name"`
	Weight float64 `gorm:"column:weight"`
}
