package city

import (
	"context"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/dto"
)

type CityGeneric interface {
	GetAll(ctx context.Context, city dto.FilterPaging)(*dto.ListCities, error)
	FindById(context.Context, int32)(*dto.City, error)
	Create(ctx context.Context, city dto.City)(*dto.City , error)
	Update(ctx context.Context, city dto.City)(*dto.City , error)
	Delete(ctx context.Context, id int32)(*bool , error)
	GetAllDistrict(ctx context.Context, district dto.FilterPaging)(*dto.ListDistrict, error)
	DistrictFindById(context.Context, int32)(*dto.District, error)
	CreateDistrict(ctx context.Context, district dto.District)(*dto.District , error)
	UpdateDistrict(ctx context.Context, district dto.District)(*dto.District , error)
	DeleteDistrict(ctx context.Context, id int32)(*bool , error)
	ListProvince(ctx context.Context, paging dto.FilterPaging)([]*dto.Province,*int32, error)
	AddProvince(ctx context.Context, province dto.Province)(*dto.Province, error)
	EditProvince(ctx context.Context, province dto.Province)(*dto.Province, error)
	DetailProvince(ctx context.Context, province dto.Province)(*dto.Province, error)
	RemoveProvince(ctx context.Context, province dto.Province)error
}