package handler

import "log"

func HandleError(err error) {
	if err != nil {
		log.Print(err.Error())
	}
}
