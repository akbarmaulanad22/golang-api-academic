package entity

type Classroom struct {
	Entity
	Name     string `gorm:"column:name"`
	Capacity int    `gorm:"column:capacity"`
	Location string `gorm:"column:location"`
}
