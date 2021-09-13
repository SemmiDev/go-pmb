package payment

import "github.com/leekchan/accounting"

type FromRegistrant struct {
}

type MidtransPayload struct {
	Id     string
	Name   string
	Email  string
	Amount int64
}

func (t *MidtransPayload) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}
