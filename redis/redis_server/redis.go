package redis_server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

//MainAPi :  Struct of the redis package
type MainAPi struct{}

//NewClient :  Creates a new client to use Redis
func NewClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return redisClient
}

//Ping :  Make a ping to redis
func Ping(client *redis.Client) (string, error) {
	result, err := client.Ping().Result()

	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}

//GetRedis :  Get data stored in redis
func (p *MainAPi) GetRedis(client *redis.Client, key string) (string, error) {
	result, err := client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

//GetDB :  Get data stored in other servers
func (p *MainAPi) GetDB(client *redis.Client, url string, data interface{}) interface{} {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	return &data

}

//MarshalBinary :  Transform structs to binary
func (p *MainAPi) MarshalBinary(e interface{}) ([]byte, error) {
	return json.Marshal(e)
}

//UnmarshalBinary : Transform Binary data to structs
func (p *MainAPi) UnmarshalBinary(data []byte, e interface{}) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}

	return nil
}
