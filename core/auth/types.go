package auth

type SignInPayload struct {
	IdToken string `json:"idToken" binding:"required"`
}