package types

import "github.com/akhil-is-watching/techletics_alumni_reg/models"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Volunteer models.Volunteer `json:"volunteer"`
	Token     string           `json:"token"`
}
