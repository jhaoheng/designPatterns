package facade

type IFacade interface {
	DoSomething()
}

type Facade struct{}

func NewFacede() IFacade {
	return &Facade{}
}

func (f *Facade) DoSomething() {

}
