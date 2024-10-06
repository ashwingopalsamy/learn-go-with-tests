package main

import (
	"fmt"
	"time"
)

type Vehicle interface {
	Start()
	Stop()
	GetInfo() string
}

type VehicleDetails interface {
	Name() string
	Make() string
	Brand() string
	Capacity() string
}

type VehicleAttr struct {
	name     string
	make     string
	brand    string
	capacity string
}

type VehicleBehavior struct {
	VehicleDetails
}

type Car struct {
	VehicleBehavior
}

type Bike struct {
	VehicleBehavior
}

type Truck struct {
	VehicleBehavior
}

func (v VehicleAttr) Name() string     { return v.name }
func (v VehicleAttr) Make() string     { return v.make }
func (v VehicleAttr) Brand() string    { return v.brand }
func (v VehicleAttr) Capacity() string { return v.capacity }

func (v *VehicleBehavior) Start() {
	fmt.Printf("%s %s Engine Started!\n", v.Name(), v.Brand())
}

func (v *VehicleBehavior) Stop() {
	fmt.Printf("%s %s Engine Stopped!\n", v.Name(), v.Brand())
}

func main() {
	vehicles := []Vehicle{
		NewCar("City", "Honda", "2022"),
		NewBike("R15", "Yamaha"),
		NewTruck("Cruize", "Bharat Benz", "150 Tonnes"),
	}

	operateVehicles(vehicles)
}

func operateVehicles(ve []Vehicle) {
	for _, v := range ve {
		v.GetInfo()
		v.Start()
		time.Sleep(time.Millisecond * 500)
	}
}

func NewCar(name, brand, make string) *Car {
	if name == "" || brand == "" {
		panic("Car name and brand cannot be empty")
	}
	return &Car{VehicleBehavior: VehicleBehavior{VehicleAttr{name, brand, make, ""}}}
}

func NewBike(name, brand string) *Bike {
	if name == "" || brand == "" {
		panic("Bike name and brand cannot be empty")
	}
	return &Bike{VehicleBehavior: VehicleBehavior{VehicleAttr{name, brand, "", ""}}}
}

func NewTruck(name, brand, capacity string) *Truck {
	if name == "" || brand == "" {
		panic("Truck name and brand cannot be empty")
	}
	return &Truck{VehicleBehavior: VehicleBehavior{VehicleAttr{name, brand, "", capacity}}}
}

func (c *Car) GetInfo() string {
	return fmt.Sprintf("Car %s (%s)", c.Name(), c.Brand())
}

func (b *Bike) GetInfo() string {
	return fmt.Sprintf("Bike %s (%s)", b.Name(), b.Brand())
}

func (t *Truck) GetInfo() string {
	return fmt.Sprintf("Truck %s (%s) with Capacity %s", t.Name(), t.Brand(), t.Capacity())
}
