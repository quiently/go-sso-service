package main

import (
	"fmt"

	"github.com/quiently/go-sso-service/internal/config"
)

func main() {
	config := config.MustLoad()
	fmt.Println(config)
}
