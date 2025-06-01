package entity

import "time"

type Attendance struct {
	Entity
	Status string    `gorm:"column:status"`
	Time   time.Time `gorm:"column:time"`

	// foreign key
	ScheduleId uint `gorm:"column:schedule_id"`
	UserId     uint `gorm:"column:user_id"`

	// relationship
	Schedule Schedule `gorm:"foreignKey:schedule_id;references:id;-create"`
	User     User     `gorm:"foreignKey:user_id;references:id;-create"`
}
