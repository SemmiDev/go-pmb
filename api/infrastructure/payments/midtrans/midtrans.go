package midtrans

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/entities"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/environments"
	"github.com/veritrans/go-midtrans"
)

type IMidtrans interface {
	GetPaymentURL(p *Payment, register *entities.Registration) (string, error)
}

type Midtrans struct {
	Client midtrans.Client
}

func New() IMidtrans {
	mid := Midtrans{}
	mid.Client.ServerKey = environments.MidtransServerKey
	mid.Client.ClientKey = environments.MidtransClientKey
	mid.Client.APIEnvType = midtrans.Sandbox
	return &mid
}

func (s *Midtrans) GetPaymentURL(p *Payment, register *entities.Registration) (string, error) {
	snapGateway := midtrans.SnapGateway{
		Client: s.Client,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: register.Email,
			FName: register.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  p.ID,
			GrossAmt: p.Amount,
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
