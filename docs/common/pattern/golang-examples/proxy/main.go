package main

import "fmt"

type Subject interface {
	Request() string
}

type RealSubject struct{}

func (r *RealSubject) Request() string {
	return "RealSubject: Handling request"
}

type Proxy struct {
	realSubject *RealSubject
}

func NewProxy() *Proxy {
	return &Proxy{}
}

func (p *Proxy) Request() string {
	if p.realSubject == nil {
		fmt.Println("Proxy: Creating real subject object")
		p.realSubject = &RealSubject{}
	}
	fmt.Println("Proxy: Forwarding request to real subject")
	return p.realSubject.Request()
}

// Пример использования
func main() {
	proxy := NewProxy()
	fmt.Println(proxy.Request()) // Output: Proxy: Creating real subject object
	//         Proxy: Forwarding request to real subject
	//         RealSubject: Handling request
	fmt.Println(proxy.Request()) // Output: Proxy: Forwarding request to real subject
	//         RealSubject: Handling request
}
