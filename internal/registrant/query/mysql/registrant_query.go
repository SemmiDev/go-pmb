package mysql

import (
	"database/sql"
	"github.com/SemmiDev/go-pmb/internal/registrant/query"
	"github.com/SemmiDev/go-pmb/internal/registrant/storage"
	"golang.org/x/crypto/bcrypt"
)

const (
	GetByIDQuery       = `SELECT * FROM registrants WHERE registrant_id = ?`
	GetByUsernameQuery = `SELECT * FROM registrants WHERE username = ?`
)

type RegistrantQueryMySql struct {
	DB *sql.DB
}

func (s RegistrantQueryMySql) GetByID(id string) <-chan query.RegistrantQueryResult {
	result := make(chan query.RegistrantQueryResult)

	go func() {
		rowsData := storage.RegistrantResult{}

		err := s.DB.QueryRow(GetByIDQuery, id).Scan(
			&rowsData.RegistrantId,
			&rowsData.Name,
			&rowsData.Email,
			&rowsData.Phone,
			&rowsData.Username,
			&rowsData.Password,
			&rowsData.Code,
			&rowsData.PaymentURL,
			&rowsData.Program,
			&rowsData.Bill,
			&rowsData.PaymentStatus,
			&rowsData.CreatedDate,
			&rowsData.LastUpdated,
		)

		if err != nil && err != sql.ErrNoRows {
			result <- query.RegistrantQueryResult{Error: err}
		}

		if err == sql.ErrNoRows {
			result <- query.RegistrantQueryResult{Result: storage.RegistrantResult{}}
		}

		result <- query.RegistrantQueryResult{Result: rowsData}
		close(result)
	}()

	return result
}

func (s RegistrantQueryMySql) GetByUsername(username string) <-chan query.RegistrantQueryResult {
	result := make(chan query.RegistrantQueryResult)

	go func() {
		rowsData := storage.RegistrantResult{}

		err := s.DB.QueryRow(GetByUsernameQuery, username).Scan(
			&rowsData.RegistrantId,
			&rowsData.Name,
			&rowsData.Email,
			&rowsData.Phone,
			&rowsData.Username,
			&rowsData.Password,
			&rowsData.Code,
			&rowsData.PaymentURL,
			&rowsData.Program,
			&rowsData.Bill,
			&rowsData.PaymentStatus,
			&rowsData.CreatedDate,
			&rowsData.LastUpdated,
		)

		if err != nil && err != sql.ErrNoRows {
			result <- query.RegistrantQueryResult{Error: err}
		}

		if err == sql.ErrNoRows {
			result <- query.RegistrantQueryResult{Result: storage.RegistrantResult{}}
		}

		result <- query.RegistrantQueryResult{Result: rowsData}
		close(result)
	}()

	return result
}

func (s RegistrantQueryMySql) GetByUsernameAndPassword(username, password string) <-chan query.RegistrantQueryResult {
	result := make(chan query.RegistrantQueryResult)

	go func() {
		rowsData := storage.RegistrantResult{}

		err := s.DB.QueryRow(GetByUsernameQuery, username).Scan(
			&rowsData.RegistrantId,
			&rowsData.Name,
			&rowsData.Email,
			&rowsData.Phone,
			&rowsData.Username,
			&rowsData.Password,
			&rowsData.Code,
			&rowsData.PaymentURL,
			&rowsData.Program,
			&rowsData.Bill,
			&rowsData.PaymentStatus,
			&rowsData.CreatedDate,
			&rowsData.LastUpdated,
		)

		if err != nil && err != sql.ErrNoRows {
			result <- query.RegistrantQueryResult{Error: err}
		}

		if err == sql.ErrNoRows {
			result <- query.RegistrantQueryResult{Result: storage.RegistrantResult{}}
		}

		if err := bcrypt.CompareHashAndPassword([]byte(rowsData.Password), []byte(password)); err != nil {
			result <- query.RegistrantQueryResult{Error: err}
		}

		result <- query.RegistrantQueryResult{Result: rowsData}
		close(result)
	}()

	return result
}

func NewRegistrantMySqlQuery(database *sql.DB) query.RegistrantQuery {
	return RegistrantQueryMySql{
		DB: database,
	}
}
