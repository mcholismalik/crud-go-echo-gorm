package auth

import "crud-go-echo-gorm/pkg/util/response"

type RegisterDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseDTO struct {
	Token string `json:"token"`
}

// for response login
type LoginResponseDOC struct {
	Body struct {
		Meta response.MetaResponse
		Data LoginResponseDTO
	}
}
