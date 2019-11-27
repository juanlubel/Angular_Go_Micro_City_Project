package admin

//AdminDTO type: Lo usaremos para generar los mapas que devolveremos
type AdminDTO struct {
	ID    uint   `json:"id,string,omitempty"`
	Name  string `json:"name"`
	Pass  string   `json:"pass"`
	Token string `json:"token"`
}

type LogInAdminDTO struct {
	Name  string `json:"name"`
	Pass  string   `json:"pass"`
}

type LoggedDTO struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}
