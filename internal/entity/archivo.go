package entity

type Archivo struct {
	ID            int64  `json:"id"`
	InstanciaID   int64  `json:"instanciaId"`
	TipoInstancia string `json:"tipoInstancia"`
	Titulo        string `json:"titulo"`
	Filename      string `json:"filename"`
	Orden         int    `json:"orden"`
	Url           string `json:"url"`
	Path          string `json:"path"`
	Mime          string `json:"mime"`
	Extension     string `json:"extension"`
	Size          int64  `json:"size"`
	Bucket        string `json:"bucket"`
}
