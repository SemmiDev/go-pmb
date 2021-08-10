package models

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewLoginResponse(ID string, name string, username string, accessToken string, refreshToken string) *LoginResponse {
	return &LoginResponse{ID: ID, Name: name, Username: username, AccessToken: accessToken, RefreshToken: refreshToken}
}
