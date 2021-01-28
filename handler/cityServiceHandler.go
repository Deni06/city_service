package handler

import (
	"context"
	"fmt"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/constant"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/dto"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/manager/city"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/util"

	cityservice "gitlab.visionet.co.id/pokota/xanadu/CityService/proto/cityservice"
)

type CityService struct{}

var manager = cityManager.CityManagerPostgreImpl{}


// Call is a single request handler called via client.Call or the generated client code
func (e *CityService) Healthcheck(ctx context.Context, req *cityservice.EmptyRequest, rsp *cityservice.Response) error {
	util.WriteLogMain("Call Healthcheck Service")
	rsp.Msg = "Success"
	return nil
}

func (e *CityService) List(ctx context.Context, req *cityservice.FilterPaging, rsp *cityservice.ListCityResponse) error {
	util.WriteLogMain("Call List City")
	cities := make([]*cityservice.CityData,0)
	paging := cityservice.Paging{}

	if req.Pagination==nil{
		req.Pagination = &cityservice.Paging{}
	}
	if req.Filter==nil{
		req.Filter = &cityservice.FilterKey{}
	}
	input := dto.FilterPaging{
		Filter: dto.FilterCity{
			Keyword:req.Filter.Keyword,
			ProvinceID:req.Filter.ProvinceId,
		},
		Paging: dto.Paging{
			Page:req.Pagination.Page,
			PerPage:req.Pagination.PerPage,
		},
	}
	err := req.ValidateGet()
	if err!=nil{
		return err
	}
	results,err:=manager.GetAll(ctx, input)
	if err!=nil{
		return err
	}
	for _,val := range results.City{
		citymodel := cityservice.CityData{}
		citymodel.CityName = val.CityName
		citymodel.CityId = val.CityID
		citymodel.ProvinceId = val.ProvinceID
		citymodel.ProvinceName = val.ProvinceName
		cities = append(cities, &citymodel)
	}
	paging.PerPage = results.Paging.PerPage
	paging.Page = results.Paging.Page
	paging.Count = results.Paging.Count
	rsp.Paging = &paging
	rsp.Cities = cities
	return nil
}

func (e *CityService) Detail(ctx context.Context, req *cityservice.CityKey, rsp *cityservice.CityData) error {
	util.WriteLogMain("Received Detail City request")
	err := req.Validate()
	if err!=nil{
		return err
	}
	result , err := manager.FindById(ctx, req.Id)
	if err!=nil{
		return err
	}
	rsp.CityId = result.CityID
	rsp.CityName = result.CityName
	rsp.ProvinceId = result.ProvinceID
	rsp.ProvinceName = result.ProvinceName

	return nil
}

func (e *CityService) Add(ctx context.Context, req *cityservice.CityData, rsp *cityservice.CityData) error {
	util.WriteLogMain("Received Add City request")
	util.WriteLogMain(fmt.Sprintf("validate request : %v",req))
	err := req.Validate()
	if err!=nil{
		return err
	}
	input := dto.City{CityName: req.CityName, ProvinceID:req.ProvinceId}
	result , err := manager.Create(ctx, input)
	if err!=nil{
		return err
	}
	rsp.CityId = result.CityID
	rsp.CityName = result.CityName
	rsp.ProvinceId = result.ProvinceID
	rsp.ProvinceName = result.ProvinceName

	return nil
}

func (e *CityService) Edit(ctx context.Context, req *cityservice.CityData, rsp *cityservice.CityData) error {
	util.WriteLogMain("Received Edit City request")
	util.WriteLogMain(fmt.Sprintf("validate request : %v",req))
	err := req.ValidateUpdate()
	if err!=nil{
		return err
	}
	input := dto.City{CityName: req.CityName, CityID:req.CityId,ProvinceID:req.ProvinceId}
	result , err := manager.Update(ctx, input)
	if err!=nil{
		return err
	}
	rsp.CityId = result.CityID
	rsp.CityName = result.CityName
	rsp.ProvinceId = result.ProvinceID
	rsp.ProvinceName = result.ProvinceName

	return nil
}

func (e *CityService) Remove(ctx context.Context, req *cityservice.CityKey, rsp *cityservice.IsRemoved) error {
	util.WriteLogMain("Received Remove City request")
	err := req.Validate()
	if err!=nil{
		return err
	}
	result , err := manager.Delete(ctx, req.Id)
	if err!=nil{
		return err
	}
	rsp.IsRemoved=*result
	return nil
}

func (e *CityService) ListDistrict(ctx context.Context, req *cityservice.FilterPaging, rsp *cityservice.ListDistrictResponse) error {
	util.WriteLogMain("Call List District")
	if req.Pagination==nil{
		req.Pagination = &cityservice.Paging{}
	}
	if req.Filter==nil{
		req.Filter = &cityservice.FilterKey{}
	}
	input := dto.FilterPaging{
		Filter: dto.FilterCity{
			Keyword:req.Filter.Keyword,
			CityID:req.Filter.CityId,
		},
		Paging: dto.Paging{
			Page:req.Pagination.Page,
			PerPage:req.Pagination.PerPage,
		},
	}
	err := req.ValidateGet()
	if err!=nil{
		return err
	}
	results,err:=manager.GetAllDistrict(ctx, input)
	if err!=nil{
		return err
	}
	districts := make([]*cityservice.DistrictData,0)
	paging := cityservice.Paging{}

	for _,val := range results.District{
		districtmodel := cityservice.DistrictData{}
		districtmodel.DistrictName = val.DistrictName
		districtmodel.DistrictId = val.DistrictID
		districtmodel.CityId = val.CityID
		districtmodel.CityName = val.CityName
		districts = append(districts, &districtmodel)
	}
	paging.PerPage = results.Paging.PerPage
	paging.Page = results.Paging.Page
	paging.Count = results.Paging.Count
	rsp.Paging = &paging
	rsp.District = districts

	return nil
}

