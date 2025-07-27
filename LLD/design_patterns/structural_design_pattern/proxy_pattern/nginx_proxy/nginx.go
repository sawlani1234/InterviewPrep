package nginxproxy

import "fmt"



type Nginx struct {
    a Application
}

func NewNginx(a Application) Nginx {
	return Nginx{a}
}

func (n Nginx) isRateLimited(r interface{}) {
	fmt.Println("nginx proxy rate limiter activated")
}

func (n Nginx) isCached(r interface{}) {
	fmt.Println("Is Request Cached")
}

func (n Nginx) HandleRequest(r interface{}) {

	n.isRateLimited(r)
	n.isCached(r)
	fmt.Println("nginx checks completed routing to application")

	n.a.HandleRequest(r)

}