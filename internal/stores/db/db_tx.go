package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type txStore struct {
	tx           *sqlx.Tx
	personaStore personaStore
	archivoStore archivoStore
}

func NewTxStore(tx *sqlx.Tx) Store {

	return &txStore{
		tx:           tx,
		personaStore: NewPersonaStore(tx),
		archivoStore: NewArchivoStore(tx),
	}
}

func (s txStore) PersonaStore() PersonaStore {
	return s.personaStore
}

func (s txStore) ArchivoStore() ArchivoStore {
	return s.archivoStore
}

func (dbs txStore) ExecuteTx(ctx context.Context, txFn TxFn) error {
	return fmt.Errorf("nested transactions not supported")

}
