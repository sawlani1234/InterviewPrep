package nginxproxy


type RequestHandler interface {
	HandleRequest(r interface{})
}