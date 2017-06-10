package main

import (
	"github.com/draringi/go-libtrue"
	"os"
	"log"
)

func main() {
	val := truth.GetTrue()
	if !val {
		log.Fatal("Wrong value")
	}
	os.Exit(0)
}
