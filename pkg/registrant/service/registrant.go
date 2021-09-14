package service

import (
	"errors"
	"github.com/SemmiDev/go-pmb/pkg/common/helper"
	"github.com/SemmiDev/go-pmb/pkg/payment"
	"github.com/SemmiDev/go-pmb/pkg/registrant/entity"
	"github.com/SemmiDev/go-pmb/pkg/registrant/models"
	"github.com/SemmiDev/go-pmb/pkg/registrant/repository"
	"github.com/google/uuid"
)

type Service struct {
	Repo repository.Repository
	Mid  payment.IMidtrans
}

func NewService(store *repository.Repository, mid payment.IMidtrans) *Service {
	return &Service{Repo: *store, Mid: mid}
}

func (s *Service) Register(data *models.RegisterReq) (*models.RegisterResponse, error) {
	password := helper.GeneratePassword()
	result := entity.NewRegistrant(data.Name, data.Email, data.Phone, password, data.Program)

	paymentPayload := &payment.MidtransPayload{
		Id:     uuid.NewString(),
		Name:   result.Name,
		Email:  result.Email,
		Amount: result.Bill.Val(),
	}

	paymentURL, err := s.Mid.GetPaymentURL(paymentPayload)
	if err != nil {
		return nil, err
	}
	result.PaymentURL = paymentURL

	err = s.Repo.Save(result)
	if err != nil {
		return nil, err
	}

	return models.ToRegisterRegistrantResp(result, password, paymentPayload.AmountFormatIDR()), nil
}

func (s *Service) UpdatePaymentStatus(data *models.UpdatePaymentStatusReq) error {
	var status entity.PaymentStatus
	if data.PaymentType == "credit_card" && data.PaymentStatus == "capture" && data.FraudStatus == "accept" {
		status = entity.PaymentStatusPaid
	} else if data.PaymentStatus == "settlement" {
		status = entity.PaymentStatusPaid
	} else if data.PaymentStatus == "deny" || data.PaymentStatus == "expire" || data.PaymentStatus == "cancel" {
		status = entity.PaymentStatusCancel
	} else {
		return errors.New("not recognize")
	}

	err := s.Repo.UpdatePaymentStatus(data.RegisterID, status)
	if err != nil {
		return err
	}

	return nil
}
