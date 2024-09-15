package main

import (
	"fmt"
	"github.com/rinnothing/boyscout"

	"github.com/rinnothing/boyscout/examples/extension"
)

type Speaker interface {
	Speak() (string, error)
}

func main() {
	println("Hello World")
	extension.Hook()

	speakers := boyscout.DefaultGrandscout.GetScout("speakers")
	list, err := speakers.List()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < 5; i++ {
		for _, bucket := range list {
			speaker, ok := bucket.Val.(Speaker)
			if !ok {
				fmt.Printf("Couldn't convert %s to Speaker interface\n", bucket.Name)
				continue
			}

			speech, err := speaker.Speak()
			if err != nil {
				fmt.Printf("Speaker %s couldn't speak due to %s error\n", bucket.Name, err.Error())
				continue
			}
			fmt.Printf("Speaker %s said %s", bucket.Name, speech)
		}
	}
}
