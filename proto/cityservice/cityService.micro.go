// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/cityservice/cityService.proto

/*
Package go_micro_srv_CityService is a generated protocol buffer package.

It is generated from these files:
	proto/cityservice/cityService.proto

It has these top-level messages:
	ProvinceData
	ListProvinceData
	EmptyRequest
	Response
	ListCityResponse
	ListDistrictResponse
	CityData
	DistrictData
	FilterPaging
	IsRemoved
	CityKey
	FilterKey
	Paging
	RoleData
	PermissionData
*/
package go_micro_srv_CityService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for City service

type CityService interface {
	Healthcheck(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*Response, error)
	List(ctx context.Context, in *FilterPaging, opts ...client.CallOption) (*ListCityResponse, error)
	Detail(ctx context.Context, in *CityKey, opts ...client.CallOption) (*CityData, error)
	Add(ctx context.Context, in *CityData, opts ...client.CallOption) (*CityData, error)
	Edit(ctx context.Context, in *CityData, opts ...client.CallOption) (*CityData, error)
	Remove(ctx context.Context, in *CityKey, opts ...client.CallOption) (*IsRemoved, error)
	ListDistrict(ctx context.Context, in *FilterPaging, opts ...client.CallOption) (*ListDistrictResponse, error)
	DetailDistrict(ctx context.Context, in *CityKey, opts ...client.CallOption) (*DistrictData, error)
	AddDistrict(ctx context.Context, in *DistrictData, opts ...client.CallOption) (*DistrictData, error)
	EditDistrict(ctx context.Context, in *DistrictData, opts ...client.CallOption) (*DistrictData, error)
	RemoveDistrict(ctx context.Context, in *CityKey, opts ...client.CallOption) (*IsRemoved, error)
	AddProvince(ctx context.Context, in *ProvinceData, opts ...client.CallOption) (*ProvinceData, error)
	EditProvince(ctx context.Context, in *ProvinceData, opts ...client.CallOption) (*ProvinceData, error)
	RemoveProvince(ctx context.Context, in *ProvinceData, opts ...client.CallOption) (*Response, error)
	DetailProvince(ctx context.Context, in *ProvinceData, opts ...client.CallOption) (*ProvinceData, error)
	ListProvince(ctx context.Context, in *FilterPaging, opts ...client.CallOption) (*ListProvinceData, error)
}

type cityService struct {
	c    client.Client
	name string
}

func NewCityService(name string, c client.Client) CityService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.CityService"
	}
	return &cityService{
		c:    c,
		name: name,
	}
}

