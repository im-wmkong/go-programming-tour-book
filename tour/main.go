package main

import (
	"log"

	"github.com/go-programming-tour-book/tour/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
