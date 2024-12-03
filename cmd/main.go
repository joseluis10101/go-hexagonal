package main

import (
	"context"
	"fmt"

	"albatross.pe/hexagonal/config"
	"albatross.pe/hexagonal/internal/core/admin"
	"albatross.pe/hexagonal/internal/entity"
	"albatross.pe/hexagonal/internal/stores/db"
)

func main() {
	ctx := context.Background()
	dua := &entity.Persona{Nombres: "Dua Leyla", Paterno: "EC"}

	dbx := config.NewMySQLConnection()

	dbs := db.NewStore(dbx)
	personaService := admin.NewPersonaService(dbs)

	err := personaService.ListPersona(ctx)
	if err != nil {
		fmt.Println(err)
	}

	err = personaService.SavePersona(ctx, dua)
	if err != nil {
		fmt.Println(err)
	}

}
