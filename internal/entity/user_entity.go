package entity

type User struct {
	Model
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	RoleID   int    `gorm:"column:role_id"`

	// relationship belongs to (one to one)
	// Student  Student  `gorm:"foreignKey:id;references:user_id"`
	// Lecturer Lecturer `gorm:"foreignKey:id;references:user_id"`
	Role *Role `gorm:"foreignKey:role_id;references:id"`
}
