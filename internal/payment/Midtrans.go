package payment

import (
	"github.com/SemmiDev/go-pmb/internal/common/config"
	"github.com/SemmiDev/go-pmb/internal/registrant/domain"
	"github.com/leekchan/accounting"
	"github.com/veritrans/go-midtrans"
)

type Payload struct {
	ID     string
	Amount int64
}

func (t *Payload) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}

type IMidtrans interface {
	GetPaymentURL(p *Payload, register *domain.Registrant) (string, error)
}

type Midtrans struct {
	Client midtrans.Client
}

func NewMidtrans() IMidtrans {
	mid := Midtrans{}
	mid.Client.ServerKey = config.MidtransServerKey
	mid.Client.ClientKey = config.MidtransClientKey
	mid.Client.APIEnvType = midtrans.Sandbox
	return &mid
}

func (s *Midtrans) GetPaymentURL(p *Payload, register *domain.Registrant) (string, error) {
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
