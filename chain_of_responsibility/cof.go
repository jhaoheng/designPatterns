package chainofresponsibility

import (
	"fmt"
	"test/chain_of_responsibility/handler"
)

func ChainOfResposibility() {
	h := handler.NewHandler()
	h.AddMiddleware(FeatureOne())
	h.AddMiddleware(FeatureTwo())
	h.Start()
}

func FeatureOne() handler.MiddlewareFunc {
	return func(next handler.HandlerFunc) handler.HandlerFunc {
		return func() error {
			fmt.Println("hello, FeatureOne")
			return next()
		}
	}
}

func FeatureTwo() handler.MiddlewareFunc {
	return func(next handler.HandlerFunc) handler.HandlerFunc {
		return func() error {
			fmt.Println("hello, FeatureTwo")
			return next()
		}
	}
}
