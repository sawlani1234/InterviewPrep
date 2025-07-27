package nginxproxy

import "fmt"



type Application struct {

}

func NewApplication() Application {
	return Application{}
}

func (a Application) HandleRequest(r interface{}) {
	fmt.Println("Request is now handled by Applicaiton")
}