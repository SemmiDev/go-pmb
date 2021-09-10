package registrant

func ToRegisterRegistrantResp(result *Registrant, p string, idr string) *RegisterResponse {
	return &RegisterResponse{
		ID:         string(result.id),
		Email:      result.Email(),
		Username:   result.Username(),
		Password:   p,
		Bill:       idr,
		PaymentURL: result.PaymentURL(),
	}
}
