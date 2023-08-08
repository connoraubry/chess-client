package client

import (
	"fmt"
	"os"
)

func Config() {
	if len(os.Args) > 2 {
		switch os.Args[2] {
		case "get":
			fmt.Println("Get chosen")
		case "set":
			fmt.Println("set chosen")
		default:
			fmt.Println(ConfigHelp)
		}
	} else {
		fmt.Println(ConfigHelp)
	}
}
