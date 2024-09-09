package response

import "time"

type TokenResponse struct {
	Token string `json:"token"`
	Exp   time.Time  `json:"expires_at"`
}

type SignUpResponse struct {
	Message string   `json:"message"`
	Token   TokenResponse `json:"token"`
}

func NewSignUpResponse(token string, exp *time.Time) (*SignUpResponse, error) {
	return &SignUpResponse{
		Message: "User created successfully",
		Token: TokenResponse{
			Token: token,
			Exp:   *exp,
		},
	}, nil
}

type SignInResponse struct {
	Message string   `json:"message"`
	Token   TokenResponse `json:"token"`
}

func NewSignInResponse(token string, exp *time.Time) (*SignInResponse, error) {
	return &SignInResponse{
		Message: "Signed in successfully",
		Token: TokenResponse{
			Token: token,
			Exp:   *exp,
		},
	}, nil
}
