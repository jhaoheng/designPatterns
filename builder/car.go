package builder

import "fmt"

type Color string
type Wheel string
type Light string

const (
	RedColor   Color = "red"
	WhiteColor Color = "white"
	BlueColor  Color = "blue"
)

const (
	RaceWheel    Wheel = "race"
	GeneralWheel Wheel = "general"
)

const (
	Funture Light = "future"
	Retro   Light = "retro"
)

type ICar interface {
	SetColor(color Color) *Car
	SetWheel(wheel Wheel) *Car
	SetLight(light Light) *Car
	Build()
}

type Car struct {
	ColorType Color
	WheelType Wheel
	LightType Light
}

func NewCar() ICar {
	return &Car{}
}

func (c *Car) SetColor(color Color) *Car {
	c.ColorType = color
	return c
}

func (c *Car) SetWheel(wheel Wheel) *Car {
	c.WheelType = wheel
	return c
}

func (c *Car) SetLight(light Light) *Car {
	c.LightType = light
	return c
}

func (c *Car) Build() {
	fmt.Printf("%#v\n", *c)
}
