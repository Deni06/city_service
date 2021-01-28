package go_micro_srv_CityService

import "gitlab.visionet.co.id/pokota/xanadu/CityService/util"

func (city *CityData) Validate() error {
	if city.CityName==""{
		return util.UnhandledError{ErrorMessage:"city name required"}
	}
	if city.ProvinceId==0{
		return util.UnhandledError{ErrorMessage:"province id required"}
	}
	return nil
}

func (city *CityData) ValidateUpdate() error {
	if city.CityName==""{
		return util.UnhandledError{ErrorMessage:"city name required"}
	}
	if city.ProvinceId==0{
		return util.UnhandledError{ErrorMessage:"province id required"}
	}
	if city.CityId==0{
		return util.UnhandledError{ErrorMessage:"city id required"}
	}
	return nil
}

func (city *CityKey) Validate() error {
	if city.Id==0{
		return util.UnhandledError{ErrorMessage:"id required"}
	}
	return nil
}

func (district *DistrictData) ValidateAddDistrict() error {
	if district.DistrictName==""{
		return util.UnhandledError{ErrorMessage:"district name required"}
	}
	if district.CityId==0{
		return util.UnhandledError{ErrorMessage:"city id required"}
	}
	return nil
}

func (district *DistrictData) ValidateUpdateDistrict() error {
	if district.DistrictName==""{
		return util.UnhandledError{ErrorMessage:"district name required"}
	}
	if district.CityId==0{
		return util.UnhandledError{ErrorMessage:"city id required"}
	}
	if district.DistrictId==0{
		return util.UnhandledError{ErrorMessage:"district id required"}
	}
	return nil
}

func (city *FilterPaging) ValidateGet() error {
	if city.Pagination.Page!=0 || city.Pagination.PerPage!=0{
		if city.Pagination.Page==0 {
			return util.UnhandledError{ErrorMessage:"page number required"}
		}
		if city.Pagination.PerPage==0{
			return util.UnhandledError{ErrorMessage:"page size required"}
		}
	}
	return nil
}