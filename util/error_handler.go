package util

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/constant"
)

func ErrorHandler (ctx context.Context, span opentracing.Span, err error)error{
	span.SetTag(constant.TRACE_STATUS, constant.STATUS_ERROR)
	span.SetTag(constant.TRACE_ERROR_MESSAGE, err.Error())
	return err
}

type UnhandledError struct {
	ErrorMessage string
}

func (ce UnhandledError) Error()string{
	return ce.ErrorMessage
}
