package main

import (
	"log"
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

const interval = 180 * time.Second

func main() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	// Select keys to be pressed: A + B)
	//kb.SetKeys(keybd_event.VK_A, keybd_event.VK_B)

	// Set shift to be pressed
	kb.HasSHIFT(true)

	// Press the selected keys
	for i := 1; true; i++ {
		log.Println("Round", i)
		err = kb.Launching()
		if err != nil {
			panic(err)
		}
		time.Sleep(interval)
	}
	// Or you can use Press and Release
	// kb.Press()
	// time.Sleep(10 * time.Millisecond)
	// kb.Release()
}