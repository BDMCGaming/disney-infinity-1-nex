package main

import (
	"sync"

	"github.com/BDMCGaming/disney-infinity-1-nex/nex"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go nex.StartAuthenticationServer()
	go nex.StartSecureServer()

	wg.Wait()
}
