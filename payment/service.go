package payment

import (
	"github.com/SemmiDev/fiber-go-clean-arch/entity"
	"github.com/veritrans/go-midtrans"
)

type service struct {
	midClient midtrans.Client
}

type Service interface {
	GetPaymentURL(ts *Payment, register *entity.Registration) (string, error)
}

func NewService(client midtrans.Client) Service {
	return &service{midClient: client}
}

func (s *service) GetPaymentURL(ts *Payment, register *entity.Registration) (string, error) {
	snapGateway := midtrans.SnapGateway{
		Client: s.midClient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: register.Email,
			FName: register.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  ts.ID,
			GrossAmt: ts.Amount,
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
