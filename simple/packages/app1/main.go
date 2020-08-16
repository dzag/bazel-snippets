package main

import (
	"fmt"
	"os"
	"simple/packages/shared"
)

func main() {
	fmt.Println("Current working dir")
	fmt.Println(os.Getwd())

	config, err := shared.InitConfigs("./packages/configs", "")
	if err != nil {
		os.Exit(3)
	}

	fmt.Printf("Configs: \n%#v\n", config)
}
