package main

import (
	"fmt"
	"github.com/bartmeuris/anventory"
)

func main() {

	avset, serr := anventory.LoadSettings("")
	if serr != nil {
		fmt.Printf("ERROR: %s\n", serr)
	}

	av, err := anventory.New(avset)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
	if av == nil {
		fmt.Printf("ERROR: av is nil");
	}
}
