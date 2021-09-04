package domain

type Bill int64

const (
	BillS1D3D4 Bill = 152000
	BillS2     Bill = 252000
)

func (b Bill) Value() int64 {
	switch b {
	case BillS1D3D4:
		return int64(BillS1D3D4)
	case BillS2:
		return int64(BillS2)
	default:
		return 0
	}
}
