// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	// "time"
)

func incrementing(ch2 chan int, done chan bool) {
	for k := 0; k < 1_000_000; k++ {
		ch2 <- 1
	}
	done <- true
	//TODO: increment i 1000000 times
}

func decrementing(ch1 chan int, done chan bool) {
	for k := 0; k < 500_000; k++ {
		ch1 <- 1
	}
	done <- true
	//TODO: decrement i 1000000 times
}

func server(ch1 chan int, ch2 chan int, ferdig chan int, svar chan int) {
	var i = 0
	for {
		select {
		case <-ch1:
			i--
		case <-ch2:
			i++
		case <-ferdig:
			svar <- i
		}

	}

}


func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)

	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan bool)
	ferdig := make(chan int)
	svar := make(chan int)

	go server(ch1, ch2, ferdig, svar)

	go incrementing(ch2, done)
	go decrementing(ch1, done)
	// TODO: Spawn both functions as goroutines

	<-done
	<-done

	ferdig <- 1

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	//time.Sleep(500 * time.Millisecond)
	resultat := <-svar
	Println("The magic number is:", resultat)
}
