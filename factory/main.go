package main

import (
	"factory/pkg"
)

func main() {
	pkg.NewNotification(pkg.FatherType).DoPhone()
	pkg.NewNotification(pkg.FatherType).DoMessage("hello")
}
