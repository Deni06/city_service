package city

import (
	"context"
	"fmt"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/constant"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/dto"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/helper"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/util"
	"strings"
	"time"
)

type CityGormPostgreImpl struct {}

func (impl CityGormPostgreImpl) GetAll(ctx context.Context, city dto.FilterPaging) (*dto.ListCities, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance := helper.GetDBInstance()
	dbInstance.LogMode(true)
	resp := dto.ListCities{}
	paging := dto.Paging{}
	offset := (city.Paging.Page * city.Paging.PerPage) - city.Paging.PerPage
	query := fmt.Sprintf("SELECT c.city_id, c.city_name,c.province_id, coalesce(p.province_name,'') FROM %v c left join %v p on p.province_id = c.province_id where 1 = 1 ", constant.CITY_TABLE_NAME, constant.PROVINCE_TABLE_NAME)
	condition := ""
	if city.Filter.ProvinceID!=0{
		condition = condition+fmt.Sprintf(" And c.province_id=%v",city.Filter.ProvinceID)
	}
	if city.Filter.Keyword!=""{
		cityName := ""
		if strings.Contains(city.Filter.Keyword, "'"){
			for _, data := range strings.Split(city.Filter.Keyword, ""){
				if data =="'"{
					cityName += "'"
					cityName += data
				}else{
					cityName += data
				}
			}
			city.Filter.Keyword = cityName
		}
		key := "%"+strings.ToLower(city.Filter.Keyword)+"%"
		condition += fmt.Sprintf(" AND LOWER(c.city_name) like '%v' ",key)
	}
	newQuery := dbInstance.Raw(query+condition)
	if city.Paging.Page != 0 || city.Paging.PerPage != 0 {
		newQuery = newQuery.Order("c.city_id").Offset(offset).Limit(city.Paging.PerPage)
	}
	util.WriteLogMain(fmt.Sprintf("start execute : %v",query))
	rows, err := newQuery.Rows()
	if err != nil {
		util.WriteLogMain(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	defer rows.Close()

	cities := make([]dto.City, 0)
	for rows.Next() {
		cityModel := new(dto.City)
		err := rows.Scan(&cityModel.CityID, &cityModel.CityName, &cityModel.ProvinceID, &cityModel.ProvinceName)
		if err != nil {
			util.WriteLogMain(err)
		}
		cities = append(cities, *cityModel)
	}
	if err = rows.Err(); err != nil {
		util.WriteLogMain(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	query = fmt.Sprintf("select count(1) from %v c where 1 = 1 ", constant.CITY_TABLE_NAME)
	row := dbInstance.Raw(query+condition).Row()
	count := 0
	err = row.Scan(&count)
	if err != nil {
		util.WriteLogMain(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	resp.City = cities
	paging.Page = city.Paging.Page
	paging.PerPage = city.Paging.PerPage
	paging.Count =int32(count)
	resp.Paging = paging
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return &resp,nil
}

func (impl CityGormPostgreImpl)FindById(ctx context.Context,id int32) (*dto.City, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance := helper.GetDBInstance()
	dbInstance.LogMode(true)
	query := fmt.Sprintf("SELECT c.city_id, c.city_name, c.province_id, coalesce(p.province_name,'') FROM %v c left join %v p on p.province_id = c.province_id where c.city_id = %v", constant.CITY_TABLE_NAME, constant.PROVINCE_TABLE_NAME, id)
	util.WriteLogMain(fmt.Sprintf("start execute : %v",query))
	row := dbInstance.Raw(query).Row()
	city := new(dto.City)
	err := row.Scan(&city.CityID, &city.CityName, &city.ProvinceID,&city.ProvinceName)
	if err != nil {
		util.WriteLogMain(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return city,nil
}

func (impl CityGormPostgreImpl)Create(ctx context.Context,city dto.City)(*dto.City , error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance := helper.GetDBInstance()
	cityName := ""
	if strings.Contains(city.CityName, "'"){
		for _, data := range strings.Split(city.CityName, ""){
			if data =="'"{
				cityName += "'"
				cityName += data
			}else{
				cityName += data
			}
		}
		city.CityName = cityName
	}
	_, err := impl.DetailProvince(ctx, dto.Province{ID: city.ProvinceID})
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, util.UnhandledError{ErrorMessage:fmt.Sprintf("province with id %v not exist", city.ProvinceID)})
	}
	dbInstance.LogMode(true)
	tx := dbInstance.Begin()
	var id int
	query := fmt.Sprintf("INSERT INTO %v(city_name, province_id, created_at) VALUES ('%v', '%v', '%v')",constant.CITY_TABLE_NAME, city.CityName, city.ProvinceID, time.Now().Format(time.RFC3339))
	newQuery := helper.GetCurrentImplementation().GetInsertQueryCity(tx, query, "city_id", constant.CITY_TABLE_NAME)
	err = newQuery.Row().Scan(&id)

	util.WriteLogMain(fmt.Printf("id : %v", id))
	if err!=nil{
		tx.Rollback()
		return nil,util.ErrorHandler(ctx, child, err)
	}
	tx.Commit()
	util.WriteLogMain(fmt.Sprintf("find last id : %v",(int32)(id)))
	cityResult ,err := impl.FindById(ctx,(int32)(id))
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return cityResult,nil
}

func (impl CityGormPostgreImpl)Update(ctx context.Context,city dto.City)(*dto.City , error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance := helper.GetDBInstance()
	cityName := ""
	if strings.Contains(city.CityName, "'"){
		for _, data := range strings.Split(city.CityName, ""){
			if data =="'"{
				cityName += "'"
				cityName += data
			}else{
				cityName += data
			}
		}
		city.CityName = cityName
	}
	_, err := impl.DetailProvince(ctx, dto.Province{ID: city.ProvinceID})
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, util.UnhandledError{ErrorMessage:fmt.Sprintf("province with id %v not exist", city.ProvinceID)})
	}
	dbInstance.LogMode(true)
	tx := dbInstance.Begin()
	query := fmt.Sprintf("update %v set city_name='%v', province_id='%v', updated_at='%v' where city_id=%v",constant.CITY_TABLE_NAME, city.CityName, city.ProvinceID,time.Now().Format(time.RFC3339),city.CityID)
	util.WriteLogMain(fmt.Sprintf("start execute : %v",query))
	err = dbInstance.Exec(query).Error
	if err!=nil{
		tx.Rollback()
		return nil,util.ErrorHandler(ctx, child, err)
	}
	tx.Commit()
	util.WriteLogMain(fmt.Sprintf("find last id : %v",city.CityID))
	cityResult ,err := impl.FindById(ctx,city.CityID)
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return cityResult,nil
}

func (impl CityGormPostgreImpl)Delete(ctx context.Context,id int32)(*bool , error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance := helper.GetDBInstance()
	queryCek := fmt.Sprintf("SELECT count(1) FROM %v where city_id = '%v'", constant.DISTRICT_TABLE_NAME, id)
	util.WriteLogMain(fmt.Sprintf("start execute : %v",queryCek))
	countCheck := 0
	_ = dbInstance.Raw(queryCek).Row().Scan(&countCheck)
	if countCheck > 0{
		return nil, util.UnhandledError{ErrorMessage:"City Id Is Still Used By The District!"}
	}
	dbInstance.LogMode(true)
	tx := dbInstance.Begin()
	query := fmt.Sprintf("delete FROM %v where city_id = %v", constant.CITY_TABLE_NAME, id)
	util.WriteLogMain(fmt.Sprintf("start execute : %v",query))
	delete := dbInstance.Exec(query)
	if delete.Error != nil {
		tx.Rollback()
		util.WriteLogMain(delete.Error)
		return nil,util.ErrorHandler(ctx, child, delete.Error)
	}
	util.WriteLogMain(fmt.Printf("effected : %v", delete.RowsAffected))
	if delete.RowsAffected<=0{
		return nil, util.ErrorHandler(ctx, child, util.UnhandledError{ErrorMessage:fmt.Sprintf("failed to delete City With id %v", id)})
	}
	tx.Commit()
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	result := true
	return &result,nil
}

func (impl CityGormPostgreImpl) GetAllDistrict(ctx context.Context, district dto.FilterPaging) (*dto.ListDistrict, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance := helper.GetDBInstance()
	dbInstance.LogMode(true)
	resp := dto.ListDistrict{}
	paging := dto.Paging{}
	offset := (district.Paging.Page * district.Paging.PerPage) - district.Paging.PerPage
	query := fmt.Sprintf("SELECT a.district_id, a.district_name, a.city_id, b.city_name FROM %v a LEFT JOIN %v b on a.city_id = b.city_id where 1 = 1 ", constant.DISTRICT_TABLE_NAME, constant.CITY_TABLE_NAME)
	condition := ""
	if district.Filter.CityID!=0{
		condition = condition+fmt.Sprintf(" And a.city_id=%v",district.Filter.CityID)
	}
	if district.Filter.Keyword!=""{
		districtName := ""
		if strings.Contains(district.Filter.Keyword, "'"){
			for _, data := range strings.Split(district.Filter.Keyword, ""){
				if data =="'"{
					districtName += "'"
					districtName += data
				}else{
					districtName += data
				}
			}
			district.Filter.Keyword = districtName
		}
		key := "%"+strings.ToLower(district.Filter.Keyword)+"%"
		condition += fmt.Sprintf(" AND LOWER(a.district_name) like '%v' ",key)
	}
	newQuery := dbInstance.Raw(query+condition)
	if district.Paging.Page != 0 || district.Paging.PerPage != 0 {
		newQuery = newQuery.Order("a.district_id").Offset(offset).Limit(district.Paging.PerPage)
	}
	util.WriteLogMain(fmt.Sprintf("start execute : %v",query))
	rows, err := newQuery.Rows()
	if err != nil {
		util.WriteLogMain(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	defer rows.Close()

	districts := make([]dto.District, 0)
	for rows.Next() {
		districtModel := new(dto.District)
		err := rows.Scan(&districtModel.DistrictID, &districtModel.DistrictName, &districtModel.CityID, &districtModel.CityName)
		if err != nil {
			util.WriteLogMain(err)
		}
		districts = append(districts, *districtModel)
	}
	if err = rows.Err(); err != nil {
		util.WriteLogMain(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	query = fmt.Sprintf("select count(1) from %v a where 1 = 1 ", constant.DISTRICT_TABLE_NAME)
	row := dbInstance.Raw(query+condition).Row()
	count := 0
	err = row.Scan(&count)
	if err != nil {
		util.WriteLogMain(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	resp.District = districts
	paging.Page = district.Paging.Page
	paging.PerPage = district.Paging.PerPage
	paging.Count =int32(count)
	resp.Paging = paging
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return &resp,nil
}

func (impl CityGormPostgreImpl)DistrictFindById(ctx context.Context,id int32) (*dto.District, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance := helper.GetDBInstance()
	dbInstance.LogMode(true)
	query := fmt.Sprintf("SELECT a.district_id, a.district_name, a.city_id, b.city_name FROM %v a LEFT JOIN %v b on a.city_id = b.city_id where district_id = %v", constant.DISTRICT_TABLE_NAME, constant.CITY_TABLE_NAME, id)
	util.WriteLogMain(fmt.Sprintf("start execute : %v",query))
	row := dbInstance.Raw(query).Row()
	district := new(dto.District)
	err := row.Scan(&district.DistrictID, &district.DistrictName, &district.CityID, &district.CityName)
	if err != nil {
		util.WriteLogMain(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return district,nil
}

func (impl CityGormPostgreImpl)CreateDistrict(ctx context.Context,district dto.District)(*dto.District , error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance := helper.GetDBInstance()
	_, err := impl.FindById(ctx, district.CityID)
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, util.UnhandledError{ErrorMessage:fmt.Sprintf("city with id %v not exist", district.CityID)})
	}
	districtName := ""
	if strings.Contains(district.DistrictName, "'"){
		for _, data := range strings.Split(district.DistrictName, ""){
			if data =="'"{
				districtName += "'"
				districtName += data
			}else{
				districtName += data
			}
		}
		district.DistrictName = districtName
	}
	dbInstance.LogMode(true)
	tx := dbInstance.Begin()
	var id int
	query := fmt.Sprintf("INSERT INTO %v(district_name, city_id, created_at) VALUES ('%v', '%v', '%v')",constant.DISTRICT_TABLE_NAME, district.DistrictName, district.CityID, time.Now().Format(time.RFC3339))
	newQuery := helper.GetCurrentImplementation().GetInsertQueryCity(tx, query, "district_id", constant.DISTRICT_TABLE_NAME)
	err = newQuery.Row().Scan(&id)
	if err!=nil{
		tx.Rollback()
		return nil,util.ErrorHandler(ctx, child, err)
	}
	tx.Commit()
	util.WriteLogMain(fmt.Sprintf("find last id : %v",(int32)(id)))
	districtResult ,err := impl.DistrictFindById(ctx,(int32)(id))
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return districtResult,nil
}

func (impl CityGormPostgreImpl)UpdateDistrict(ctx context.Context,district dto.District)(*dto.District , error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance := helper.GetDBInstance()
	_, err := impl.FindById(ctx, district.CityID)
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, util.UnhandledError{ErrorMessage:fmt.Sprintf("city with id %v not exist", district.CityID)})
	}
	districtName := ""
	if strings.Contains(district.DistrictName, "'"){
		for _, data := range strings.Split(district.DistrictName, ""){
			if data =="'"{
				districtName += "'"
				districtName += data
			}else{
				districtName += data
			}
		}
		district.DistrictName = districtName
	}
	dbInstance.LogMode(true)
	tx := dbInstance.Begin()
	query := fmt.Sprintf("update %v set district_name='%v', city_id='%v', updated_at='%v' where district_id=%v",constant.DISTRICT_TABLE_NAME, district.DistrictName, district.CityID,time.Now().Format(time.RFC3339),district.DistrictID)
	util.WriteLogMain(fmt.Sprintf("start execute : %v",query))
	err = tx.Exec(query).Error
	if err!=nil{
		tx.Rollback()
		return nil,util.ErrorHandler(ctx, child, err)
	}
	tx.Commit()
	util.WriteLogMain(fmt.Sprintf("find last id : %v",district.DistrictID))
	districtResult ,err := impl.DistrictFindById(ctx,district.DistrictID)
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return districtResult,nil
}

func (impl CityGormPostgreImpl)DeleteDistrict(ctx context.Context,id int32)(*bool , error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance := helper.GetDBInstance()
	dbInstance.LogMode(true)
	tx := dbInstance.Begin()
	query := fmt.Sprintf("delete FROM %v where district_id = %v",constant.DISTRICT_TABLE_NAME, id)
	util.WriteLogMain(fmt.Sprintf("start execute : %v",query))
	delete := tx.Exec(query)
	if delete.Error != nil {
		tx.Rollback()
		util.WriteLogMain(delete.Error)
		return nil,util.ErrorHandler(ctx, child, delete.Error)
	}
	util.WriteLogMain(fmt.Printf("effected : %v", delete.RowsAffected))
	if delete.RowsAffected<=0{
		return nil, util.ErrorHandler(ctx, child, util.UnhandledError{ErrorMessage:fmt.Sprintf("failed to delete District With id %v", id)})
	}
	tx.Commit()
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	result := true
	return &result,nil
}

func (impl CityGormPostgreImpl) ListProvince(ctx context.Context, city dto.FilterPaging) ([]*dto.Province,*int32, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance := helper.GetDBInstance()
	dbInstance.LogMode(true)
	offset := (city.Paging.Page * city.Paging.PerPage) - city.Paging.PerPage
	query := fmt.Sprintf("SELECT province_id, province_name FROM %v where 1 = 1 ", constant.PROVINCE_TABLE_NAME)
	condition := ""
	if city.Filter.ProvinceID!=0{
		condition = condition+fmt.Sprintf(" And province_id=%v",city.Filter.ProvinceID)
	}
	if city.Filter.Keyword!=""{
		provinceName := ""
		if strings.Contains(city.Filter.Keyword, "'"){
			for _, data := range strings.Split(city.Filter.Keyword, ""){
				if data =="'"{
					provinceName += "'"
					provinceName += data
				}else{
					provinceName += data
				}
			}
			city.Filter.Keyword = provinceName
		}
		key := "%"+strings.ToLower(city.Filter.Keyword)+"%"
		condition += fmt.Sprintf(" AND LOWER(province_name) like '%v' ",key)
	}
	newQuery := dbInstance.Raw(query+condition)
	if city.Paging.Page != 0 || city.Paging.PerPage != 0 {
		newQuery = newQuery.Order("province_id").Offset(offset).Limit(city.Paging.PerPage)
	}
	util.WriteLogMain(fmt.Sprintf("start execute : %v",query))
	rows, err := newQuery.Rows()
	if err != nil {
		util.WriteLogMain(err)
		return nil,nil,util.ErrorHandler(ctx, child, err)
	}
	defer rows.Close()

	provs := make([]*dto.Province, 0)
	for rows.Next() {
		prov := new(dto.Province)
		err := rows.Scan(&prov.ID, &prov.ProvinceName)
		if err != nil {
			util.WriteLogMain(err)
		}
		provs = append(provs, prov)
	}
	if err = rows.Err(); err != nil {
		util.WriteLogMain(err)
		return nil,nil,util.ErrorHandler(ctx, child, err)
	}
	query = fmt.Sprintf("select count(1) from %v where 1 = 1 ", constant.PROVINCE_TABLE_NAME)
	row := dbInstance.Raw(query+condition).Row()
	count := 0
	err = row.Scan(&count)
	if err != nil {
		util.WriteLogMain(err)
		return nil,nil,util.ErrorHandler(ctx, child, err)
	}
	countInt32 :=  int32(count)
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return provs,&countInt32,nil
}

func (impl CityGormPostgreImpl)DetailProvince(ctx context.Context,province dto.Province) (*dto.Province, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance := helper.GetDBInstance()
	dbInstance.LogMode(true)
	query := fmt.Sprintf("SELECT province_id, province_name FROM %v where province_id = %v",constant.PROVINCE_TABLE_NAME, province.ID)
	util.WriteLogMain(fmt.Sprintf("start execute : %v",query))
	row := dbInstance.Raw(query).Row()
	prov := new(dto.Province)
	err := row.Scan(&prov.ID, &prov.ProvinceName)
	if err != nil {
		util.WriteLogMain(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return prov,nil
}

func (impl CityGormPostgreImpl)AddProvince(ctx context.Context,province dto.Province) (*dto.Province, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance := helper.GetDBInstance()
	provinceName := ""
	if strings.Contains(province.ProvinceName, "'"){
		for _, data := range strings.Split(province.ProvinceName, ""){
			if data =="'"{
				provinceName += "'"
				provinceName += data
			}else{
				provinceName += data
			}
		}
		province.ProvinceName = provinceName
	}
	dbInstance.LogMode(true)
	tx := dbInstance.Begin()
	var id int
	query := fmt.Sprintf("INSERT INTO %v(province_name, created_at) VALUES ('%v', '%v')",constant.PROVINCE_TABLE_NAME, province.ProvinceName, time.Now().Format(time.RFC3339))
	newQuery := helper.GetCurrentImplementation().GetInsertQueryCity(tx, query, "province_id", constant.PROVINCE_TABLE_NAME)
	err := newQuery.Row().Scan(&id)
	if err!=nil{
		tx.Rollback()
		return nil,util.ErrorHandler(ctx, child, err)
	}
	tx.Commit()
	util.WriteLogMain(fmt.Sprintf("find last id : %v",(int32)(id)))
	cityResult ,err := impl.DetailProvince(ctx, dto.Province{ID: (int32)(id)})
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return cityResult,nil
}

func (impl CityGormPostgreImpl)EditProvince(ctx context.Context,province dto.Province) (*dto.Province, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance := helper.GetDBInstance()
	provinceName := ""
	if strings.Contains(province.ProvinceName, "'"){
		for _, data := range strings.Split(province.ProvinceName, ""){
			if data =="'"{
				provinceName += "'"
				provinceName += data
			}else{
				provinceName += data
			}
		}
		province.ProvinceName = provinceName
	}
	dbInstance.LogMode(true)
	tx := dbInstance.Begin()
	query := fmt.Sprintf("update %v set province_name='%v', updated_at='%v' where province_id=%v",constant.PROVINCE_TABLE_NAME, province.ProvinceName,time.Now().Format(time.RFC3339),province.ID)
	util.WriteLogMain(fmt.Sprintf("start execute : %v",query))
	err := tx.Exec(query).Error
	if err!=nil{
		tx.Rollback()
		return nil,util.ErrorHandler(ctx, child, err)
	}
	tx.Commit()
	util.WriteLogMain(fmt.Sprintf("find last id : %v",province.ID))
	cityResult ,err := impl.DetailProvince(ctx, dto.Province{ID: province.ID})
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return cityResult,nil
}

func (impl CityGormPostgreImpl)RemoveProvince(ctx context.Context,province dto.Province) error{
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance := helper.GetDBInstance()
	queryCek := fmt.Sprintf("SELECT count(1) FROM %v where province_id = '%v'", constant.CITY_TABLE_NAME, province.ID)
	util.WriteLogMain(fmt.Sprintf("start execute : %v",queryCek))
	countCheck := 0
	_ = dbInstance.Raw(queryCek).Row().Scan(&countCheck)
	if countCheck > 0{
		return util.UnhandledError{ErrorMessage:"Province Id Is Still Used By Some City!"}
	}
	dbInstance.LogMode(true)
	tx := dbInstance.Begin()
	query := fmt.Sprintf("delete FROM %v where province_id = %v",constant.PROVINCE_TABLE_NAME, province.ID)
	util.WriteLogMain(fmt.Sprintf("start execute : %v",query))
	delete := tx.Exec(query)
	if delete.Error != nil {
		tx.Rollback()
		util.WriteLogMain(delete.Error)
		return util.ErrorHandler(ctx, child, delete.Error)
	}
	util.WriteLogMain(fmt.Printf("effected : %v", delete.RowsAffected))
	if delete.RowsAffected<=0{
		return util.ErrorHandler(ctx, child, util.UnhandledError{ErrorMessage:fmt.Sprintf("failed to delete Province With id %v", province.ID)})
	}
	tx.Commit()
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return nil
}