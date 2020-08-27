package cmd

import (
	"log"
)

// fatalIf help on manege the fatal messages
func fatalIf(err error, messages ...string) {
	if err != nil {
		log.Fatal(err, messages)
	}
}
