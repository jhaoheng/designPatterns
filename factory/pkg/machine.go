package pkg

import "fmt"

type TargetType int

const (
	FatherType TargetType = 1 << iota
	MotherType
)

type INotification interface {
	DoPhone() INotification
	DoMessage(message string) INotification
}

type Notification struct {
	Name  string
	Age   string
	Phone string
}

func NewNotification(nt TargetType) INotification {
	n := &Notification{}
	switch nt {
	case FatherType:
		n.Name = "father"
		n.Age = "36"
		n.Phone = "0919111111"
	case MotherType:
		n.Name = "mother"
		n.Age = "30"
		n.Phone = "0919222222"
	}
	return n
}

func (n *Notification) DoPhone() INotification {
	fmt.Println("call phone")
	return n
}

func (n *Notification) DoMessage(message string) INotification {
	fmt.Println("send message")
	return n
}
