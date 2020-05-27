package cmd

import "log"

func fatalIf(err error, messages ...string) {
	if err != nil {
		log.Fatal(err, messages)
	}
}
