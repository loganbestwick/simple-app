package main

import (
	"fmt"

	"github.com/loganbestwick/simple-app/cmd/api"
)

func main() {
	fmt.Println("Starting App...")
	api.Setup()
}
