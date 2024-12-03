package admin

import (
	"context"
	"errors"
	"fmt"

	"albatross.pe/hexagonal/internal/entity"
	"albatross.pe/hexagonal/internal/stores/db"
)

func (s personaService) ListPersona(ctx context.Context) error {
	return s.db.PersonaStore().ListPersona(ctx)
}

func (s personaService) SavePersona(ctx context.Context, persona *entity.Persona) error {

	txFn := func(txStore db.Store) error {
		persona.GenerateExtras()

		txStore.PersonaStore().ListPersona(ctx)

		err := txStore.PersonaStore().SavePersona(ctx, persona)
		if err != nil {
			return err
		}

		persona.Nombres = "Luke"
		persona.Paterno = "Green"
		err = txStore.PersonaStore().SavePersona(ctx, persona)
		if err != nil {
			return err
		}

		if true {
			return errors.New("error forzado")
		}

		fmt.Println("---------", persona.ID)

		return nil
	}

	return s.db.ExecuteTx(ctx, txFn)
}
