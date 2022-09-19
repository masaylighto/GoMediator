package Mediator

import (
	"Mediator/Interfaces"
	"errors"
	"reflect"
)

type RequestType = Interfaces.IRequest[Interfaces.IRequestResult]
type RequestHandleType = Interfaces.IRequestHandler[RequestType, Interfaces.IRequestResult]
type Mediator struct {
	handlers []RequestHandleType
}

func (mediator *Mediator) Process(request Interfaces.IRequest[Interfaces.IRequestResult]) (Interfaces.IRequestResult, error) {

	for _, handler := range mediator.handlers {
		if !IsTheRightHandle(&handler, &request) {
			continue
		}
		return handler.Execute(request)
	}
	return nil, errors.New("couldn't find appropriate handler ")
}
func IsTheRightHandle(handel *RequestHandleType, Request RequestType) bool {
	return reflect.TypeOf(Request).AssignableTo(reflect.TypeOf(handel))
}
