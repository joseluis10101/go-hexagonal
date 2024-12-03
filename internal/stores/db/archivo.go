package db

import (
	"context"
	"fmt"

	"albatross.pe/hexagonal/internal/entity"
	"github.com/jmoiron/sqlx"
)

type ArchivoStore interface {
	SaveArchivo(ctx context.Context, p *entity.Archivo) error
}

type archivoStore struct {
	db sqlx.ExtContext
}

func NewArchivoStore(db sqlx.ExtContext) archivoStore {
	return archivoStore{
		db: db,
	}
}

func (s archivoStore) SaveArchivo(ctx context.Context, p *entity.Archivo) error {
	fmt.Println("saveArchivo store method")
	return nil
}
