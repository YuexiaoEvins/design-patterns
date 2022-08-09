package Proxy

import "fmt"

type server struct {
}

func NewServer() *server {
	return &server{}
}

func (s *server) handlerRequest(url, para string) {
	fmt.Printf("server handler url %s para %s\n", url, para)
}

type Nginx struct {
	s          *server
	maxRequest int
	urlCounter map[string]int
}

func (n *Nginx) rateLimit(url string) bool {
	if _, ok := n.urlCounter[url]; !ok {
		n.urlCounter[url] = 1
	} else {
		if n.urlCounter[url]++; n.urlCounter[url] > n.maxRequest {
			return true
		}
	}

	return false
}

func NewNginx(maxReq int) *Nginx {
	return &Nginx{
		s:          NewServer(),
		maxRequest: maxReq,
		urlCounter: make(map[string]int, 0),
	}
}

func (n *Nginx) HandlerRequest(url, para string) {
	if !n.rateLimit(url) {
		n.s.handlerRequest(url, para)
	}
}
