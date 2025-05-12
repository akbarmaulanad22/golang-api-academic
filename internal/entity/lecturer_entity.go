package entity

import "time"

type Lecturer struct {
	Nidn       int       `gorm:"column:nidn;primaryKey"`
	Name       string    `gorm:"column:name"`
	Gender     string    `gorm:"column:gender"`
	Degree     string    `gorm:"column:degree"`
	IsFullTime bool      `gorm:"column:is_full_time"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
