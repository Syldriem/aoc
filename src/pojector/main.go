package main

import (
	"fmt"
	"log"

	"github.com/Syldriem/aoc/pkg/projector"
)

func main() {
    opts, err := projector.GetOpts()
    if err != nil {
        log.Fatalf("unable to get options %v", err)
    }

    fmt.Printf("opts: %+v", opts)
}
