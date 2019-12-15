package main

/*import (
	forums "Go_Gingonic_Redis/forum"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

//CacheServerAPi :  Is in charge to provide all the posible transactions to the router
type CacheServerAPI struct {
	redisClient *redis.Client
}

//ProvideCacheAPI :  Provides to the router the entry api
func ProvideCacheAPI(redisClient *redis.Client) CacheServerAPI {
	return CacheServerAPI{redisClient: redisClient}
}

func MarshalBinary(e interface{}) ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func UnmarshalBinary(data []byte, e interface{}) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}

	return nil
}
func (p *CacheServerAPI) GetTopics(c *gin.Context) {
	// Open Redis Connection

	result, err := p.redisClient.Get("topics").Result()
	if err != nil {
		topic := forums.Topics{}
		response, err := http.Get("http://forum:3010/topics")
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		} else {
			defer response.Body.Close()
			err := json.NewDecoder(response.Body).Decode(&topic)
			if err != nil {
				panic(err)
			}

			topics, _ := MarshalBinary(&topic)
			p.redisClient.Set("topics", topics, 0)

			c.JSON(200, topic)
		}

	} else {
		var topics forums.Topics
		fmt.Print(result)
		err := UnmarshalBinary([]byte(result), &topics)
		if err != nil {
			panic("Not get propertly")
		}
		fmt.Println("Desde redis redis")
		c.JSON(200, topics)
	}

}

func (p *CacheServerAPI) GetKeys(c *gin.Context) {
	redisClient := NewClient()
	res := redisClient.Keys("topics*")
	fmt.Print(res)
}

func Set(c *gin.Context) {

}
func Del(c *gin.Context) {

}
func Update(c *gin.Context) {

}
*/
