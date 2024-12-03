package db

import (
	"context"
	"fmt"

	"albatross.pe/hexagonal/internal/entity"
	"github.com/jmoiron/sqlx"
)

type PersonaStore interface {
	SavePersona(ctx context.Context, p *entity.Persona) error
	ListPersona(ctx context.Context) error
}

type personaStore struct {
	db sqlx.ExtContext
}

func NewPersonaStore(db sqlx.ExtContext) personaStore {
	return personaStore{
		db: db,
	}
}

func (s personaStore) ListPersona(ctx context.Context) error {

	rows, err := s.db.QueryxContext(ctx, "SELECT nombres FROM gen_persona")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var persona entity.Persona
		if err := rows.StructScan(&persona); err != nil {
			return err
		}

		fmt.Println(persona.Nombres)
	}

	return err
}

func (s personaStore) SavePersona(ctx context.Context, p *entity.Persona) error {

	re, err := s.db.ExecContext(ctx, "INSERT INTO gen_persona (nombres, paterno) VALUES (?, ?)", p.Nombres, p.Paterno)
	if err != nil {
		return err
	}
	id, err := re.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = id
	return err
}
