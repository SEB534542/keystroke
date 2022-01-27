package main

import (
	"log"
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

const interval = 180 * time.Second
const key = "shift

func main() {
	log.Printf("Pressing %v button every %v", key, interval)
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
	// Set shift to be pressed
	kb.HasSHIFT(true)
	// Press the selected keys
	for i := 1; true; i++ {
		log.Println("Round", i)
		kb.Press()
		time.Sleep(10 * time.Millisecond)
		kb.Release()
		time.Sleep(interval)
	}
}
