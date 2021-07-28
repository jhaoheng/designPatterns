package pkg

type Pizza interface {
	GetPrice() int
}

//
type VeggeMania struct {
}

func (v *VeggeMania) GetPrice() int {
	return 15
}

//
type TomatoTopping struct {
	Pizza Pizza
}

func (t *TomatoTopping) GetPrice() int {
	return t.Pizza.GetPrice() + 7
}

//
type CheeseTopping struct {
	Pizza Pizza
}

func (c *CheeseTopping) GetPrice() int {
	return c.Pizza.GetPrice() + 10
}
