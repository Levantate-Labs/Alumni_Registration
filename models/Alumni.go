package models

type Alumni struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Email       string `gorm:"unique" json:"email"`
	PhoneNo     string `json:"phone_no"`
	ImageURL    string `json:"image_url"`
	PassoutYear string `json:"passout_year"`
	CheckedIn   bool   `json:"checked_in"`
}
