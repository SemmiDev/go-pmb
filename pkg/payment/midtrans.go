package payment

import (
	"github.com/SemmiDev/go-pmb/pkg/common/config"
	mid "github.com/veritrans/go-midtrans"
)

type IMidtrans interface {
	GetPaymentURL(p *MidtransPayload) (string, error)
}

type midtrans struct {
	Client mid.Client
}

func NewMidtrans() *midtrans {
	midtrans := new(midtrans)

	midtrans.Client.ServerKey = config.MidtransServerKey
	midtrans.Client.ClientKey = config.MidtransClientKey
	midtrans.Client.APIEnvType = mid.Sandbox

	return midtrans
}

func (s *midtrans) GetPaymentURL(p *MidtransPayload) (string, error) {
	snapGateway := mid.SnapGateway{
		Client: s.Client,
	}

	snapReq := &mid.SnapReq{
		CustomerDetail: &mid.CustDetail{
			Email: p.Email,
			FName: p.Name,
		},
		TransactionDetails: mid.TransactionDetails{
			OrderID:  p.Id,
			GrossAmt: p.Amount,
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
