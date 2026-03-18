package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
	interactive := isTerminal(os.Stdin) && isTerminal(os.Stdout)
	reader := bufio.NewReader(os.Stdin)

	for {
		if interactive {
			fmt.Print(prompt())
		}

		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}

		cmd := strings.TrimSpace(line)
		if cmd == "eaas exit" {
			return
		}
		if cmd != "" && message != "" {
			fmt.Println(message)
		}
	}
}

func isTerminal(f *os.File) bool {
	info, err := f.Stat()
	if err != nil {
		return false
	}
	return (info.Mode() & os.ModeCharDevice) != 0
}

func prompt() string {
	wd, err := os.Getwd()
	if err != nil {
		return "% "
	}
	name := filepath.Base(wd)
	if name == "" || name == "." || name == string(filepath.Separator) {
		return "% "
	}
	return name + " % "
}
