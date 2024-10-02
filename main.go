package main

import (
	"fmt"

	"github.com/lucassantos-dev/go-api/config"
	"github.com/lucassantos-dev/go-api/router"
)

func main() {
	err := config.Init()

	if err != nil {
		fmt.Println(err)
		return
	}
	router.Initializer()
}
