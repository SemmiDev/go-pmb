package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
)

// TXHandler is handler for working with transaction.
// This is wrapper function for commit and rollback.
func TXHandler(db *sqlx.DB, f func(*sqlx.Tx) error) (err error) {
	tx, err := db.Beginx()
	if err != nil {
		return errors.Wrap(err, "start transaction failed")
	}

	defer func() {
		if p := recover(); p != nil || err != nil {
			rollBackErr := tx.Rollback()
			if rollBackErr != nil {
				log.Fatalf("rollback failed: %v", rollBackErr)
			}
			if p != nil {
				err = errors.New(fmt.Sprintf("recovered: %v", p))
			} else {
				err = errors.Wrap(err, "transaction: operation failed")
			}
		}
	}()

	err = f(tx)
	return err
}
