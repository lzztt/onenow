package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

func main() {
	flag.Parse()
	n, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		n = 1
	}

	for i := 0; i < n; i++ {
		id := uuid.New()
		fmt.Println(id.String())
	}
}
