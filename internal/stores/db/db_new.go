package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TxFn func(store Store) error

type Store interface {
	PersonaStore() PersonaStore
	ArchivoStore() ArchivoStore
	ExecuteTx(ctx context.Context, txFn TxFn) error
}

type store struct {
	db           *sqlx.DB
	personaStore personaStore
	archivoStore archivoStore
}

func NewStore(db *sqlx.DB) Store {

	return &store{
		db:           db,
		personaStore: NewPersonaStore(db),
		archivoStore: NewArchivoStore(db),
	}
}

func (s store) PersonaStore() PersonaStore {
	return s.personaStore
}

func (s store) ArchivoStore() ArchivoStore {
	return s.archivoStore
}

func (dbs store) ExecuteTx(ctx context.Context, txFn TxFn) error {

	tx, err := dbs.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	txStore := NewTxStore(tx)

	err = txFn(txStore)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return err
}
