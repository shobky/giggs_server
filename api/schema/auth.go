package schema

type LoginBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

type SignupBody struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password"`
	GivenName  string `json:"given_name" validate:"required,min=4"`
	FamilyName string `json:"family_name" validate:"required,min=4"`
	Name       string `json:"name" validate:"required,min=8"`
}

type TokenPayload struct {
	ID    uint
	Email string
}

type CallbackToken struct {
	Token string `json:"token"`
}

type GoogleResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	GivenName  string `json:"givin_name"`
	FamilyName string `json:"family_name"`
	Name       string `json:"name"`
	Verified   bool   `json:"verified_email"`
	Picture    string `json:"picture"`
	Provider   string `json:"provider"`
}
