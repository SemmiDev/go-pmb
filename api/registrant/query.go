package registrant

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

const (
	GetByIDQuery       = `SELECT * FROM registrants WHERE registrant_id = ?`
	GetByUsernameQuery = `SELECT * FROM registrants WHERE username = ?`
)

type QueryMySql struct {
	DB *sql.DB
}

func (s QueryMySql) GetByID(id string) <-chan QueryResult {
	result := make(chan QueryResult)

	go func() {
		rowsData := ReadResult{}

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
			result <- QueryResult{Error: err}
		}

		if err == sql.ErrNoRows {
			result <- QueryResult{Result: ReadResult{}}
		}

		result <- QueryResult{Result: rowsData}
		close(result)
	}()

	return result
}

func (s QueryMySql) GetByUsername(username string) <-chan QueryResult {
	result := make(chan QueryResult)

	go func() {
		rowsData := ReadResult{}

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
			result <- QueryResult{Error: err}
		}

		if err == sql.ErrNoRows {
			result <- QueryResult{Result: ReadResult{}}
		}

		result <- QueryResult{Result: rowsData}
		close(result)
	}()

	return result
}

func (s QueryMySql) GetByUsernameAndPassword(username, password string) <-chan QueryResult {
	result := make(chan QueryResult)

	go func() {
		rowsData := ReadResult{}

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
			result <- QueryResult{Error: err}
		}

		if err == sql.ErrNoRows {
			result <- QueryResult{Result: ReadResult{}}
		}

		if err := bcrypt.CompareHashAndPassword([]byte(rowsData.Password), []byte(password)); err != nil {
			result <- QueryResult{Error: err}
		}

		result <- QueryResult{Result: rowsData}
		close(result)
	}()

	return result
}

func NewMySqlQuery(database *sql.DB) QueryRepository {
	return QueryMySql{
		DB: database,
	}
}
