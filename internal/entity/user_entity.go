package entity

type User struct {
	Entity
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Token    string `gorm:"column:token"`
	RoleID   uint   `gorm:"column:role_id"`

	// relationship
	Role *Role `gorm:"foreignKey:role_id;references:id"`
}
