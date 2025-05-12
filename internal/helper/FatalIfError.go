package helper

import "log"

func FatalIfErrorWithMessage(err error, message string) {

	if err != nil {
		log.Fatal(message)
	}

}
