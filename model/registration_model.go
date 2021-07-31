package model

type RegistrationRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Program string `json:"program"`
}

type RegistrationResponse struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Bill           Bill   `json:"bill"`
	VirtualAccount string `json:"virtual_account"`
}

type UpdateStatus struct {
	VirtualAccount string `json:"virtual_account"`
}
