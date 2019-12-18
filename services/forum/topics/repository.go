package topics

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql
)

// TopicRepository : obtain the context of the database
type TopicRepository struct {
	DB *gorm.DB
}

// ProvideTopicRepostiory : provides the repository to use the db context
func ProvideTopicRepostiory(DB *gorm.DB) TopicRepository {
	return TopicRepository{DB: DB}
}

//FindAll : obtain from db all the data of the indicated gorm model
func (p *TopicRepository) FindAll() []Topic {
	var topic []Topic
	p.DB.Find(&topic)

	return topic
}

/* func (p *TopicRepository) FindByBankAndName(owner string, bank string) Topic {
	var topic Topic
	println(owner)
	p.DB.Where("Owner = ? AND Bank=?", owner, bank).First(&topic)

	return topic
} */

//FindByOwner :  obtain from db the data that contains the param name
func (p *TopicRepository) FindByOwner(owner string) Topic {
	var topic Topic
	println(owner)
	p.DB.Where("topic_title = ?", owner).First(&topic)

	return topic
}
func (p *TopicRepository) FindByTopic(topicTittle string) []Comment {
	var comment []Comment
	/* var topic Topic */
	/* p.DB.Find(&comment) */
	/* p.DB.Model(&topic).Related(&comment, "TopicTitle") */
	p.DB.Where("topic_tittle = ?", topicTittle).Find(&comment)
	return comment
}

//Save : stores the changes occurred when creating or updating entries in the database
func (p *TopicRepository) Save(topic Topic) Topic {
	p.DB.Save(&topic)

	return topic
}
func (p *TopicRepository) SaveComment(Comment Comment) Comment {
	p.DB.Save(&Comment)

	return Comment
}

// Delete : Deletes from the data base
func (p *TopicRepository) Delete(topic Topic) {
	p.DB.Unscoped().Delete(&topic)
}
func (p *TopicRepository) DeleteComment(topicTittle string, comment []Comment) {
	p.DB.Where("topic_tittle = ?", topicTittle).Unscoped().Delete(&comment)
}
