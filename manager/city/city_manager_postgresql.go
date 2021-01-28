package cityManager

import (
	"context"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/dto"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/repository/city"
)

type CityManagerPostgreImpl struct {}

var repo = city.CityGormPostgreImpl{}

func (impl CityManagerPostgreImpl) GetAll(ctx context.Context, city dto.FilterPaging) (*dto.ListCities, error){
	results,err :=repo.GetAll(ctx, city)
	if err!=nil{
		return nil, err
	}
	return results,nil
}

func (impl CityManagerPostgreImpl)FindById(ctx context.Context,id int32) (*dto.City, error){
	result, err := repo.FindById(ctx, id)
	if err!=nil{
		return nil, err
	}
	return result,nil
}

func (impl CityManagerPostgreImpl)Create(ctx context.Context,city dto.City)(*dto.City, error){
	result, err := repo.Create(ctx,city)
	if err!=nil{
		return nil, err
	}
	return result, nil
}

func (impl CityManagerPostgreImpl)Update(ctx context.Context,city dto.City)(*dto.City, error){
	result, err := repo.Update(ctx,city)
	if err!=nil{
		return nil, err
	}
	return result, nil
}

func (impl CityManagerPostgreImpl)Delete(ctx context.Context,id int32)(*bool , error){
	result, err := repo.Delete(ctx,id)
	if err!=nil{
		return nil, err
	}
	return result,nil
}

func (impl CityManagerPostgreImpl) GetAllDistrict(ctx context.Context, district dto.FilterPaging) (*dto.ListDistrict, error){
	results,err :=repo.GetAllDistrict(ctx, district)
	if err!=nil{
		return nil, err
	}
	return results,nil
}

func (impl CityManagerPostgreImpl)DistrictFindById(ctx context.Context,id int32) (*dto.District, error){
	result, err := repo.DistrictFindById(ctx, id)
	if err!=nil{
		return nil, err
	}
	return result,nil
}

func (impl CityManagerPostgreImpl)CreateDistrict(ctx context.Context,district dto.District)(*dto.District , error){
	result, err := repo.CreateDistrict(ctx,district)
	if err!=nil{
		return nil, err
	}
	return result, nil
}

func (impl CityManagerPostgreImpl)UpdateDistrict(ctx context.Context,district dto.District)(*dto.District, error){
	result, err := repo.UpdateDistrict(ctx,district)
	if err!=nil{
		return nil, err
	}
	return result, nil
}

func (impl CityManagerPostgreImpl)DeleteDistrict(ctx context.Context,id int32)(*bool , error){
	result, err := repo.DeleteDistrict(ctx,id)
	if err!=nil{
		return nil, err
	}
	return result,nil
}

func (impl CityManagerPostgreImpl)ListProvince(ctx context.Context,paging dto.FilterPaging)([]*dto.Province ,*int32, error){
	result, count,err := repo.ListProvince(ctx,paging)
	if err!=nil{
		return nil, nil,err
	}
	return result,count,nil
}

func (impl CityManagerPostgreImpl)AddProvince(ctx context.Context,province dto.Province)(*dto.Province , error){
	result, err := repo.AddProvince(ctx,province)
	if err!=nil{
		return nil, err
	}
	return result,nil
}

func (impl CityManagerPostgreImpl)EditProvince(ctx context.Context,province dto.Province)(*dto.Province , error){
	result, err := repo.EditProvince(ctx,province)
	if err!=nil{
		return nil, err
	}
	return result,nil
}

func (impl CityManagerPostgreImpl)DetailProvince(ctx context.Context,province dto.Province)(*dto.Province , error){
	result, err := repo.DetailProvince(ctx,province)
	if err!=nil{
		return nil, err
	}
	return result,nil
}

func (impl CityManagerPostgreImpl)RemoveProvince(ctx context.Context,province dto.Province)error{
	err := repo.RemoveProvince(ctx,province)
	if err!=nil{
		return err
	}
	return nil
}