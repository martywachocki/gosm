package main

import (
	"fmt"
)

// CurrentConfig The current application configuration
var CurrentConfig Config

func main() {
	CurrentConfig = ParseConfigFile()
	fmt.Println(CurrentConfig)
}
