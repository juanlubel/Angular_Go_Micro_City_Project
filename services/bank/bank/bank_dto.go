package bank

//BankDTO type: Lo usaremos para generar los mapas que devolveremos
type BankDTO struct {
	ID         uint    `json:"id,string,omitempty"`
	BankName   string  `json:"BankName"`
	created_at []uint8 `json:"created_at"`
	/* updated_at time.Time `json:"updated_at"` */
}
