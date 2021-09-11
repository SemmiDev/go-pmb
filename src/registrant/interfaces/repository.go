package interfaces

import (
	"github.com/SemmiDev/go-pmb/src/registrant/entities"
	"github.com/SemmiDev/go-pmb/src/registrant/models"
)

type IRepository interface {
	GetByID(id string) <-chan models.QueryResult
	GetByUsername(username string) <-chan models.QueryResult
	GetByUsernameAndPassword(username, password string) <-chan models.QueryResult

	Save(registrant *entities.Registrant) <-chan error
	UpdateStatus(id string, status entities.PaymentStatus) <-chan error
}
