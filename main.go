package main

import (
	"fmt"

	"honnef.co/go/js/dom"
)

type CounterEvent int

const (
	DownEvent CounterEvent = iota
	UpEvent
)

type NumState struct {
	Num int
}

func (n *NumState) Add() {
	n.Num++
}

func (n *NumState) Sub() {
	n.Num--
}

func ReceiveCouterEvent(eventChan chan CounterEvent, numChan chan int, numState *NumState) {
	for {
		switch <-eventChan {
		case UpEvent:
			numState.Add()
		case DownEvent:
			numState.Sub()
		}
		numChan <- numState.Num
	}
}

func UpEventListener(eventChan chan CounterEvent) {
	dom.GetWindow().Document().QuerySelector("#up").AddEventListener("click", false, func(e dom.Event) {
		eventChan <- UpEvent
	})
}

func DownEventListener(eventChan chan CounterEvent) {
	dom.GetWindow().Document().QuerySelector("#down").AddEventListener("click", false, func(e dom.Event) {
		eventChan <- DownEvent
	})
}

func UpdateDisplay(numChan chan int) {
	for {
		select {
		case num := <-numChan:
			dom.GetWindow().Document().QuerySelector("#num").SetTextContent(fmt.Sprint(num))
		}
	}
}

func main() {
	var numState NumState
	eventChan := make(chan CounterEvent, 1)
	numChan := make(chan int, 1)

	go UpdateDisplay(numChan)
	go ReceiveCouterEvent(eventChan, numChan, &numState)
	UpEventListener(eventChan)
	DownEventListener(eventChan)
}
