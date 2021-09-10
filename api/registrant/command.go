package registrant

import (
	"database/sql"
	"log"
)

const (
	SaveQuery = `INSERT INTO registrants
				(registrant_id,name,email,phone,username,password,code,payment_url,program,bill,payment_status,created_date,last_updated)
				 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	UpdateStatusQuery = `UPDATE registrants SET payment_status = ?`
)

type CommandMysql struct {
	DB *sql.DB
}

func (s *CommandMysql) Save(registrant *Registrant) <-chan error {
	result := make(chan error)

	go func() {
		trx, err := s.DB.Begin()
		if err != nil {
			result <- err
		}

		log.Println(*registrant)

		_, err = trx.Exec(SaveQuery,
			registrant.Id(),
			registrant.Name(),
			registrant.Email(),
			registrant.Phone(),
			registrant.Username(),
			registrant.Password(),
			registrant.Code(),
			registrant.PaymentURL(),
			registrant.Program(),
			registrant.Bill(),
			registrant.PaymentStatus(),
			registrant.CreatedDate(),
			registrant.LastUpdated(),
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

func (s *CommandMysql) UpdateStatus(id string, status PaymentStatus) <-chan error {
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

func NewCommandMysql(database *sql.DB) CmdRepository {
	return &CommandMysql{
		DB: database,
	}
}
