package mysql

import (
	"database/sql"
	"github.com/SemmiDev/go-pmb/pkg/registrant/domain"
	"github.com/SemmiDev/go-pmb/pkg/registrant/repository"
)

type RegistrantRepositoryMysql struct {
	DB *sql.DB
}

func (s *RegistrantRepositoryMysql) Save(registrant *domain.Registrant) <-chan error {
	result := make(chan error)

	go func() {
		_, err := s.DB.Exec(`INSERT INTO registrants
				(registrant_id, name, email, phone, username, password, code, payment_url, program, bill, payment_status, created_date, last_updated)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			registrant.RegistrantId,
			registrant.Name,
			registrant.Email,
			registrant.Phone,
			registrant.Username,
			registrant.Password,
			registrant.Code,
			registrant.PaymentURL,
			registrant.Program,
			registrant.Bill,
			registrant.PaymentStatus,
			registrant.CreatedDate,
			registrant.LastUpdated,
		)

		if err != nil {
			result <- err
		}

		result <- nil
		close(result)
	}()

	return result
}

func (s *RegistrantRepositoryMysql) UpdateStatus(id string, status domain.PaymentStatus) <-chan error {
	result := make(chan error)

	go func() {
		_, err := s.DB.Exec(`UPDATE registrants SET payment_status = ?`, status)
		if err != nil {
			result <- err
		}

		result <- nil
		close(result)
	}()

	return result
}

func NewRegistrantRepositoryMysql(database *sql.DB) repository.RegistrantRepository {
	return &RegistrantRepositoryMysql{
		DB: database,
	}
}
