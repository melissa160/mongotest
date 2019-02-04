package models

// User example
type UserExample struct {
	ID         uint   `json:"id"`
	Name       string `gorm:"not null" json:"name"`
	Email      string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password   string `gorm:"not null" json:"password"`
	Role       string `gorm:"not null;default:'user'; type:varchar(50) CHECk(role = 'user' or role = 'admin')" json:"role"`
	WebhookURL string `gorm:"not null" json:"webhook_url"`
	Token      string `gorm:"not null" json:"token"`
}
