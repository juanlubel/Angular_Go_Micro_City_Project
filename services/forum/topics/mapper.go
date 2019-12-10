package topics

//ToTopic : mapper single objectDTO to gorm struct
func ToTopic(TopicDTO TopicDTO) Topic {
	return Topic{
		Author:     TopicDTO.Author,
		TopicTitle: TopicDTO.TopicTitle,
	}
}

//ToComment : mapper single objectDTO to gorm struct
func ToComment(CommentDTO CommentDTO) Comment {
	return Comment{
		Author:      CommentDTO.Author,
		Body:        CommentDTO.Body,
		TopicTittle: CommentDTO.TopicTittle,
	}
}

/*ToTopicDTO : mapper to gorm struct to single object DTO */
func ToTopicDTO(topic Topic) TopicDTO {
	return TopicDTO{
		Author:     topic.Author,
		TopicTitle: topic.TopicTitle,
	}
}
func ToCommentDTO(comment Comment) CommentDTO {
	return CommentDTO{
		Author:      comment.Author,
		Body:        comment.Body,
		TopicTittle: comment.TopicTittle,
	}
}

func ToTopicWithCommentsDTO(comment []Comment, topic Topic) TopicWithCommentsDTO {
	return TopicWithCommentsDTO{
		Topic:    ToTopicDTO(topic),
		Comments: ToCommentDTOs(comment),
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
func ToCommentDTOs(comment []Comment) []CommentDTO {

	commentDtos := make([]CommentDTO, len(comment))

	for i, itm := range comment {
		commentDtos[i] = ToCommentDTO(itm)
	}

	return commentDtos
}
