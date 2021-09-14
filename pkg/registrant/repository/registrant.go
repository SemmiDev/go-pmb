package repository

import (
	"database/sql"
	"errors"
	"github.com/SemmiDev/go-pmb/pkg/common/database"
	"github.com/SemmiDev/go-pmb/pkg/registrant/entity"
	"github.com/SemmiDev/go-pmb/pkg/registrant/storage"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	db *sqlx.DB
}

func NewMySqlRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (s Repository) Save(r *entity.Registrant) error {
	if err := database.TXHandler(s.db, func(tx *sqlx.Tx) error {
		_, err := tx.Exec(`INSERT INTO registrants
		(registrant_id,name,email,phone,username,password,code,payment_url,program,bill,payment_status,created_date,last_updated)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			r.ID,
			r.Name,
			r.Email,
			r.Phone,
			r.Username,
			r.Password,
			r.Code,
			r.PaymentURL,
			r.Program,
			r.Bill,
			r.PaymentStatus,
			r.CreatedDate,
			r.LastUpdated,
		)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s Repository) FindByID(id string) (queryResult storage.QueryResult) {
	rowsData := storage.ReadResult{}
	err := s.db.QueryRowx(`SELECT * FROM registrants WHERE registrant_id = ?`, id).Scan(
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
		queryResult.Error = err
	}

	if err == sql.ErrNoRows {
		queryResult.Error = errors.New("registrant not found")
	}

	queryResult.ReadResult = &rowsData
	return
}

func (s Repository) FindByUsername(u string) (queryResult storage.QueryResult) {
	rowsData := storage.ReadResult{}
	err := s.db.QueryRowx(`SELECT * FROM registrants WHERE username = ?`, u).Scan(
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
		queryResult.Error = err
	}

	if err == sql.ErrNoRows {
		queryResult.Error = errors.New("registrant not found")
	}

	queryResult.ReadResult = &rowsData
	return
}

func (s Repository) FindByUsernameAndPassword(u, p string) (queryResult storage.QueryResult) {
	rowsData := storage.ReadResult{}
	err := s.db.QueryRowx(`SELECT * FROM registrants WHERE username = ?`, u).Scan(
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
		queryResult.Error = err
	}

	if err == sql.ErrNoRows {
		queryResult.Error = errors.New("registrant not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(rowsData.Password), []byte(p)); err != nil {
		queryResult.Error = err
	}

	queryResult.ReadResult = &rowsData
	return
}

func (s Repository) UpdatePaymentStatus(id string, paymentStatus entity.PaymentStatus) error {
	if err := database.TXHandler(s.db, func(tx *sqlx.Tx) error {
		_, err := tx.Exec(`UPDATE registrants SET payment_status = ? WHERE registrant_id = ?`, paymentStatus, id)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
