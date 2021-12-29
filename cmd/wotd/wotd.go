package main

import (
	"context"
	"fmt"

	"github.com/tstromberg/wotd/pkg/wotd"
)

func main() {
	rs, err := wotd.All(context.Background())
	if err != nil {
		panic(err.Error())
	}

	for _, r := range rs {
		fmt.Printf("%s —  %s\n", r.Word, r.URL)
		for _, d := range r.Definitions {
			for _, ps := range d.Parts {
				fmt.Printf("    %-9.9s: %s\n", ps.Kind, ps.Text)
			}

			// Only show the first definition
			break
		}

		/* Skip examples - they are wordy.
		if len(r.Examples) > 0 {
			for _, ex := range r.Examples {
				fmt.Printf("  example: %s\n", ex.Text)
			}
		}
		*/
		fmt.Printf("\n")
	}
}
