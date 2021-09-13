package registrant

import (
	"errors"
	"github.com/SemmiDev/go-pmb/pkg/common/helper"
	"github.com/SemmiDev/go-pmb/pkg/payment"
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
	mid  payment.IMidtrans
}

func NewService(store *Repository, mid payment.IMidtrans) *Service {
	return &Service{repo: *store, mid: mid}
}

func (s *Service) Register(data *RegisterReq) (*RegisterResponse, error) {
	password := helper.GeneratePassword()
	result := NewRegistrant(data.Name, data.Email, data.Phone, password, data.Program)

	paymentPayload := &payment.MidtransPayload{
		Id:     uuid.NewString(),
		Name:   result.Name,
		Email:  result.Email,
		Amount: result.Bill.Val(),
	}

	paymentURL, err := s.mid.GetPaymentURL(paymentPayload)
	if err != nil {
		return nil, err
	}
	result.PaymentURL = paymentURL

	err = s.repo.Save(result)
	if err != nil {
		return nil, err
	}

	return ToRegisterRegistrantResp(result, password, paymentPayload.AmountFormatIDR()), nil
}

func (s *Service) UpdatePaymentStatus(data *UpdatePaymentStatusReq) error {
	var status PaymentStatus
	if data.PaymentType == "credit_card" && data.PaymentStatus == "capture" && data.FraudStatus == "accept" {
		status = PaymentStatusPaid
	} else if data.PaymentStatus == "settlement" {
		status = PaymentStatusPaid
	} else if data.PaymentStatus == "deny" || data.PaymentStatus == "expire" || data.PaymentStatus == "cancel" {
		status = PaymentStatusCancel
	} else {
		return errors.New("not recognize")
	}

	err := s.repo.UpdatePaymentStatus(data.RegisterID, status)
	if err != nil {
		return err
	}

	return nil
}
