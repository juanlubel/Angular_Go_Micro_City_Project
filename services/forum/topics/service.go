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

/* func (p *TopicService) FindByBankAndName(owner string, bank string) Topic {
	return p.TopicRepository.FindByBankAndName(owner, bank)
} */
/* func (p *TopicService) FindByOwner(owner string) Topic {
	return p.TopicRepository.FindByOwner(owner)
} */

func (p *TopicService) Save(topic Topic) Topic {
	p.TopicRepository.Save(topic)

	return topic
}

func (p *TopicService) Delete(topic Topic) {
	p.TopicRepository.Delete(topic)
}
