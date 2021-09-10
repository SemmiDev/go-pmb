package registrant

import (
	"github.com/SemmiDev/go-pmb/common/config"
	"github.com/leekchan/accounting"
	mid "github.com/veritrans/go-midtrans"
)

type payload struct {
	id     string
	amount int64
}

func (t *payload) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.amount)
}

type IMidtrans interface {
	GetPaymentURL(p *payload, register *Registrant) (string, error)
}

type midtrans struct {
	Client mid.Client
}

func NewMidtrans() IMidtrans {
	midtrans := new(midtrans)
	midtrans.Client.ServerKey = config.MidtransServerKey
	midtrans.Client.ClientKey = config.MidtransClientKey
	midtrans.Client.APIEnvType = mid.Sandbox

	return midtrans
}

func (s *midtrans) GetPaymentURL(p *payload, register *Registrant) (string, error) {
	snapGateway := mid.SnapGateway{
		Client: s.Client,
	}

	snapReq := &mid.SnapReq{
		CustomerDetail: &mid.CustDetail{
			Email: register.Email(),
			FName: register.Name(),
		},
		TransactionDetails: mid.TransactionDetails{
			OrderID:  p.id,
			GrossAmt: p.amount,
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