func (e *CityService) DetailDistrict(ctx context.Context, req *cityservice.CityKey, rsp *cityservice.DistrictData) error {
	util.WriteLogMain("Received Detail District request")
	err := req.Validate()
	if err!=nil{
		return err
	}
	result , err := manager.DistrictFindById(ctx, req.Id)
	if err!=nil{
		return err
	}
	rsp.DistrictId = result.DistrictID
	rsp.DistrictName = result.DistrictName
	rsp.CityId = result.CityID
	rsp.CityName = result.CityName

	return nil
}

func (e *CityService) AddDistrict(ctx context.Context, req *cityservice.DistrictData, rsp *cityservice.DistrictData) error {
	util.WriteLogMain("Received Add District request")
	util.WriteLogMain(fmt.Sprintf("validate request : %v",req))
	err := req.ValidateAddDistrict()
	if err!=nil{
		return err
	}
	input := dto.District{DistrictName: req.DistrictName, CityID:req.CityId}
	result , err := manager.CreateDistrict(ctx, input)
	if err!=nil{
		return err
	}
	rsp.DistrictName = result.DistrictName
	rsp.DistrictId = result.DistrictID
	rsp.CityId = result.CityID

	return nil
}

func (e *CityService) EditDistrict(ctx context.Context, req *cityservice.DistrictData, rsp *cityservice.DistrictData) error {
	util.WriteLogMain("Received Edit District request")
	util.WriteLogMain(fmt.Sprintf("validate request : %v",req))
	err := req.ValidateUpdateDistrict()
	if err!=nil{
		return err
	}
	input := dto.District{DistrictName: req.DistrictName, DistrictID:req.DistrictId, CityID:req.CityId}
	result , err := manager.UpdateDistrict(ctx, input)
	if err!=nil{
		return err
	}
	rsp.DistrictName = result.DistrictName
	rsp.DistrictId = result.DistrictID
	rsp.CityId = result.CityID
	rsp.CityName = result.CityName


	return nil
}

func (e *CityService) RemoveDistrict(ctx context.Context, req *cityservice.CityKey, rsp *cityservice.IsRemoved) error {
	util.WriteLogMain("Received Remove District request")
	err := req.Validate()
	if err!=nil{
		return err
	}
	result , err := manager.DeleteDistrict(ctx, req.Id)
	if err!=nil{
		return err
	}
	rsp.IsRemoved=*result
	return nil
}

func (e *CityService) DetailProvince(ctx context.Context, req *cityservice.ProvinceData, rsp *cityservice.ProvinceData) error {
	util.WriteLogMain("Received DetailProvince request")
	input :=  dto.Province{
		ID:req.Id,
	}
	result , err := manager.DetailProvince(ctx, input)
	if err!=nil{
		return err
	}
	rsp.ProvinceName = result.ProvinceName
	rsp.Id = result.ID
	return nil
}

func (e *CityService) ListProvince(ctx context.Context, req *cityservice.FilterPaging, rsp *cityservice.ListProvinceData) error {
	util.WriteLogMain("Received ListProvince request")
	filter:= dto.FilterCity{}
	if req.Filter !=nil{
		filter.Keyword = req.Filter.Keyword
		filter.ProvinceID =req.Filter.ProvinceId
	}
	pagination := dto.Paging{}
	if req.Pagination !=nil{
		pagination.Page = req.Pagination.Page
		pagination.PerPage = req.Pagination.PerPage
	}else{
		pagination.Page = constant.DEFAULT_PAGE
		pagination.PerPage = constant.DEFAULT_PERPAGE
	}
	input :=  dto.FilterPaging{
		Paging:pagination,
		Filter:filter,
	}
	results ,count, err := manager.ListProvince(ctx, input)
	if err!=nil{
		return err
	}

	resultService := make([]*cityservice.ProvinceData,0)

	for _, data := range results{
		result := cityservice.ProvinceData{
			Id:data.ID,
			ProvinceName:data.ProvinceName,
		}
		resultService = append(resultService, &result)
	}

	rsp.Provinces = resultService
	rsp.Paging = &cityservice.Paging{
		Count:*count,
		Page:pagination.Page,
		PerPage:pagination.PerPage,
	}
	return nil
}

func (e *CityService) EditProvince(ctx context.Context, req *cityservice.ProvinceData, rsp *cityservice.ProvinceData) error {
	util.WriteLogMain("Received EditProvince request")
	input :=  dto.Province{
		ID:req.Id,
		ProvinceName:req.ProvinceName,
	}
	result , err := manager.EditProvince(ctx, input)
	if err!=nil{
		return err
	}
	rsp.ProvinceName = result.ProvinceName
	rsp.Id = result.ID
	return nil
}

func (e *CityService) AddProvince(ctx context.Context, req *cityservice.ProvinceData, rsp *cityservice.ProvinceData) error {
	util.WriteLogMain("Received AddProvince request")
	input :=  dto.Province{
		ID:req.Id,
		ProvinceName:req.ProvinceName,
	}
	result , err := manager.AddProvince(ctx, input)
	if err!=nil{
		return err
	}
	rsp.ProvinceName = result.ProvinceName
	rsp.Id = result.ID
	return nil
}

func (e *CityService) RemoveProvince(ctx context.Context, req *cityservice.ProvinceData, rsp *cityservice.Response) error {
	util.WriteLogMain("Received RemoveProvince request")
	input :=  dto.Province{
		ID:req.Id,
	}
	err := manager.RemoveProvince(ctx, input)
	if err!=nil{
		return err
	}
	rsp.Msg ="success"
	return nil
}