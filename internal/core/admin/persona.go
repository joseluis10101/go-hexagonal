package admin

import (
	"context"

	"albatross.pe/hexagonal/internal/entity"
	"albatross.pe/hexagonal/internal/stores/db"
)

type PersonaService interface {
	ListPersona(ctx context.Context) error
	SavePersona(ctx context.Context, ficha *entity.Persona) error
}

type PersonaRepsository interface {
	PersonaStore() db.PersonaStore
	ExecuteTx(ctx context.Context, txFn db.TxFn) error
}

type personaService struct {
	db PersonaRepsository
}

func NewPersonaService(db PersonaRepsository) *personaService {
	return &personaService{
		db: db,
	}
}
