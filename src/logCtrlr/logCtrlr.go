package logCtrlr

import (
	"log"
)

func DError(err error) {
	log.Println("*** (ERROR)", err)
}

func Error(err error) {
	log.Println("** (ERROR)", err)
}

func Log(message string) {
	log.Println("* (SYSTEM)" + message)
}
