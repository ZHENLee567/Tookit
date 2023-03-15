package main

import (
	"github.com/ZHENLee567/Toolkit/cmd"
	"log"
)

var name string

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
