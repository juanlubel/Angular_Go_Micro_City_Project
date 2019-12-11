package topics

import (
	"Go_Gingonic_Server/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TopicAPI :  Is in charge to provide all the posible transactions to the router
type TopicAPI struct {
	TopicService TopicService
	JWT          *utils.JWT
}

//ProvideTopicAPI :  Provides to the router the entry api
func ProvideTopicAPI(p TopicService, jwt *utils.JWT) TopicAPI {
	return TopicAPI{TopicService: p, JWT: jwt}
}

/* Simple Crud Opperations */

//FindAll : The method to obtain all the data of the table
func (p *TopicAPI) FindAll(c *gin.Context) {
	topics := p.TopicService.FindAll()

	c.JSON(http.StatusOK, gin.H{"topics": ToTopicDTOs(topics)})
}

// Create : Create a new entry in the table (Used in Register Functionality)
func (p *TopicAPI) Create(c *gin.Context) {
	var fullTopicDTO TopicDTO
	err := c.BindJSON(&fullTopicDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdTopic := p.TopicService.Save(ToTopic(fullTopicDTO))

	c.JSON(http.StatusOK, gin.H{"topic": ToTopicDTO(createdTopic)})
}

func (p *TopicAPI) CreateComment(c *gin.Context) {
	var CommentDTO CommentDTO
	err := c.BindJSON(&CommentDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdComment := p.TopicService.SaveComment(ToComment(CommentDTO))

	c.JSON(http.StatusOK, gin.H{"Comment": ToCommentDTO(createdComment)})
}

//FindByOwner :  The method to obtain specific data
func (p *TopicAPI) FindByOwner(c *gin.Context) {
	owner := c.Param("name")
	topic := p.TopicService.FindByOwner(owner)
	/* c.JSON(http.StatusOK, gin.H{"topic": ToTopicDTO(account)}) */
	comments := p.TopicService.FindHisComments(owner)
	c.JSON(http.StatusOK, gin.H{"topic": ToTopicWithCommentsDTO(comments, topic)})
}

/*
// Update :  Updates an specific content of the table
func (p *TopicAPI) Update(c *gin.Context) {
	var fullTopicDTO TopicDTO
	err := c.BindJSON(&fullTopicDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	owner, _ := strconv.Atoi(c.Param("owner"))
	account := p.TopicService.FindByOwner(string(owner))
	if account == (Topic{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	account.Bank = fullTopicDTO.Bank
	p.TopicService.Save(account)

	c.Status(http.StatusOK)
}
// Delete :  Deletes an specific content of the table
func (p *TopicAPI) Delete(c *gin.Context) {
	owner, _ := strconv.Atoi(c.Param("owner"))
	account := p.TopicService.FindByOwner(string(owner))
	if account == (Topic{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.TopicService.Delete(account)

	c.Status(http.StatusOK)
} */
