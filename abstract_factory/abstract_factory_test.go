package abstractfactory

import "testing"

func Test_Victorian(t *testing.T) {
	v := NewAbstractFactroy(NewVictorianFactory())
	v.CreateChair().SetColor().SetLegs().Produce()
	v.CreateCoffeeTable().SetColor().Produce()
	v.CreateSofa().SetColor().Produce()
}

func Test_Modern(t *testing.T) {
	m := NewAbstractFactroy(NewModernFactory())
	m.CreateChair().SetColor().SetLegs().Produce()
	m.CreateCoffeeTable().SetColor().Produce()
	m.CreateSofa().SetColor().Produce()
}
