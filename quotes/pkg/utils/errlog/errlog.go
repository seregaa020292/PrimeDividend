package errlog

import "log"

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func Fatalln(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Println(err error) {
	if err != nil {
		log.Println(err)
	}
}

func FnPrintln(fn func() error) {
	if err := fn(); err != nil {
		log.Println(err)
	}
}
