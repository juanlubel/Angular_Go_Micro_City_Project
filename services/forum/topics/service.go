package topics

//TopicService : Is in charge of communicate the api with the repository
type TopicService struct {
	TopicRepository TopicRepository
}

func ProvideTopicService(p TopicRepository) TopicService {
	return TopicService{TopicRepository: p}
}

func (p *TopicService) FindAll() []Topic {
	return p.TopicRepository.FindAll()
}

func (p *TopicService) FindByOwner(owner string) Topic {
	return p.TopicRepository.FindByOwner(owner)
}
func (p *TopicService) FindHisComments(topic string) []Comment {
	return p.TopicRepository.FindByTopic(topic)
}

func (p *TopicService) Save(topic Topic) Topic {
	p.TopicRepository.Save(topic)

	return topic
}
func (p *TopicService) SaveComment(comment Comment) Comment {
	p.TopicRepository.SaveComment(comment)

	return comment
}

func (p *TopicService) Delete(topic Topic) {
	p.TopicRepository.Delete(topic)
}
func (p *TopicService) DeleteComment(topic string, comment []Comment) {
	p.TopicRepository.DeleteComment(topic, comment)
}
