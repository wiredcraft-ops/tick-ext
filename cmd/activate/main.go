// how many agents are running
package main

import (
	"log"

	"github.com/Wiredcraft/tick-ext"
)

func main() {

	log.SetFlags(log.Lshortfile)

	uuid, err := tick.GetUUID()
	if err != nil {
		log.Fatalf("E! %v", err)
	}
	err = tick.Store(string(uuid))
	if err != nil {
		log.Printf("E! %v\n", err)
	}
}
