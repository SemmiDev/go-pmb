package repositories

import (
	"database/sql"
	"github.com/SemmiDev/go-pmb/src/registrant/entities"
	"github.com/SemmiDev/go-pmb/src/registrant/interfaces"
	"github.com/SemmiDev/go-pmb/src/registrant/models"
	"golang.org/x/crypto/bcrypt"
)

const (
	SaveQuery = `INSERT INTO registrants
				(registrant_id,name,email,phone,username,password,code,payment_url,program,bill,payment_status,created_date,last_updated)
				 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	UpdateStatusQuery  = `UPDATE registrants SET payment_status = ?`
	GetByIDQuery       = `SELECT * FROM registrants WHERE registrant_id = ?`
	GetByUsernameQuery = `SELECT * FROM registrants WHERE username = ?`
)

type registrantRepository struct {
	DB *sql.DB
}

func (s registrantRepository) GetByID(id string) <-chan models.QueryResult {
	result := make(chan models.QueryResult)

	go func() {
		rowsData := models.ReadResult{}

		err := s.DB.QueryRow(GetByIDQuery, id).Scan(
			&rowsData.ID,
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
			result <- models.QueryResult{Error: err}
		}

		if err == sql.ErrNoRows {
			result <- models.QueryResult{Result: models.ReadResult{}}
		}

		result <- models.QueryResult{Result: rowsData}
		close(result)
	}()

	return result
}

func (s registrantRepository) GetByUsername(username string) <-chan models.QueryResult {
	result := make(chan models.QueryResult)

	go func() {
		rowsData := models.ReadResult{}

		err := s.DB.QueryRow(GetByUsernameQuery, username).Scan(
			&rowsData.ID,
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
			result <- models.QueryResult{Error: err}
		}

		if err == sql.ErrNoRows {
			result <- models.QueryResult{Result: models.ReadResult{}}
		}

		result <- models.QueryResult{Result: rowsData}
		close(result)
	}()

	return result
}

func (s registrantRepository) GetByUsernameAndPassword(username, password string) <-chan models.QueryResult {
	result := make(chan models.QueryResult)

	go func() {
		rowsData := models.ReadResult{}

		err := s.DB.QueryRow(GetByUsernameQuery, username).Scan(
			&rowsData.ID,
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
			result <- models.QueryResult{Error: err}
		}

		if err == sql.ErrNoRows {
			result <- models.QueryResult{Result: models.ReadResult{}}
		}

		if err := bcrypt.CompareHashAndPassword([]byte(rowsData.Password), []byte(password)); err != nil {
			result <- models.QueryResult{Error: err}
		}

		result <- models.QueryResult{Result: rowsData}
		close(result)
	}()

	return result
}

func (s *registrantRepository) Save(registrant *entities.Registrant) <-chan error {
	result := make(chan error)

	go func() {
		trx, err := s.DB.Begin()
		if err != nil {
			result <- err
		}

		_, err = trx.Exec(SaveQuery,
			registrant.ID,
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

		err = trx.Commit()
		if err != nil {
			result <- nil
		}

		result <- nil
		close(result)
	}()

	return result
}

func (s *registrantRepository) UpdateStatus(id string, status entities.PaymentStatus) <-chan error {
	result := make(chan error)

	go func() {
		trx, err := s.DB.Begin()
		if err != nil {
			result <- err
		}

		_, err = trx.Exec(UpdateStatusQuery, status)
		if err != nil {
			result <- err
		}

		err = trx.Commit()
		if err != nil {
			result <- nil
		}

		result <- nil
		close(result)
	}()

	return result
}

func NewRegistrantRepository(DB *sql.DB) interfaces.IRepository {
	return &registrantRepository{DB: DB}
}
