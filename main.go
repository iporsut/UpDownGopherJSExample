package main

import (
	"fmt"

	"honnef.co/go/js/console"
	"honnef.co/go/js/dom"
)

func UpdateDisplay(numChan chan int) {
	for {
		select {
		case num := <-numChan:
			console.Log("UpdateDisplay")
			dom.GetWindow().Document().QuerySelector("#num").SetTextContent(fmt.Sprint(num))
		}
	}
}

type NumState struct {
	Num int
}

func (n *NumState) Add() {
	n.Num++
}

func (n *NumState) Sub() {
	n.Num--
}

func main() {
	var numState NumState
	numChan := make(chan int, 1)

	go UpdateDisplay(numChan)

	dom.GetWindow().Document().QuerySelector("#up").AddEventListener("click", false, func(e dom.Event) {
		numState.Add()
		console.Log(numState.Num)
		numChan <- numState.Num
	})

	dom.GetWindow().Document().QuerySelector("#down").AddEventListener("click", false, func(e dom.Event) {
		numState.Sub()
		numChan <- numState.Num
	})
}
