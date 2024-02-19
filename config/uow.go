package config

import "database/sql"

type Uow struct {
	Db *sql.DB
	Tx *sql.Tx
}

func NewUnitOfWork(db *sql.DB) *Uow {
	return &Uow{Db: db}
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
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}
	u.Tx = tx
	return nil
}

func (u *Uow) Commit() error {
	return u.Tx.Commit()
}

func (u *Uow) Rollback() error {
	return u.Tx.Rollback()
}
