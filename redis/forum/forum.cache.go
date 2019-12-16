package forums

import (
	"Go_Gingonic_Redis/redis_server"
	"fmt"

	/* "encoding/json"
	"fmt"
	"net/http"
	"os" */

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

//CacheServerAPI :  Is in charge to provide all the posible transactions to the router
type CacheServerAPI struct {
	redisClient *redis.Client
	redisStruct *redis_server.MainAPi
}

//ProvideCacheAPI :  Provides to the router the entry api
func ProvideCacheAPI(redisClient *redis.Client, redisStruct *redis_server.MainAPi) CacheServerAPI {
	return CacheServerAPI{redisClient: redisClient, redisStruct: redisStruct}
}

//GetTopics : Get the topics stored in redis or by default get the topics from the Forum server
func (p *CacheServerAPI) GetTopics(c *gin.Context) {
	// Open Redis Connection
	result, err := p.redisStruct.GetRedis(p.redisClient, "topics") //get from redis
	if err != nil {                                                //if redis dont have the data get it from db
		topic := Topics{} //Set the data struct
		//get de data from the corresponding server
		result := p.redisStruct.GetDB(p.redisClient, "http://forum:3010/topics", &topic)
		println("get from db")

		//store the data in the correspondig key
		toset, _ := p.redisStruct.MarshalBinary(&result)
		p.redisClient.Set("topics", toset, 0)

		c.JSON(200, result) //return the correspondig data to the client
	} else { //if redis have the data unmarshal them, set the corresponding structure and send it to the client
		topic := Topics{}
		err := p.redisStruct.UnmarshalBinary([]byte(result), &topic)
		println("get from redis")
		if err != nil {
			panic("Not get propertly")
		}
		c.JSON(200, topic)
	}
}

//GetKeys :  Get The Keys storded in redis
func (p *CacheServerAPI) GetKeys(c *gin.Context) {
	res := p.redisClient.Keys("topics*")
	fmt.Print(res)
}

//DeleteKeys : Deletes the specified key from redis
func (p *CacheServerAPI) DeleteKeys(c *gin.Context) {
	res := p.redisClient.Del("topics")
	fmt.Print(res)
}
