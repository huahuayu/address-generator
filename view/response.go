package view

import "time"

type (
	LoginRes struct {
		Token string `json:"token"`
	}

	InfoRes struct {
		Id       string    `json:"id"`
		Email    string    `json:"email"`
		Username string    `json:"username"`
		CreateAt time.Time `json:"createdAt"`
	}
)
