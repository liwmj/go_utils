package go_utils

import "log"

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintError(err error) {
	if err != nil {
		log.Println(err)
	}
}
