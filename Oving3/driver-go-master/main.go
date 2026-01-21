package main

import (
	"Driver-go/elevio"
	"fmt"
)

func main() {

	numFloors := 4

	elevio.Init("localhost:15657", numFloors)

	cab1 := &elevio.Elevator{}

	var d elevio.MotorDirection = elevio.MD_Up
	//cab1.SetMotorDirection(d)

	drv_buttons := make(chan elevio.ButtonEvent)
	drv_floors := make(chan int)
	drv_obstr := make(chan bool)
	drv_stop := make(chan bool)
	OrderChan := make(chan elevio.ButtonEvent)
	BtnPress := make(chan bool)

	go elevio.PollButtons(drv_buttons)
	go elevio.PollFloorSensor(drv_floors, BtnPress)
	go elevio.PollObstructionSwitch(drv_obstr)
	go elevio.PollStopButton(drv_stop)

	for {
		select {
		case a := <-drv_buttons:
			fmt.Printf("%+v\n", a)
			elevio.SetButtonLamp(a.Button, a.Floor, true)
			go elevio.UpdateOrderList(OrderChan, cab1)
			OrderChan <- a
			BtnPress <- true

		case a := <-drv_floors:
			cab1.UpdateFloor(a)
			go cab1.ExecuteOrder()
			cab1.ClearOrderFloor()

			/*

			   fmt.Printf("%+v\n", a)
			   if a == numFloors-1 {
			       d = elevio.MD_Down
			   } else if a == 0 {
			       d = elevio.MD_Up
			   }
			   cab1.SetMotorDirection(d)
			*/

		case a := <-drv_obstr:
			fmt.Printf("%+v\n", a)
			if a {
				cab1.SetMotorDirection(elevio.MD_Stop)
			} else {
				cab1.SetMotorDirection(d)
			}

		case a := <-drv_stop:
			fmt.Printf("%+v\n", a)
			for f := 0; f < numFloors; f++ {
				for b := elevio.ButtonType(0); b < 3; b++ {
					elevio.SetButtonLamp(b, f, false)
				}
			}
		}
	}
}
