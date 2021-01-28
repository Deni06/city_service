package city

import (
	"context"
	"fmt"
	"github.com/micro/go-log"
	"github.com/opentracing/opentracing-go"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/constant"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/driver"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/dto"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/util"
	"strings"
	"time"
)

type CityPostgreImpl struct {}

func (impl CityPostgreImpl) GetAll(ctx context.Context, city dto.FilterPaging) (*dto.ListCities, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance,errDBInstance := driver.GetDBInstance()
	if(errDBInstance!=nil){
		log.Log(errDBInstance)
		return nil,util.ErrorHandler(ctx, child, errDBInstance)
	}

	resp := dto.ListCities{}
	paging := dto.Paging{}
	offset := (city.Paging.Page * city.Paging.PerPage) - city.Paging.PerPage
	query := "SELECT c.city_id, c.city_name,c.province_id, coalesce(p.province_name,'') FROM city c left join province p on p.province_id = c.province_id where true "
	if city.Filter.ProvinceID!=0{
		query = query+fmt.Sprintf(" And c.province_id=%v",city.Filter.ProvinceID)
	}
	if city.Filter.Keyword!=""{
		key := "%"+strings.ToLower(city.Filter.Keyword)+"%"
		query += fmt.Sprintf(" AND LOWER(c.city_name) like '%v' ",key)
	}
	if city.Paging.Page != 0 || city.Paging.PerPage != 0 {
		query += fmt.Sprintf(" offset %v limit %v ", offset, city.Paging.PerPage)
	}
	log.Log(fmt.Sprintf("start execute : %v",query))
	rows, err := dbInstance.Query(query)
	if err != nil {
		log.Log(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	defer rows.Close()

	cities := make([]dto.City, 0)
	for rows.Next() {
		citymodel := new(dto.City)
		err := rows.Scan(&citymodel.CityID, &citymodel.CityName, &citymodel.ProvinceID, &citymodel.ProvinceName)
		if err != nil {
			log.Log(err)
		}
		cities = append(cities, *citymodel)
	}
	if err = rows.Err(); err != nil {
		log.Log(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	query = fmt.Sprintf("select count(1) from city where true ")
	if city.Filter.ProvinceID!=0{
		query = query+fmt.Sprintf(" And province_id=%v",city.Filter.ProvinceID)
	}
	if city.Filter.Keyword!=""{
		key := "%"+strings.ToLower(city.Filter.Keyword)+"%"
		query += fmt.Sprintf(" AND LOWER(city_name) like '%v' ",key)
	}
	row := dbInstance.QueryRow(query)
	count := 0
	err = row.Scan(&count)
	if err != nil {
		log.Log(err)
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

func (impl CityPostgreImpl)FindById(ctx context.Context,id int32) (*dto.City, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance,errDBInstance := driver.GetDBInstance()
	if(errDBInstance!=nil){
		log.Log(errDBInstance)
		return nil,util.ErrorHandler(ctx, child, errDBInstance)
	}
	query := fmt.Sprintf("SELECT c.city_id, c.city_name, c.province_id, coalesce(p.province_name,'') FROM city  c  left join province p on p.province_id = c.province_id where c.city_id = %v",id)
	log.Log(fmt.Sprintf("start execute : %v",query))
	row := dbInstance.QueryRow(query)
	city := new(dto.City)
	err := row.Scan(&city.CityID, &city.CityName, &city.ProvinceID,&city.ProvinceName)
	if err != nil {
		log.Log(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return city,nil
}

func (impl CityPostgreImpl)Create(ctx context.Context,city dto.City)(*dto.City , error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance,errDBInstance := driver.GetDBInstance()
	if(errDBInstance!=nil){
		log.Log(errDBInstance)
		return nil,util.ErrorHandler(ctx, child, errDBInstance)
	}
	_, err := impl.DetailProvince(ctx, dto.Province{ID: city.ProvinceID})
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, util.UnhandledError{ErrorMessage:fmt.Sprintf("province with id %v not exist", city.ProvinceID)})
	}
	query := fmt.Sprintf("INSERT INTO city(city_name, province_id, created_at) VALUES ('%v', '%v', '%v') RETURNING city_id",city.CityName, city.ProvinceID, time.Now().Format(time.RFC3339))
	log.Log(fmt.Sprintf("start execute : %v",query))
	var id int
	err = dbInstance.QueryRow(query).Scan(&id)
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	log.Log(fmt.Sprintf("find last id : %v",(int32)(id)))
	cityResult ,err := impl.FindById(ctx,(int32)(id))
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return cityResult,nil
}

func (impl CityPostgreImpl)Update(ctx context.Context,city dto.City)(*dto.City , error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance,errDBInstance := driver.GetDBInstance()
	if(errDBInstance!=nil){
		log.Log(errDBInstance)
		return nil,util.ErrorHandler(ctx, child, errDBInstance)
	}
	query := fmt.Sprintf("update city set city_name='%v', province_id='%v', updated_at='%v' where city_id=%v",city.CityName, city.ProvinceID,time.Now().Format(time.RFC3339),city.CityID)
	log.Log(fmt.Sprintf("start execute : %v",query))
	_,err := dbInstance.Exec(query)
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	log.Log(fmt.Sprintf("find last id : %v",city.CityID))
	cityResult ,err := impl.FindById(ctx,city.CityID)
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return cityResult,nil
}

func (impl CityPostgreImpl)Delete(ctx context.Context,id int32)(*bool , error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance,errDBInstance := driver.GetDBInstance()
	if(errDBInstance!=nil){
		log.Log(errDBInstance)
		return nil,util.ErrorHandler(ctx, child, errDBInstance)
	}
	querycek := fmt.Sprintf("SELECT count(1) FROM district where city_id = '%v'", id)
	log.Log(fmt.Sprintf("start execute : %v",querycek))
	countcheck := 0
	_ = dbInstance.QueryRow(querycek).Scan(&countcheck)
	if countcheck > 0{
		return nil, util.UnhandledError{ErrorMessage:"City Id Is Still Used By The District!"}
	}
	query := fmt.Sprintf("delete FROM city where city_id = %v",id)
	log.Log(fmt.Sprintf("start execute : %v",query))
	res,err:= dbInstance.Exec(query)
	if err != nil {
		log.Log(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	effected, err := res.RowsAffected()
	if err!=nil{
		return nil, util.ErrorHandler(ctx, child, err)
	}
	if effected<=0{
		return nil, util.ErrorHandler(ctx, child, util.UnhandledError{ErrorMessage:fmt.Sprintf("failed to delete city with id %v", id)})
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	result := true
	return &result,nil
}

func (impl CityPostgreImpl) GetAllDistrict(ctx context.Context, district dto.FilterPaging) (*dto.ListDistrict, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance,errDBInstance := driver.GetDBInstance()
	if(errDBInstance!=nil){
		log.Log(errDBInstance)
		return nil,util.ErrorHandler(ctx, child, errDBInstance)
	}
	resp := dto.ListDistrict{}
	paging := dto.Paging{}
	offset := (district.Paging.Page * district.Paging.PerPage) - district.Paging.PerPage
	query := "SELECT a.district_id, a.district_name, a.city_id, b.city_name FROM district a LEFT JOIN city b on a.city_id = b.city_id where true "
	if district.Filter.ProvinceID!=0{
		query = query+fmt.Sprintf(" And a.city_id=%v",district.Filter.CityID)
	}
	if district.Filter.Keyword!=""{
		key := "%"+strings.ToLower(district.Filter.Keyword)+"%"
		query += fmt.Sprintf(" AND LOWER(a.district_name) like '%v' ",key)
	}
	if district.Paging.Page != 0 || district.Paging.PerPage != 0 {
		query += fmt.Sprintf(" offset %v limit %v ", offset, district.Paging.PerPage)
	}
	log.Log(fmt.Sprintf("start execute : %v",query))
	rows, err := dbInstance.Query(query)
	if err != nil {
		log.Log(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	defer rows.Close()

	districts := make([]dto.District, 0)
	for rows.Next() {
		districtmodel := new(dto.District)
		err := rows.Scan(&districtmodel.DistrictID, &districtmodel.DistrictName, &districtmodel.CityID, &districtmodel.CityName)
		if err != nil {
			log.Log(err)
		}
		districts = append(districts, *districtmodel)
	}
	if err = rows.Err(); err != nil {
		log.Log(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	query = fmt.Sprintf("select count(1) from district where true ")
	if district.Filter.ProvinceID!=0{
		query = query+fmt.Sprintf(" And district_id=%v",district.Filter.ProvinceID)
	}
	if district.Filter.Keyword!=""{
		key := "%"+strings.ToLower(district.Filter.Keyword)+"%"
		query += fmt.Sprintf(" AND LOWER(district_name) like '%v' ",key)
	}
	row := dbInstance.QueryRow(query)
	count := 0
	err = row.Scan(&count)
	if err != nil {
		log.Log(err)
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

func (impl CityPostgreImpl)DistrictFindById(ctx context.Context,id int32) (*dto.District, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance,errDBInstance := driver.GetDBInstance()
	if(errDBInstance!=nil){
		log.Log(errDBInstance)
		return nil,util.ErrorHandler(ctx, child, errDBInstance)
	}
	query := fmt.Sprintf("SELECT a.district_id, a.district_name, a.city_id, b.city_name FROM district a LEFT JOIN city b on a.city_id = b.city_id where district_id = %v",id)
	log.Log(fmt.Sprintf("start execute : %v",query))
	row := dbInstance.QueryRow(query)
	district := new(dto.District)
	err := row.Scan(&district.DistrictID, &district.DistrictName, &district.CityID, &district.CityName)
	if err != nil {
		log.Log(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return district,nil
}

func (impl CityPostgreImpl)CreateDistrict(ctx context.Context,district dto.District)(*dto.District , error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance,errDBInstance := driver.GetDBInstance()
	if(errDBInstance!=nil){
		log.Log(errDBInstance)
		return nil,util.ErrorHandler(ctx, child, errDBInstance)
	}
	_, err := impl.FindById(ctx, district.CityID)
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, util.UnhandledError{ErrorMessage:fmt.Sprintf("city with id %v not exist", district.CityID)})
	}
	query := fmt.Sprintf("INSERT INTO district(district_name, city_id, created_at) VALUES ('%v', '%v', '%v') RETURNING district_id",district.DistrictName, district.CityID, time.Now().Format(time.RFC3339))
	log.Log(fmt.Sprintf("start execute : %v",query))
	var id int
	err = dbInstance.QueryRow(query).Scan(&id)
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	log.Log(fmt.Sprintf("find last id : %v",(int32)(id)))
	districtResult ,err := impl.DistrictFindById(ctx,(int32)(id))
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return districtResult,nil
}

func (impl CityPostgreImpl)UpdateDistrict(ctx context.Context,district dto.District)(*dto.District , error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance,errDBInstance := driver.GetDBInstance()
	if(errDBInstance!=nil){
		log.Log(errDBInstance)
		return nil,util.ErrorHandler(ctx, child, errDBInstance)
	}
	_, err := impl.FindById(ctx, district.CityID)
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, util.UnhandledError{ErrorMessage:fmt.Sprintf("city with id %v not exist", district.CityID)})
	}
	query := fmt.Sprintf("update district set district_name='%v', city_id='%v', updated_at='%v' where district_id=%v",district.DistrictName, district.CityID,time.Now().Format(time.RFC3339),district.DistrictID)
	log.Log(fmt.Sprintf("start execute : %v",query))
	_,err = dbInstance.Exec(query)
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	log.Log(fmt.Sprintf("find last id : %v",district.DistrictID))
	districtResult ,err := impl.DistrictFindById(ctx,district.DistrictID)
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return districtResult,nil
}

func (impl CityPostgreImpl)DeleteDistrict(ctx context.Context,id int32)(*bool , error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance,errDBInstance := driver.GetDBInstance()
	if(errDBInstance!=nil){
		log.Log(errDBInstance)
		return nil,util.ErrorHandler(ctx, child, errDBInstance)
	}
	query := fmt.Sprintf("delete FROM district where district_id = %v",id)
	log.Log(fmt.Sprintf("start execute : %v",query))
	res,err:= dbInstance.Exec(query)
	if err != nil {
		log.Log(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	effected, err := res.RowsAffected()
	if err!=nil{
		return nil, util.ErrorHandler(ctx, child, err)
	}
	if effected<=0{
		return nil, util.ErrorHandler(ctx, child, util.UnhandledError{ErrorMessage:fmt.Sprintf("failed to delete district with id %v", id)})
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	result := true
	return &result,nil
}

func (impl CityPostgreImpl) ListProvince(ctx context.Context, city dto.FilterPaging) ([]*dto.Province,*int32, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance,errDBInstance := driver.GetDBInstance()
	if(errDBInstance!=nil){
		log.Log(errDBInstance)
		return nil,nil,util.ErrorHandler(ctx, child, errDBInstance)
	}
	offset := (city.Paging.Page * city.Paging.PerPage) - city.Paging.PerPage
	query := "SELECT province_id, province_name FROM province where true "
	if city.Filter.ProvinceID!=0{
		query = query+fmt.Sprintf(" And province_id=%v",city.Filter.ProvinceID)
	}
	if city.Filter.Keyword!=""{
		key := "%"+strings.ToLower(city.Filter.Keyword)+"%"
		query += fmt.Sprintf(" AND LOWER(province_name) like '%v' ",key)
	}
	if city.Paging.Page != 0 || city.Paging.PerPage != 0 {
		query += fmt.Sprintf(" offset %v limit %v ", offset, city.Paging.PerPage)
	}
	log.Log(fmt.Sprintf("start execute : %v",query))
	rows, err := dbInstance.Query(query)
	if err != nil {
		log.Log(err)
		return nil,nil,util.ErrorHandler(ctx, child, err)
	}
	defer rows.Close()

	provs := make([]*dto.Province, 0)
	for rows.Next() {
		prov := new(dto.Province)
		err := rows.Scan(&prov.ID, &prov.ProvinceName)
		if err != nil {
			log.Log(err)
		}
		provs = append(provs, prov)
	}
	if err = rows.Err(); err != nil {
		log.Log(err)
		return nil,nil,util.ErrorHandler(ctx, child, err)
	}
	query = fmt.Sprintf("select count(1) from province where true ")
	if city.Filter.ProvinceID!=0{
		query = query+fmt.Sprintf(" And province_id=%v",city.Filter.ProvinceID)
	}
	if city.Filter.Keyword!=""{
		key := "%"+strings.ToLower(city.Filter.Keyword)+"%"
		query += fmt.Sprintf(" AND LOWER(province_name) like '%v' ",key)
	}
	row := dbInstance.QueryRow(query)
	count := 0
	err = row.Scan(&count)
	if err != nil {
		log.Log(err)
		return nil,nil,util.ErrorHandler(ctx, child, err)
	}
	countInt32 :=  int32(count)
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return provs,&countInt32,nil
}

func (impl CityPostgreImpl)DetailProvince(ctx context.Context,province dto.Province) (*dto.Province, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance,errDBInstance := driver.GetDBInstance()
	if(errDBInstance!=nil){
		log.Log(errDBInstance)
		return nil,util.ErrorHandler(ctx, child, errDBInstance)
	}
	query := fmt.Sprintf("SELECT province_id, province_name FROM province where province_id = %v",province.ID)
	log.Log(fmt.Sprintf("start execute : %v",query))
	row := dbInstance.QueryRow(query)
	prov := new(dto.Province)
	err := row.Scan(&prov.ID, &prov.ProvinceName)
	if err != nil {
		log.Log(err)
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return prov,nil
}

func (impl CityPostgreImpl)AddProvince(ctx context.Context,province dto.Province) (*dto.Province, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance,errDBInstance := driver.GetDBInstance()
	if(errDBInstance!=nil){
		log.Log(errDBInstance)
		return nil,util.ErrorHandler(ctx, child, errDBInstance)
	}
	query := fmt.Sprintf("INSERT INTO province(province_name, created_at) VALUES ('%v', '%v') RETURNING province_id",province.ProvinceName, time.Now().Format(time.RFC3339))
	log.Log(fmt.Sprintf("start execute : %v",query))
	var id int
	err := dbInstance.QueryRow(query).Scan(&id)
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	log.Log(fmt.Sprintf("find last id : %v",(int32)(id)))
	cityResult ,err := impl.DetailProvince(ctx, dto.Province{ID: (int32)(id)})
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return cityResult,nil
}

func (impl CityPostgreImpl)EditProvince(ctx context.Context,province dto.Province) (*dto.Province, error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance,errDBInstance := driver.GetDBInstance()
	if(errDBInstance!=nil){
		log.Log(errDBInstance)
		return nil,util.ErrorHandler(ctx, child, errDBInstance)
	}
	query := fmt.Sprintf("update province set province_name='%v', updated_at='%v' where province_id=%v",province.ProvinceName,time.Now().Format(time.RFC3339),province.ID)
	log.Log(fmt.Sprintf("start execute : %v",query))
	_,err := dbInstance.Exec(query)
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	log.Log(fmt.Sprintf("find last id : %v",province.ID))
	cityResult ,err := impl.DetailProvince(ctx, dto.Province{ID: province.ID})
	if err!=nil{
		return nil,util.ErrorHandler(ctx, child, err)
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return cityResult,nil
}

func (impl CityPostgreImpl)RemoveProvince(ctx context.Context,province dto.Province) (error){
	child := initTraceAndLog(ctx)
	defer child.Finish()
	dbInstance,errDBInstance := driver.GetDBInstance()
	if(errDBInstance!=nil){
		log.Log(errDBInstance)
		return util.ErrorHandler(ctx, child, errDBInstance)
	}
	querycek := fmt.Sprintf("SELECT count(1) FROM city where province_id = '%v'", province.ID)
	log.Log(fmt.Sprintf("start execute : %v",querycek))
	countcheck := 0
	_ = dbInstance.QueryRow(querycek).Scan(&countcheck)
	if countcheck > 0{
		return util.UnhandledError{ErrorMessage:"Province Id Is Still Used By Some City!"}
	}
	query := fmt.Sprintf("delete FROM province where province_id = %v",province.ID)
	log.Log(fmt.Sprintf("start execute : %v",query))
	res,err:= dbInstance.Exec(query)
	if err != nil {
		log.Log(err)
		return util.ErrorHandler(ctx, child, err)
	}
	effected, err := res.RowsAffected()
	if err!=nil{
		return  util.ErrorHandler(ctx, child, err)
	}
	if effected<=0{
		return util.ErrorHandler(ctx, child, util.UnhandledError{ErrorMessage:fmt.Sprintf("failed to delete province with id %v", province.ID)})
	}
	child.SetTag(constant.TRACE_STATUS, constant.STATUS_SUCCESS)
	return nil
}


func initTraceAndLog(ctx context.Context)opentracing.Span{
	tracer := opentracing.GlobalTracer()
	span := opentracing.SpanFromContext(ctx)
	child := tracer.StartSpan(constant.CITY_POSTGRESQL_REPOSITORY, opentracing.ChildOf(span.Context()))
	return child
}

