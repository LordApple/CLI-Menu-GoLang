package main

import (
	"fmt"
	"os"
	"strings"
)

func mainMenu() {
	const asciiart = `
	___  ___                 
	|  \/  |                 
	| .  . | ___ _ __  _   _ 
	| |\/| |/ _ \ '_ \| | | |
	| |  | |  __/ | | | |_| |
	\_|  |_/\___|_| |_|\__,_|
							 




							
	`
	fmt.Println(strings.Repeat("\n", 1000) + asciiart)

	fmt.Println("[1] Menu")
	fmt.Println("[2] Menu")
	fmt.Println("[3] Exit")

	switch ask([]string{"1", "2", "3"}) {
	case "1":
		menu1()
	case "2":
		menu2()
	case "3":
		os.Exit(404)
	}

}

func menu1() {
	const asciiart = `
	___  ___                 
	|  \/  |                 
	| .  . | ___ _ __  _   _ 
	| |\/| |/ _ \ '_ \| | | |
	| |  | |  __/ | | | |_| |
	\_|  |_/\___|_| |_|\__,_|
							 




							
	`
	fmt.Println(strings.Repeat("\n", 1000) + asciiart)

	fmt.Println("[1] Menu")
	fmt.Println("[2] Back")

	switch ask([]string{"1", "2"}) {
	case "1":
		//Do stuff here
	case "2":
		mainMenu()
	}

}

func menu2() {
	const asciiart = `
	___  ___                 
	|  \/  |                 
	| .  . | ___ _ __  _   _ 
	| |\/| |/ _ \ '_ \| | | |
	| |  | |  __/ | | | |_| |
	\_|  |_/\___|_| |_|\__,_|
							 




							
	`
	fmt.Println(strings.Repeat("\n", 1000) + asciiart)

	fmt.Println("[1] Menu")
	fmt.Println("[2] Menu")
	fmt.Println("[3] Back")

	switch ask([]string{"1", "2", "3"}) {
	case "1":
		//Do stuff here
	case "2":
		//Do stuff here
	case "3":
		mainMenu()
	}

}
