package domains

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Age       uint   `json:"age"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
