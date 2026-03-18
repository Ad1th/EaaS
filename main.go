package main

import (
	"fmt"
	"os"
)

func main() {
	dry := false
	corporate := false
	reassure := false
	optimistic := false
	silent := false

	for _, arg := range os.Args[1:] {
		switch arg {
		case "--dry":
			dry = true
		case "--corporate":
			corporate = true
		case "--reassure":
			reassure = true
		case "--optimistic":
			optimistic = true
		case "--silent":
			silent = true
		}
	}

	if silent {
		return
	}

	switch {
	case dry:
		fmt.Println("Noted.")
	case corporate:
		fmt.Println("On the roadmap.")
	case reassure:
		fmt.Println("This can wait.")
	case optimistic:
		fmt.Println("It will happen.")
	default:
		fmt.Println("Eventually.")
	}
}