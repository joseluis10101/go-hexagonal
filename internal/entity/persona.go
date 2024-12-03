package entity

import (
	"fmt"

	"albatross.pe/hexagonal/internal/constants"
	"albatross.pe/hexagonal/pkg"
)

type Persona struct {
	ID            int64               `json:"id"`
	Paterno       string              `json:"paterno"`
	Materno       string              `json:"materno"`
	Nombres       string              `json:"nombres"`
	Sexo          bool                `json:"sexo"`
	EmailCompania string              `json:"emailCompania"`
	NumeroDni     string              `json:"numeroDni"`
	Avatar        string              `json:"avatar"`
	Color         string              `json:"color"`
	Celular       string              `json:"celular"`
	Estado        constants.GenStatus `json:"estado"`
}

func (p *Persona) GenerateExtras() {

	fullname := fmt.Sprintf("%s %s", p.Nombres, p.Paterno)

	avatar := pkg.FindAbbreviation(fullname)
	p.Avatar = avatar

	if p.ID == 0 {
		color := pkg.FindRandomColor()
		p.Color = color
		p.Estado = constants.GEN_STATUS_ACTIVE
	}
}
