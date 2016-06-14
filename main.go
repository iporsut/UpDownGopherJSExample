package main

import (
	"fmt"
	"sync"

	"honnef.co/go/js/dom"
)

func main() {
	num := 0
	var numMutex sync.Mutex
	numChan := make(chan int, 1)

	go func(numChan chan int) {
		for {
			select {
			case num := <-numChan:
				if el := dom.GetWindow().Document().QuerySelector("#num"); el != nil {
					el.SetTextContent(fmt.Sprint(num))
				}
			}
		}
	}(numChan)

	if el := dom.GetWindow().Document().QuerySelector("#up"); el != nil {
		el.AddEventListener("click", false, func(e dom.Event) {
			numMutex.Lock()
			num += 1
			numMutex.Unlock()
			numChan <- num
		})
	}

	if el := dom.GetWindow().Document().QuerySelector("#down"); el != nil {
		el.AddEventListener("click", false, func(e dom.Event) {
			numMutex.Lock()
			num -= 1
			numMutex.Unlock()
			numChan <- num
		})
	}
}
