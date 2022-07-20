package abstractfactory

import "fmt"

//
type ModernFactory struct {
}

func NewModernFactory() *ModernFactory {
	return &ModernFactory{}
}

func (m *ModernFactory) CreateChair() IChair {
	return &ModernChair{}
}
func (m *ModernFactory) CreateCoffeeTable() ICoffeeTable {
	return &ModernCoffeeTable{}
}
func (m *ModernFactory) CreateSofa() ISofa {
	return &ModernSofa{}
}

//
type ModernChair struct {
}

func (mc *ModernChair) SetColor() IChair {
	return mc
}

func (mc *ModernChair) SetLegs() IChair {
	return mc
}

func (mc *ModernChair) Produce() {
	fmt.Println("Modern Chair")
}

//
type ModernCoffeeTable struct{}

func (mct *ModernCoffeeTable) SetColor() ICoffeeTable {
	return mct
}

func (mct *ModernCoffeeTable) Produce() {
	fmt.Println("Modern Coffee Table")
}

//
type ModernSofa struct{}

func (ms *ModernSofa) SetColor() ISofa {
	return ms
}

func (ms *ModernSofa) Produce() {
	fmt.Println("Modern Sofa")
}
