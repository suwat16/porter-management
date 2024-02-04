package config

import "database/sql"

type Uow struct {
	db *sql.DB
	tx *sql.Tx
}

func NewUnitOfWork(db *sql.DB) *Uow {
	return &Uow{db: db}
}

func (u *Uow) DoInTransaction(f func() error) error {
	err := u.Begin()
	if err != nil {
		return err
	}
	err = f()
	if err != nil {
		u.Rollback()
		return err
	}
	return u.Commit()
}

func (u *Uow) Begin() error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	u.tx = tx
	return nil
}

func (u *Uow) Commit() error {
	return u.tx.Commit()
}

func (u *Uow) Rollback() error {
	return u.tx.Rollback()
}
