package cmd

import (
	"log"
	"os"
)

func fatalIf(err error, messages ...string) {
	if err != nil {
		log.Fatal(err, messages)
		os.Exit(1)
	}
}
