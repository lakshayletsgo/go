// Mutexs in go
// It is similar to CN and OS
// It is used to lock a particular  channel.
// So that other function can't access it while it is being held by a another function
// To implement it we write
// We first define a mux lock
// We can have multiple reader at the same time but we can't have multiple writers or read-writer at the same time
// So we can have a different lock called sync.RWMutex where it has additional two locks
// 1. RLock()  2.RUnlock()
package main

import (
	"fmt"
	"sync"
)

func sampleMux(muxx *sync.Mutex) {
	muxx.Lock()
	defer muxx.Unlock() //This will unlock the mux lock after this program is executed
	// We can have mutex unlock at the end of the program it is same but we can write it as defer also in the begining it is same
	fmt.Print("Hello")

}

// In This at a time only one progress can get executed and other have to wait
func sampleMux2(muxx *sync.Mutex) {
	muxx.Lock()
	defer muxx.Unlock()

	fmt.Print("Hello2")

}

func main() {
	fmt.Println("Mutexs in go")
}
