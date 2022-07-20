package factory

import (
	"testing"
)

func Test_factory(t *testing.T) {

	//
	NewNotification(FatherType).DoPhone()
	NewNotification(FatherType).DoMessage("hello")
	//
	NewNotification(MotherType).DoPhone()
	NewNotification(MotherType).DoMessage("hello")
}
