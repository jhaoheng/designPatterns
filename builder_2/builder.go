package builder2

/*
Builder interface
*/
type Builder interface {
	setName(name string)
	setAge(age int)
	setSound(sound string)
	getResult() map[string]interface{}
}

// Direct struct
type Director struct {
	builder Builder
}

// NewDirector ...
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

//Construct Product
func (d *Director) Construct(name string, age int, sound string) {
	d.builder.setName(name)
	d.builder.setAge(age)
	d.builder.setSound(sound)
	d.builder.getResult()
}

// ConcreteBuilder
type DogBuilder struct {
	name  string
	age   int
	sound string
}

func (d *DogBuilder) setSound(sound string) { d.sound = sound } //WoofWoof
func (d *DogBuilder) setName(name string)   { d.name = name }
func (d *DogBuilder) setAge(age int)        { d.age = age }
func (d *DogBuilder) getResult() map[string]interface{} {
	return map[string]interface{}{
		"name":  d.name,
		"age":   d.age,
		"sound": d.sound,
	}
}

type CatBuilder struct {
	name  string
	age   int
	sound string
}

func (c *CatBuilder) setSound(sound string) { c.sound = sound } //MeowMeow
func (c *CatBuilder) setName(name string)   { c.name = name }
func (c *CatBuilder) setAge(age int)        { c.age = age }
func (c *CatBuilder) getResult() map[string]interface{} {
	return map[string]interface{}{
		"name":  c.name,
		"age":   c.age,
		"sound": c.sound,
	}
}

// func main() {
// 	//
// 	dog := NewDirector(&DogBuilder{})
// 	dog.Construct("Snoppy", 100, "WoofWoof")
// 	fmt.Println(dog.builder.getResult())
// 	fmt.Println(reflect.TypeOf(dog.builder.getResult()))

// 	//
// 	cat := NewDirector(&CatBuilder{})
// 	cat.Construct("Kitty", 100, "MeowMeow")
// 	fmt.Println(cat.builder.getResult())
// }
