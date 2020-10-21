package controller

// User User
type User struct {
	ID    uint   `gorm:"column:id;primary_key" json:"id"`
	Email string `gorm:"column:email" json:"email"`
	Name  string `gorm:"column:name" json:"name"`
}
