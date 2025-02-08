package main

import (
	"fmt"
	"sync"
)

// import (
// 	"fmt"

// 	"github.com/quiently/go-sso-service/internal/config"
// )

// func main() {
// 	config := config.MustLoad()
// 	fmt.Println(config)

// }
var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup
	msg = "Hello, world!"


	message("Hello, universe!", &wg)
	message("Hello, cosmos!", &wg)
	message("Hello, world!", &wg)
}

func message(message string, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		updateMessage(message)
	}()
}
