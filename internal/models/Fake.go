package models

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Fake struct {
	String string    `db:"string"`
	Date   time.Time `db:"date"`
}

func (f *Fake) Create(db *sqlx.DB) error {
	f.Date = time.Now()
	_, err := db.NamedExec("INSERT INTO fakes (string, created_at) VALUES (:string, :date)", f)

	if err != nil {
		return err
	}

	return err
}
