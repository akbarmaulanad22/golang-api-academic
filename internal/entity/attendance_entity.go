package entity

import "time"

type Attendance struct {
	Model
	Status string    `gorm:"column:status"`
	Time   time.Time `gorm:"column:date"`

	// foreign key
	ScheduleId int `gorm:"column:schedule_id"`
	UserId     int `gorm:"column:user_id"`

	// relationship
	Schedule Schedule `gorm:"foreignKey:schedule_id;references:id"`
	User     User     `gorm:"foreignKey:user_id;references:id"`
}
