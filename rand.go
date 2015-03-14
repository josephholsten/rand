package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
)

func main() {
	log.SetFlags(0) // disable time in output
	args := os.Args
	name := path.Base(args[0])

	var i int

	if len(args) >= 2 {
		n, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("%s: not an integer: %s", name, args[1])
		}
		i = rand.Intn(n)
	} else {
		i = rand.Int()
	}

	fmt.Println(i)
}
