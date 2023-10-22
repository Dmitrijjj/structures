package main

import (
	"fmt"
)

type car struct {
	wheels       wheels
	engine       string
	enginehp     float32
	transmissuin string
	suspension   string
}

type wheels struct {
	diametr float32
	Color   string
}

func main() {
	var car1 car
	car1.engine = "дизель"
	car1.wheels = wheels{diametr: 19, Color: "silver"}
	car1.enginehp = 350
	car1.transmissuin = "механика"
	car1.suspension = "пневма"

	var car2 car
	car2.engine = "электро"
	car2.wheels.Color = "gold"
	car2.wheels.diametr = 21.0
	car2.enginehp = 550
	car2.transmissuin = "вариатор"
	car2.suspension = "гидро-пневма"

	var car3 car
	car3.engine = "бензин"
	car3.wheels.Color = "purpure"
	car3.wheels.diametr = 26.0
	car3.enginehp = 850
	car3.transmissuin = "автомат"
	car3.suspension = "пневма"

	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(car3)
}
