package dto

import "encoding/json"

type City struct {
	CityID int32
	CityName string
	ProvinceID int32
	ProvinceName string
}

type District struct {
	DistrictID int32
	DistrictName string
	CityID int32
	CityName string
}

type FilterCity struct {
	ProvinceID int32
	CityID int32
	Keyword string
}

type ListCities struct {
	City []City
	Paging Paging
}

type ListDistrict struct {
	District []District
	Paging Paging
}

type Paging struct {
	Page int32
	PerPage int32
	Count int32
}

type FilterPaging struct {
	Filter FilterCity
	Paging Paging
}

type Province struct {
	ID int32
	ProvinceName string
}

func (city City) ToString()( string,error) {
	out, err := json.Marshal(city)
	if err != nil {
		return "", err
	}
	return string(out),nil
}