func (c *cityService) Healthcheck(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "City.Healthcheck", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityService) List(ctx context.Context, in *FilterPaging, opts ...client.CallOption) (*ListCityResponse, error) {
	req := c.c.NewRequest(c.name, "City.List", in)
	out := new(ListCityResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityService) Detail(ctx context.Context, in *CityKey, opts ...client.CallOption) (*CityData, error) {
	req := c.c.NewRequest(c.name, "City.Detail", in)
	out := new(CityData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityService) Add(ctx context.Context, in *CityData, opts ...client.CallOption) (*CityData, error) {
	req := c.c.NewRequest(c.name, "City.Add", in)
	out := new(CityData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityService) Edit(ctx context.Context, in *CityData, opts ...client.CallOption) (*CityData, error) {
	req := c.c.NewRequest(c.name, "City.Edit", in)
	out := new(CityData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityService) Remove(ctx context.Context, in *CityKey, opts ...client.CallOption) (*IsRemoved, error) {
	req := c.c.NewRequest(c.name, "City.Remove", in)
	out := new(IsRemoved)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityService) ListDistrict(ctx context.Context, in *FilterPaging, opts ...client.CallOption) (*ListDistrictResponse, error) {
	req := c.c.NewRequest(c.name, "City.ListDistrict", in)
	out := new(ListDistrictResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityService) DetailDistrict(ctx context.Context, in *CityKey, opts ...client.CallOption) (*DistrictData, error) {
	req := c.c.NewRequest(c.name, "City.DetailDistrict", in)
	out := new(DistrictData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityService) AddDistrict(ctx context.Context, in *DistrictData, opts ...client.CallOption) (*DistrictData, error) {
	req := c.c.NewRequest(c.name, "City.AddDistrict", in)
	out := new(DistrictData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityService) EditDistrict(ctx context.Context, in *DistrictData, opts ...client.CallOption) (*DistrictData, error) {
	req := c.c.NewRequest(c.name, "City.EditDistrict", in)
	out := new(DistrictData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityService) RemoveDistrict(ctx context.Context, in *CityKey, opts ...client.CallOption) (*IsRemoved, error) {
	req := c.c.NewRequest(c.name, "City.RemoveDistrict", in)
	out := new(IsRemoved)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityService) AddProvince(ctx context.Context, in *ProvinceData, opts ...client.CallOption) (*ProvinceData, error) {
	req := c.c.NewRequest(c.name, "City.AddProvince", in)
	out := new(ProvinceData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityService) EditProvince(ctx context.Context, in *ProvinceData, opts ...client.CallOption) (*ProvinceData, error) {
	req := c.c.NewRequest(c.name, "City.EditProvince", in)
	out := new(ProvinceData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityService) RemoveProvince(ctx context.Context, in *ProvinceData, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "City.RemoveProvince", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityService) DetailProvince(ctx context.Context, in *ProvinceData, opts ...client.CallOption) (*ProvinceData, error) {
	req := c.c.NewRequest(c.name, "City.DetailProvince", in)
	out := new(ProvinceData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityService) ListProvince(ctx context.Context, in *FilterPaging, opts ...client.CallOption) (*ListProvinceData, error) {
	req := c.c.NewRequest(c.name, "City.ListProvince", in)
	out := new(ListProvinceData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for City service

type CityHandler interface {
	Healthcheck(context.Context, *EmptyRequest, *Response) error
	List(context.Context, *FilterPaging, *ListCityResponse) error
	Detail(context.Context, *CityKey, *CityData) error
	Add(context.Context, *CityData, *CityData) error
	Edit(context.Context, *CityData, *CityData) error
	Remove(context.Context, *CityKey, *IsRemoved) error
	ListDistrict(context.Context, *FilterPaging, *ListDistrictResponse) error
	DetailDistrict(context.Context, *CityKey, *DistrictData) error
	AddDistrict(context.Context, *DistrictData, *DistrictData) error
	EditDistrict(context.Context, *DistrictData, *DistrictData) error
	RemoveDistrict(context.Context, *CityKey, *IsRemoved) error
	AddProvince(context.Context, *ProvinceData, *ProvinceData) error
	EditProvince(context.Context, *ProvinceData, *ProvinceData) error
	RemoveProvince(context.Context, *ProvinceData, *Response) error
	DetailProvince(context.Context, *ProvinceData, *ProvinceData) error
	ListProvince(context.Context, *FilterPaging, *ListProvinceData) error
}

func RegisterCityHandler(s server.Server, hdlr CityHandler, opts ...server.HandlerOption) error {
	type city interface {
		Healthcheck(ctx context.Context, in *EmptyRequest, out *Response) error
		List(ctx context.Context, in *FilterPaging, out *ListCityResponse) error
		Detail(ctx context.Context, in *CityKey, out *CityData) error
		Add(ctx context.Context, in *CityData, out *CityData) error
		Edit(ctx context.Context, in *CityData, out *CityData) error
		Remove(ctx context.Context, in *CityKey, out *IsRemoved) error
		ListDistrict(ctx context.Context, in *FilterPaging, out *ListDistrictResponse) error
		DetailDistrict(ctx context.Context, in *CityKey, out *DistrictData) error
		AddDistrict(ctx context.Context, in *DistrictData, out *DistrictData) error
		EditDistrict(ctx context.Context, in *DistrictData, out *DistrictData) error
		RemoveDistrict(ctx context.Context, in *CityKey, out *IsRemoved) error
		AddProvince(ctx context.Context, in *ProvinceData, out *ProvinceData) error
		EditProvince(ctx context.Context, in *ProvinceData, out *ProvinceData) error
		RemoveProvince(ctx context.Context, in *ProvinceData, out *Response) error
		DetailProvince(ctx context.Context, in *ProvinceData, out *ProvinceData) error
		ListProvince(ctx context.Context, in *FilterPaging, out *ListProvinceData) error
	}
	type City struct {
		city
	}
	h := &cityHandler{hdlr}
	return s.Handle(s.NewHandler(&City{h}, opts...))
}

type cityHandler struct {
	CityHandler
}

func (h *cityHandler) Healthcheck(ctx context.Context, in *EmptyRequest, out *Response) error {
	return h.CityHandler.Healthcheck(ctx, in, out)
}

func (h *cityHandler) List(ctx context.Context, in *FilterPaging, out *ListCityResponse) error {
	return h.CityHandler.List(ctx, in, out)
}

func (h *cityHandler) Detail(ctx context.Context, in *CityKey, out *CityData) error {
	return h.CityHandler.Detail(ctx, in, out)
}

func (h *cityHandler) Add(ctx context.Context, in *CityData, out *CityData) error {
	return h.CityHandler.Add(ctx, in, out)
}

func (h *cityHandler) Edit(ctx context.Context, in *CityData, out *CityData) error {
	return h.CityHandler.Edit(ctx, in, out)
}

func (h *cityHandler) Remove(ctx context.Context, in *CityKey, out *IsRemoved) error {
	return h.CityHandler.Remove(ctx, in, out)
}

func (h *cityHandler) ListDistrict(ctx context.Context, in *FilterPaging, out *ListDistrictResponse) error {
	return h.CityHandler.ListDistrict(ctx, in, out)
}

func (h *cityHandler) DetailDistrict(ctx context.Context, in *CityKey, out *DistrictData) error {
	return h.CityHandler.DetailDistrict(ctx, in, out)
}

func (h *cityHandler) AddDistrict(ctx context.Context, in *DistrictData, out *DistrictData) error {
	return h.CityHandler.AddDistrict(ctx, in, out)
}

func (h *cityHandler) EditDistrict(ctx context.Context, in *DistrictData, out *DistrictData) error {
	return h.CityHandler.EditDistrict(ctx, in, out)
}

func (h *cityHandler) RemoveDistrict(ctx context.Context, in *CityKey, out *IsRemoved) error {
	return h.CityHandler.RemoveDistrict(ctx, in, out)
}

func (h *cityHandler) AddProvince(ctx context.Context, in *ProvinceData, out *ProvinceData) error {
	return h.CityHandler.AddProvince(ctx, in, out)
}

func (h *cityHandler) EditProvince(ctx context.Context, in *ProvinceData, out *ProvinceData) error {
	return h.CityHandler.EditProvince(ctx, in, out)
}

func (h *cityHandler) RemoveProvince(ctx context.Context, in *ProvinceData, out *Response) error {
	return h.CityHandler.RemoveProvince(ctx, in, out)
}

func (h *cityHandler) DetailProvince(ctx context.Context, in *ProvinceData, out *ProvinceData) error {
	return h.CityHandler.DetailProvince(ctx, in, out)
}

func (h *cityHandler) ListProvince(ctx context.Context, in *FilterPaging, out *ListProvinceData) error {
	return h.CityHandler.ListProvince(ctx, in, out)
}