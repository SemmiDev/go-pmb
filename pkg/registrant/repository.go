package registrant

import (
	"database/sql"
	"errors"
	sqlTrx "github.com/SemmiDev/go-pmb/pkg/stores/sql"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	db *sql.DB
}

func NewMySqlRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (s Repository) Save(r *Registrant) error {
	err := sqlTrx.WithTransaction(s.db, func(trx sqlTrx.Transaction) error {
		_, err := trx.Exec(`INSERT INTO registrants
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

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s Repository) FindByID(id string) (queryResult QueryResult) {
	rowsData := ReadResult{}
	err := s.db.QueryRow(`SELECT * FROM registrants WHERE registrant_id = ?`, id).Scan(
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

func (s Repository) FindByUsername(u string) (queryResult QueryResult) {
	rowsData := ReadResult{}
	err := s.db.QueryRow(`SELECT * FROM registrants WHERE username = ?`, u).Scan(
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

func (s Repository) FindByUsernameAndPassword(u, p string) (queryResult QueryResult) {
	rowsData := ReadResult{}
	err := s.db.QueryRow(`SELECT * FROM registrants WHERE username = ?`, u).Scan(
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

func (s Repository) UpdatePaymentStatus(id string, paymentStatus PaymentStatus) error {
	err := sqlTrx.WithTransaction(s.db, func(trx sqlTrx.Transaction) error {
		_, err := trx.Exec(`UPDATE registrants SET payment_status = ? WHERE registrant_id = ?`, paymentStatus, id)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
