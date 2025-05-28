package entity

type Faculty struct {
	Entity
	Code    string `gorm:"column:code"`
	Name    string `gorm:"column:name"`
	Dekan   string `gorm:"column:dekan"`
	Address string `gorm:"column:address"`
}
