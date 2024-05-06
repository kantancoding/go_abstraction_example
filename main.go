package main

import (
	"fmt"
	"math"

	"github.com/kantancoding/abstraction/vm"
)

type vendingMachine interface {
	GetDrink(money int64, brand string) string
}

type Application struct {
	vm vendingMachine
}

func (this Application) Run() {
	//complex
	//complex
	//complex
	fmt.Println(math.Abs(-242))
	//complex
	//complex
	//complex
	myDrink := this.vm.GetDrink(100, "Cola")
	fmt.Println(myDrink)
}

func newApplication(vm vendingMachine) *Application {
	return &Application{vm: vm}
}

func main() {
	vendingMachine := vm.New()
	app := newApplication(vendingMachine)
	app.Run()
}
