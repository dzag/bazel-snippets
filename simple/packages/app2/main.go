package main

import (
	"fmt"
	"os"
	"simple/packages/shared"
)

const DEFAULT_ENV = "development"

func main() {
	fmt.Println("Current working dir")
	fmt.Println(os.Getwd())
	fmt.Println(len(os.Args))

	env := ""
	if (len(os.Args) <= 1) {
		env = DEFAULT_ENV
	} else {
		env = os.Args[1]
	}

	if env == "" {
		env = DEFAULT_ENV
	}

	config, err := shared.InitConfigs("./packages/configs", env)
	if err != nil {
		os.Exit(3)
	}

	fmt.Printf("Configs: \n%#v\n", config)
}
