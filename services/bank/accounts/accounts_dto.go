package accounts

//FullAccountDTO type: Lo usaremos para generar los mapas que devolveremos con todos los datos de las cuentas
type FullAccountDTO struct {
	AccountNumber string  `json:"AccountNumber"`
	Owner         string  `json:"Owner"`
	Balance       int     `json:"Balance"`
	Bank          string  `json:"Bank"`
	created_at    []uint8 `json:"created_at"`
	/* updated_at time.Time `json:"updated_at"` */
}

type LogInUserDTO struct {
	Owner         string
	AccountNumber string
	Bank          string
}

type LoggedUserDTO struct {
	Owner string `json:"name"`
	Token string `json:"token"`
}

/* type SimpleAccountDTO struct {
	AccountNumber string  `json:"AccountNumber,string,omitempty"`
	Owner         string  `json:"Owner,string"`
	Balance       int     `json:"Balance,string"`
	Bank          string  `json:"Bank,string"`
	created_at    []uint8 `json:"created_at"`
	updated_at time.Time `json:"updated_at"`
} */
