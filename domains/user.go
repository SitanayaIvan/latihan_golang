package domains

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       uint   `json:"age"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
