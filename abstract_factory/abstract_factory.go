package abstractfactory

type IAbstractFactroy interface {
	CreateChair() IChair
	CreateCoffeeTable() ICoffeeTable
	CreateSofa() ISofa
}

type IChair interface {
	SetColor() IChair
	SetLegs() IChair
	Produce()
}

type ICoffeeTable interface {
	SetColor() ICoffeeTable
	Produce()
}

type ISofa interface {
	SetColor() ISofa
	Produce()
}

func NewAbstractFactroy(brand interface{}) IAbstractFactroy {
	if v, ok := brand.(*VictorianFactory); ok {
		return v
	} else if m, ok := brand.(*ModernFactory); ok {
		return m
	}

	return &VictorianFactory{}
}
