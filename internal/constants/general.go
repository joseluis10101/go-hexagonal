package constants

type KeyValue struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Status struct {
	Name interface{} `json:"name"`
	Bg   string      `json:"bg"`
}

type GenStatus string

const (
	GEN_STATUS_ACTIVE   GenStatus = "activo"
	GEN_STATUS_INACTIVE GenStatus = "inactivo"
)

var (
	ALL_GEN_STATUS = []Status{
		{GEN_STATUS_ACTIVE, "bg-soft-grass"},
		{GEN_STATUS_INACTIVE, "bg-soft-grey"},
	}

	COLORS = []string{
		"blue",
		"indigo",
		"purple",
		"pink",
		"magenta",
		"red",
		"orange",
		"yellow",
		"green",
		"teal",
		"cyan",
		"acero",
		"deep",
		"sunday",
		"grass",
		"ruby",
		"marine",
		"city",
		"shadow"}

	MESES = []KeyValue{
		{"1", "Enero"},
		{"2", "Febrero"},
		{"3", "Marzo"},
		{"4", "Abril"},
		{"5", "Mayo"},
		{"6", "Junio"},
		{"7", "Julio"},
		{"8", "Agosto"},
		{"9", "Setiembre"},
		{"10", "Octubre"},
		{"11", "Noviembre"},
		{"12", "Diciembre"},
	}
)
