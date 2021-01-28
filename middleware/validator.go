package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-log/log"
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/server"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/constant"
	proto "gitlab.visionet.co.id/pokota/xanadu/CityService/proto/cityservice"
	util "gitlab.visionet.co.id/pokota/xanadu/CityService/util"
)

func NewValidatorWrapper(redis *redis.Client) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			//skip user login because get token form user login
			if req.Method()!="City.List" && req.Method()!="City.Healthcheck" && req.Method()!="City.Detail" && req.Method()!="City.ListDistrict" && req.Method()!="City.DetailDistrict" && req.Method()!="City.ListProvince" && req.Method()!="City.DetailProvince" {
				headers := req.Header()
				token := headers["Authorization"]
				roleByte,err := redis.Get(token).Bytes()
				if err!=nil{
					log.Log(fmt.Sprintf("error while get token to redis : %v ",err.Error()))
					return errors.Unauthorized(constant.API_NAME,util.UnhandledError{
						ErrorMessage:fmt.Sprintf("invalid token or you dont have access for this service %v ", req.Method())}.Error())
				}

				roles := []*proto.RoleData{}
				err = json.Unmarshal(roleByte,&roles)
				if err!=nil{
					log.Log(fmt.Sprintf("error while unmarshal roles from redis : %v ",err.Error()))
					return errors.Unauthorized(constant.API_NAME,util.UnhandledError{
						ErrorMessage:fmt.Sprintf("invalid token or you dont have access for this service %v ", req.Method())}.Error())
				}
				for _,data := range roles{
					for _,permission := range data.Permissions{
						if permission.PermissionName == req.Method(){
							return h(ctx, req, rsp)
						}
					}
				}
				return errors.Unauthorized(constant.API_NAME,util.UnhandledError{
					ErrorMessage:fmt.Sprintf("invalid token or you dont have access for this service %v ", req.Method())}.Error())
			}else{
				return h(ctx, req, rsp)
			}
		}
	}
}