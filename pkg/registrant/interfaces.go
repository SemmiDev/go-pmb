package registrant

type Saver interface {
	Save(r *Registrant) error
}

type Finder interface {
	FindByID(id string) QueryResult
	FindByUsername(u string) QueryResult
	FindByUsernameAndPassword(u, p string) QueryResult
}

type Updater interface {
	UpdatePaymentStatus(id string, paymentStatus PaymentStatus) error
}
