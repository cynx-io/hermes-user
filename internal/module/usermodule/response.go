package usermodule

import (
	"hermes/internal/model/entity"
	"time"
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type RegisterUserResponse struct {
	Id int `json:"id"`
}

type LoginUserResponse struct {
	AuthResponse
	UserResponse
}

type UserResponse struct {
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	Username    string    `json:"username"`
	ID          int       `json:"id"`
	Coin        int       `json:"coin"`
}

func NewUserResponse(user *entity.TblUser) *UserResponse {

	if user == nil {
		return nil
	}

	return &UserResponse{
		ID:          user.Id,
		Username:    user.Username,
		Coin:        user.Coin,
		CreatedDate: user.CreatedDate,
		UpdatedDate: user.UpdatedDate,
	}
}
