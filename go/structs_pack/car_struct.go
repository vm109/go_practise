package structs_pack

import (
	"strconv"
)

type Car struct {
	gear        bool
	seats       int8
	wheels      int8
	heatedSeats bool
}

func (c Car) CarAdvertise(modelName string) string{
	return "the car model is "+ modelName+ " and the car has "+ strconv.Itoa(int(c.seats)) + "seats and "+ strconv.Itoa(int(c.wheels))+ " wheels"
}

func (c *Car) ChangeCarWheelsOrSeats(wheels, seats int8){
	c.seats = seats
	c.wheels = wheels
}
func Create_Car(seats, wheels int8, gear, heatedseats bool) Car{

	car := Car{
		gear:        gear,
		seats:       seats,
		wheels:      wheels,
		heatedSeats: heatedseats,
	}
	return car
}
