package forums

type TopicDTO struct {
	Author     string
	TopicTitle string
}
type Topics struct {
	Topics []TopicDTO
}
