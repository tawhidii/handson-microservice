package models

type User struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Email   string `json:"email" gorm:"unique"`
	Address string `json:"address"`
	Birth   string `json:"birth"`
	Gender  string `json:"gender"`
	Bio     string `json:"bio"`
	About   string `json:"about"`
}
