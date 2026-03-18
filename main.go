package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		runLoop("")
		return
	}

	message := "Eventually."
	switch {
	case dry:
		message = "Noted."
	case corporate:
		message = "On the roadmap."
	case reassure:
		message = "This can wait."
	case optimistic:
		message = "It will happen."
	}

	runLoop(message)
}

func runLoop(message string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "eaas exit" {
			return
		}
		if message != "" {
			fmt.Println(message)
		}
	}
}