package utils

import "log"

func Recovered() {
	if err := recover(); err != nil {
		log.Println(err)
	}
}
