package decorator

type Pizza interface {
	GetPrice() int
}

type BasicPizza struct {
}

func NewBasicPizze() Pizza {
	return &BasicPizza{}
}

func (b *BasicPizza) GetPrice() int {
	return 15
}
