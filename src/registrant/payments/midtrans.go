package payments

import (
	"github.com/SemmiDev/go-pmb/src/common/config"
	"github.com/SemmiDev/go-pmb/src/registrant/entities"
	"github.com/leekchan/accounting"
	mid "github.com/veritrans/go-midtrans"
)

type Payload struct {
	id     string
	amount int64
}

func NewPayload(id string, amount int64) *Payload {
	return &Payload{id: id, amount: amount}
}

func (t *Payload) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.amount)
}

type IMidtrans interface {
	GetPaymentURL(p *Payload, register *entities.Registrant) (string, error)
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

func (s *midtrans) GetPaymentURL(p *Payload, register *entities.Registrant) (string, error) {
	snapGateway := mid.SnapGateway{
		Client: s.Client,
	}

	snapReq := &mid.SnapReq{
		CustomerDetail: &mid.CustDetail{
			Email: register.Email,
			FName: register.Name,
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
