package main

import (
	"fmt"
	"strings"
)

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string) string
}

type BaseHandler struct {
	nextHandler Handler
}

func (b *BaseHandler) SetNext(handler Handler) Handler {
	b.nextHandler = handler
	return handler
}

func (b *BaseHandler) Handle(request string) string {
	if b.nextHandler != nil {
		return b.nextHandler.Handle(request)
	}
	return ""
}

type StringHandler struct {
	BaseHandler
}

func (s *StringHandler) Handle(request string) string {
	if strings.Contains(request, "string") {
		return "StringHandler: Handling request"
	}
	return s.BaseHandler.Handle(request)
}

type NumberHandler struct {
	BaseHandler
}

func (n *NumberHandler) Handle(request string) string {
	if strings.Contains(request, "number") {
		return "NumberHandler: Handling request"
	}
	return n.BaseHandler.Handle(request)
}

// Пример использования
func main() {
	stringHandler := &StringHandler{}
	numberHandler := &NumberHandler{}

	// Установка связей между обработчиками
	stringHandler.SetNext(numberHandler)

	// Обработка запросов
	fmt.Println(stringHandler.Handle("This is a string request")) // Output: StringHandler: Handling request
	fmt.Println(stringHandler.Handle("This is a number request")) // Output: NumberHandler: Handling request
}
