package view

type (
	RegisterReq struct {
		Email    string
		Password string
		Username string
	}

	LoginReq struct {
		Email    string
		Password string
	}

	UpdatePasswordReq struct {
		OldPassword string
		NewPassword string
	}

	InfoReq struct{}

	UpdateUsernameReq struct {
		NewUsername string
	}
)
