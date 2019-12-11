package topics

//TopicDTO type: Lo usaremos para generar los mapas que devolveremos con todos los datos
type TopicDTO struct {
	Author     string `json:"Author"`
	TopicTitle string `json:"TopicTitle"`
}

type CommentDTO struct {
	Author      string `json:"Author"`
	Body        string `json:"Body"`
	TopicTittle string `json:TopicTitle`
}

type TopicWithCommentsDTO struct {
	Topic    TopicDTO     `json:"Topic"`
	Comments []CommentDTO `json:"Comments"`
}
