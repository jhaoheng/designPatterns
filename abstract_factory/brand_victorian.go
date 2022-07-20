package abstractfactory

import "fmt"

type VictorianFactory struct {
}

func NewVictorianFactory() *VictorianFactory {
	return &VictorianFactory{}
}

func (v *VictorianFactory) CreateChair() IChair {
	vc := &VictorianChair{}
	return vc
}

func (v *VictorianFactory) CreateCoffeeTable() ICoffeeTable {
	vct := &VictorianCoffeeTable{}
	return vct
}

func (v *VictorianFactory) CreateSofa() ISofa {
	vs := &VictorianSofa{}
	return vs
}

//
type VictorianChair struct {
}

func (vc *VictorianChair) SetColor() IChair {
	return vc
}

func (vc *VictorianChair) SetLegs() IChair {
	return vc
}

func (vc *VictorianChair) Produce() {
	fmt.Println("Victorian Chair")
}

//
type VictorianCoffeeTable struct{}

func (vct *VictorianCoffeeTable) SetColor() ICoffeeTable {
	return vct
}

func (vct *VictorianCoffeeTable) Produce() {
	fmt.Println("Victorian Coffee Table")
}

//
type VictorianSofa struct{}

func (vs *VictorianSofa) SetColor() ISofa {
	return vs
}

func (vs *VictorianSofa) Produce() {
	fmt.Println("Victorian Sofa")
}
