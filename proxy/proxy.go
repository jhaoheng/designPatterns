package proxy

import "fmt"

type IProxy interface {
	SetService(s *Service)
	Operation() error
	CheckAccess() error
}

type Proxy struct {
	realService *Service
}

func NewProxy() IProxy {
	return &Proxy{}
}

func (p *Proxy) SetService(s *Service) {
	p.realService = s
}

func (p *Proxy) CheckAccess() error {
	if p.realService == nil {
		return fmt.Errorf("not access service")
	}
	return nil
}

func (p *Proxy) Operation() error {
	if p.realService == nil {
		return fmt.Errorf("not access service")
	}
	p.realService.Operation()
	return nil
}
