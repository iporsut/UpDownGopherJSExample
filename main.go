package main

import (
	"io/ioutil"
	"net/http"

	"honnef.co/go/js/console"
	"honnef.co/go/js/dom"
)

func main() {
	if el := dom.GetWindow().Document().QuerySelector("#load"); el != nil {
		el.AddEventListener("click", false, func(e dom.Event) {
			go func() {
				resp, err := http.Get("http://localhost:8080/hellogopherjs/data.json")
				if err != nil {
					panic(err)
				}
				b, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					panic(err)
				}
				console.Log(string(b))
			}()
			console.Log("Loading")
		})
	}
}
