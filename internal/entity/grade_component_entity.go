package entity

type GradeComponent struct {
	Model
	Name   string `gorm:"column:name"`
	Weight int    `gorm:"column:weight"`
}
