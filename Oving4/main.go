package main

import (
	"Network-go/network/bcast"
	"Network-go/network/localip"
	"Network-go/network/peers"
	"fmt"
	"time"
	"os"

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

	counter := 0


	initialMode := "primary"
	if len(os.Args) > 1 && os.Args[1] == "backup" {
		initialMode = "backup"
	}

	if initialMode == "primary"{
		runPrimary(1)
	}else{
		runBackup()
	}

	
}


func runPrimary(startCount int){
	fmt.Printf("Starting primary PID: &d \n", os.Getpid())
	spawnBackup()

	count := startCount
	ticker := time.NewTicker(heartbeatPeriod)

	for{
		if count > 4{
			count = 1
		}
		time.Sleep(1 * time.second)
		printf("count: %d \n", count)
		IntTx <- count
		count ++
	}
	
}

func runBackup(){
	fmt:Printf("starting backup PID: %d \n", os.Getpid())

	buffer := make([]byte, 1024)
	lastCount := 0
	lastSeen := time.Now()


	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			if time.Since(lastSeen) > timeoutPeriod {
				fmt.Printf("Primary dead, taking over")
				runPrimary(lastCount + 1)
			}
		}
	}
}

func spawnBackup() {
	executable, err := os.executable()
	if err != nil{
		fmt:Println("Error finding exec", err)
		return
	}
	cmd := exec.command("gnome-terminal", "--", executable, "backup")
	err = cmd.Start()
}




