package main

import (
	"Network-go/network/bcast"
	//"Network-go/network/localip"
	//"Network-go/network/peers"
	"fmt"
	"time"
	"os"
	"os/exec"

)


const (
	udpPort = 10000
	broadcastAddr = "255.255.255.255"
	heartbeatPeriod = 1 * time.Second
	timeoutPeriod = 3 * time.Second
)



func main() {

	IntTx := make(chan int)
	IntRx := make(chan int)

	go bcast.Transmitter(20014, IntTx)
	go bcast.Receiver(20014, IntRx)


	if len(os.Args) > 1 && os.Args[1] == "backup" {
		runBackup(IntRx, IntTx)
	} else {
		runPrimary(1, IntTx)
	}


	
}


func runPrimary(startCount int, IntTx chan <-int){
	fmt.Printf("Starting primary PID: %d \n", os.Getpid())
	spawnBackup()

	count := startCount
	for{
		if count > 4{
			count = 1
		}
		fmt.Printf("count: %d \n", count)
		IntTx <- count
		count ++
		time.Sleep(heartbeatPeriod)
	}
	
}

func runBackup(IntRx <- chan int, IntTx chan <-int){
	fmt.Printf("starting backup PID: %d \n", os.Getpid())

	lastCount := 0
	lastSeen := time.Now()

	for{
		select{
		case msg := <- IntRx:
			lastCount = msg
			lastSeen = time.Now()
			fmt.Printf("Count received: %d \n", lastCount)
		
		default:
			if time.Since(lastSeen) > timeoutPeriod {
				fmt.Printf("Primary dead! Taking over")
				runPrimary(lastCount + 1, IntTx)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}

}

func spawnBackup() {
	executable, err := os.Executable()
	if err != nil{
		fmt.Println("Error finding exec", err)
		return
	}
	cmd := exec.Command("gnome-terminal", "--", executable, "backup")
	err = cmd.Start()
}




