package midtrans

import "github.com/leekchan/accounting"

type Payment struct {
	ID     string
	Amount int64
}

func (t *Payment) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}
