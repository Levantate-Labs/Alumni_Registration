package types

type AlumniCreateInput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNo     string `json:"phone_no"`
	ImageURL    string `json:"image_url"`
	PassoutYear string `json:"passout_year"`
}
