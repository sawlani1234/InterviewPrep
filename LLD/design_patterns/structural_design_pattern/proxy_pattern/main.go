package main

import "proxy_pattern/nginx_proxy"


/*
1. Proxy Design Pattern used when u have to do some pre processing and post porcessing 
on original obejct like authentication , caching 

2. Follows Open Closed Principle And Also Single Responsibility Prinicple

3. I.e. multiple proxies can be created for base object 

4. Client can use any type of proxy (interchangable) , when u need some sort 
of lazy initialization


*/

func main() {
	
	nginx := nginxproxy.NewNginx(nginxproxy.NewApplication())
	nginx.HandleRequest("Shubham Proxy Design Pattern")
}