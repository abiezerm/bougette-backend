package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	FirstName *string `gorm:"type:varchar(200)" json:"first_name"`
	LastName  *string `gorm:"type:varchar(200)" json:"last_name"`
	Gender    *string `gorm:"type:varchar(50)" json:"gender"`
	Password  string  `gorm:"type:varchar(200); not null " json:"password"`
	Email     string  `gorm:"type:varchar(200); not null; unique " json:"email"`
}

func (receiver UserModel) TableName() string {
	return "users"
}
