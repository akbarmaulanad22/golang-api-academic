package entity

type User struct {
	Entity
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Token    string `gorm:"column:token"`
	RoleID   uint   `gorm:"column:role_id"`

	// relationship
	Role     *Role     `gorm:"foreignKey:role_id;references:id"`
	Lecturer *Lecturer `gorm:"foreignkey:UserID;references:ID"`
	Student  Student   `gorm:"foreignkey:user_id;references:id"`
}
