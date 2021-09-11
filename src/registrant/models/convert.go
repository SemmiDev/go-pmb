package models

import "github.com/SemmiDev/go-pmb/src/registrant/entities"

// ToRegisterRegistrantResp for convert *entities.Registrant to RegisterResponse.
func ToRegisterRegistrantResp(result *entities.Registrant, p string, idr string) *RegisterResponse {
	return &RegisterResponse{
		ID:         result.ID,
		Email:      result.Email,
		Username:   result.Username,
		Password:   p,
		Bill:       idr,
		PaymentURL: result.PaymentURL,
	}
}
