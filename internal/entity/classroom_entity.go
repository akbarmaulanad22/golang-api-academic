package entity

type Classroom struct {
	Model
	Name     string `gorm:"column:name"`
	Capacity int    `gorm:"column:capacity"`
	Location string `gorm:"column:location"`
}
