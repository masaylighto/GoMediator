package Interfaces

type IRequestHandler[Param IRequest[Response], Response IRequestResult] interface {
	Execute(Param) (Response, error)
}
