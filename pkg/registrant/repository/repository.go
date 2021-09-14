package repository

import (
	"github.com/SemmiDev/go-pmb/pkg/registrant/entity"
	"github.com/SemmiDev/go-pmb/pkg/registrant/storage"
)

type Saver interface {
	Save(r *entity.Registrant) error
}

type Finder interface {
	FindByID(id string) storage.QueryResult
	FindByUsername(u string) storage.QueryResult
	FindByUsernameAndPassword(u, p string) storage.QueryResult
}

type Updater interface {
	UpdatePaymentStatus(id string, paymentStatus entity.PaymentStatus) error
}
