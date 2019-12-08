package topics

//ToTopic : mapper single objectDTO to gorm struct
func ToTopic(TopicDTO TopicDTO) Topic {
	return Topic{
		Author:     TopicDTO.Author,
		TopicTitle: TopicDTO.TopicTitle,
	}
}

/*ToTopicDTO : mapper to gorm struct to single object DTO */
func ToTopicDTO(topic Topic) TopicDTO {
	return TopicDTO{
		Author:     topic.Author,
		TopicTitle: topic.TopicTitle,
	}
}

//ToTopicDTOs : Mapper to gorm stuct to array of DTO objects
func ToTopicDTOs(topic []Topic) []TopicDTO {
	topicsDtos := make([]TopicDTO, len(topic))

	for i, itm := range topic {
		topicsDtos[i] = ToTopicDTO(itm)
	}

	return topicsDtos
}
