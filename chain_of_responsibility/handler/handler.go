package handler

import (
	"fmt"
)

type IHandler interface {
	AddMiddleware(m MiddlewareFunc)
	Start()
}

type Hanlder struct {
	Middlewares []MiddlewareFunc
}

var NotFoundHandler = func() error {
	return fmt.Errorf("error not found")
}

type HandlerFunc func() error

type MiddlewareFunc func(HandlerFunc) HandlerFunc

func NewHandler() IHandler {
	h := &Hanlder{
		Middlewares: []MiddlewareFunc{},
	}
	return h
}

func (h *Hanlder) AddMiddleware(m MiddlewareFunc) {
	h.Middlewares = append(h.Middlewares, m)
}

func (h *Hanlder) Start() {
	e := NotFoundHandler
	for i := 0; i < len(h.Middlewares); i++ {
		f := h.Middlewares[i](e)
		f()
	}
}
