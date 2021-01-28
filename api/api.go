package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/constant"
	proto "gitlab.visionet.co.id/pokota/xanadu/CityService/proto/cityservice"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/util"

	"gitlab.visionet.co.id/pokota/xanadu/CityService/middleware"
	"log"
	"net/http"
	"os"
	"strconv"

	"context"
)

type City struct {
	Client proto.CityService
}

func (s *City) Healthcheck(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.Healthcheck API request")

	response, err := s.Client.Healthcheck(ctx, &proto.EmptyRequest{})
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)

	return nil
}

func (s *City) List(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.List API request")
	reqObj := proto.FilterPaging{}
	err := json.Unmarshal([]byte(req.Body), &reqObj)
	if err!=nil{
		return errors.BadRequest(constant.API_NAME,err.Error())
	}
	response, err := s.Client.List(ctx, &reqObj)
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)

	return nil
}

func (s *City) Detail(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.Detail API request")
	reqObj := proto.CityKey{}
	err := json.Unmarshal([]byte(req.Body), &reqObj)
	if err!=nil{
		return errors.BadRequest(constant.API_NAME,err.Error())
	}
	response, err := s.Client.Detail(ctx, &reqObj)
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)

	return nil
}

func (s *City) Add(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.Add API request")
	reqObj := proto.CityData{}
	err := json.Unmarshal([]byte(req.Body), &reqObj)
	if err!=nil{
		return errors.BadRequest(constant.API_NAME,err.Error())
	}
	response, err := s.Client.Add(ctx, &reqObj)
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)

	return nil
}

func (s *City) Edit(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.Edit API request")
	reqObj := proto.CityData{}
	err := json.Unmarshal([]byte(req.Body), &reqObj)
	if err!=nil{
		return errors.BadRequest(constant.API_NAME,err.Error())
	}
	response, err := s.Client.Edit(ctx, &reqObj)
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)

	return nil
}

func (s *City) Remove(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.Remove API request")
	reqObj := proto.CityKey{}
	err := json.Unmarshal([]byte(req.Body), &reqObj)
	if err!=nil{
		return errors.BadRequest(constant.API_NAME,err.Error())
	}
	response, err := s.Client.Remove(ctx, &reqObj)
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)
	return nil
}

func (s *City) ListDistrict(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.ListDistrict API request")
	reqObj := proto.FilterPaging{}
	err := json.Unmarshal([]byte(req.Body), &reqObj)
	if err!=nil{
		return errors.BadRequest(constant.API_NAME,err.Error())
	}
	response, err := s.Client.ListDistrict(ctx, &reqObj)
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)

	return nil
}

func (s *City) DetailDistrict(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.DetailDistrict API request")
	reqObj := proto.CityKey{}
	err := json.Unmarshal([]byte(req.Body), &reqObj)
	if err!=nil{
		return errors.BadRequest(constant.API_NAME,err.Error())
	}
	response, err := s.Client.DetailDistrict(ctx, &reqObj)
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)

	return nil
}

func (s *City) AddDistrict(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.AddDistrict API request")
	reqObj := proto.DistrictData{}
	err := json.Unmarshal([]byte(req.Body), &reqObj)
	if err!=nil{
		return errors.BadRequest(constant.API_NAME,err.Error())
	}
	response, err := s.Client.AddDistrict(ctx, &reqObj)
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)

	return nil
}

func (s *City) EditDistrict(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.EditDistrict API request")
	reqObj := proto.DistrictData{}
	err := json.Unmarshal([]byte(req.Body), &reqObj)
	if err!=nil{
		return errors.BadRequest(constant.API_NAME,err.Error())
	}
	response, err := s.Client.EditDistrict(ctx, &reqObj)
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)

	return nil
}

func (s *City) RemoveDistrict(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.RemoveDistrict API request")
	reqObj := proto.CityKey{}
	err := json.Unmarshal([]byte(req.Body), &reqObj)
	if err!=nil{
		return errors.BadRequest(constant.API_NAME,err.Error())
	}
	response, err := s.Client.RemoveDistrict(ctx, &reqObj)
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)
	return nil
}

func (s *City) ListProvince(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.ListProvince API request")
	reqObj := proto.FilterPaging{}
	err := json.Unmarshal([]byte(req.Body), &reqObj)
	if err!=nil{
		return errors.BadRequest(constant.API_NAME,err.Error())
	}
	response, err := s.Client.ListProvince(ctx, &reqObj)
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)

	return nil
}

func (s *City) DetailProvince(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.DetailProvince API request")
	reqObj := proto.ProvinceData{}
	err := json.Unmarshal([]byte(req.Body), &reqObj)
	if err!=nil{
		return errors.BadRequest(constant.API_NAME,err.Error())
	}
	response, err := s.Client.DetailProvince(ctx, &reqObj)
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)

	return nil
}

func (s *City) AddProvince(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.AddProvince API request")
	reqObj := proto.ProvinceData{}
	err := json.Unmarshal([]byte(req.Body), &reqObj)
	if err!=nil{
		return errors.BadRequest(constant.API_NAME,err.Error())
	}
	response, err := s.Client.AddProvince(ctx, &reqObj)
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)

	return nil
}

func (s *City) EditProvince(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.EditProvince API request")
	reqObj := proto.ProvinceData{}
	err := json.Unmarshal([]byte(req.Body), &reqObj)
	if err!=nil{
		return errors.BadRequest(constant.API_NAME,err.Error())
	}
	response, err := s.Client.EditProvince(ctx, &reqObj)
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)

	return nil
}

func (s *City) RemoveProvince(ctx context.Context, req *api.Request, rsp *api.Response) error {
	util.WriteLogApi("Received city.RemoveProvince API request")
	reqObj := proto.ProvinceData{}
	err := json.Unmarshal([]byte(req.Body), &reqObj)
	if err!=nil{
		return errors.BadRequest(constant.API_NAME,err.Error())
	}
	response, err := s.Client.RemoveProvince(ctx, &reqObj)
	if err != nil {
		return err
	}
	rsp.StatusCode = http.StatusOK
	rspObject,err := json.Marshal(response)
	if err!=nil{
		return errors.InternalServerError(constant.API_NAME,err.Error())
	}
	rsp.Body = string(rspObject)
	return nil
}

var redisServe *redis.Client

func main() {
	redisHost := constant.REDIS_HOST
	if os.Getenv(constant.REDIS_HOST_ENV) != "" {
		redisHost = os.Getenv(constant.REDIS_HOST_ENV)
	}
	redisDB := constant.REDIS_DB
	if os.Getenv(constant.REDIS_DB_ENV) != "" {
		redisDB, _ = strconv.Atoi(os.Getenv(constant.REDIS_DB_ENV))
	}
	redisPassword := constant.REDIS_PASSWORD
	if os.Getenv(constant.REDIS_PASSWORD_ENV) != "" {
		redisPassword = os.Getenv(constant.REDIS_PASSWORD_ENV)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword, // no password set
		DB:       redisDB,       // use default DB
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}
	redisServe = redisClient
	util.SetRedisClient(redisClient)

	service := micro.NewService(
		micro.Name(constant.API_NAME),
		micro.WrapHandler(middleware.NewValidatorWrapper(redisServe)),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&City{Client: proto.NewCityService(constant.SERVICE_NAME, service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
