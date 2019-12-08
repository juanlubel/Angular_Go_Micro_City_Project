package topics

//TopicDTO type: Lo usaremos para generar los mapas que devolveremos con todos los datos
type TopicDTO struct {
	Author     string `json:"Author"`
	TopicTitle string `json:"TopicTitle"`
}
