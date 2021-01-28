package util

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/errors"
	"gitlab.visionet.co.id/pokota/xanadu/LoyaltyService/constants"
	proto "gitlab.visionet.co.id/pokota/xanadu/LoyaltyService/proto/LoyaltyService"
)

var redisClient *redis.Client

func SetRedisClient(redis *redis.Client) {
	redisClient = redis
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func GetUserData(token string) (*proto.ResponseLogin, error) {

	log.Log("Enter User Data Redis")

	responseRedis,err := redisClient.Get("full:"+token).Bytes()
	if err!=nil{
		return nil,errors.Unauthorized(constants.API_NAME,UnhandledError{
			ErrorMessage:fmt.Sprintf("invalid token ")}.Error())
	}

	responses := proto.ResponseLogin{}
	err = json.Unmarshal(responseRedis,&responses)
	if err!=nil{
		return nil,errors.Unauthorized(constants.API_NAME,UnhandledError{
			ErrorMessage:fmt.Sprintf("invalid token ")}.Error())
	}
	if err!=nil{
		return nil,errors.BadRequest(constants.API_NAME,UnhandledError{"invalid token"}.ErrorMessage)
	}
	return &responses,nil
}
