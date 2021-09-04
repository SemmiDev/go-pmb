package mysql

import (
	"database/sql"
	"github.com/SemmiDev/go-pmb/internal/registrant/command"
	"github.com/SemmiDev/go-pmb/internal/registrant/domain"
)

const (
	SaveQuery = `INSERT INTO registrants
				(registrant_id, name, email, phone, username, password, code, payment_url, program, bill, payment_status, created_date, last_updated)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	UpdateStatusQuery = `UPDATE registrants SET payment_status = ?`
)

type RegistrantCommandMysql struct {
	DB *sql.DB
}

func (s *RegistrantCommandMysql) Save(registrant *domain.Registrant) <-chan error {
	result := make(chan error)

	go func() {
		_, err := s.DB.Exec(SaveQuery,
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

func (s *RegistrantCommandMysql) UpdateStatus(id string, status domain.PaymentStatus) <-chan error {
	result := make(chan error)

	go func() {
		_, err := s.DB.Exec(UpdateStatusQuery, status)
		if err != nil {
			result <- err
		}

		result <- nil
		close(result)
	}()

	return result
}

func NewRegistrantCommandMysql(database *sql.DB) command.RegistrantCommand {
	return &RegistrantCommandMysql{
		DB: database,
	}
}
