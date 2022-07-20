package decorator

type TomatoTopping struct {
	Pizza Pizza
}

func (t *TomatoTopping) GetPrice() int {
	return t.Pizza.GetPrice() + 7
}
