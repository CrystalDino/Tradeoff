package core

import "log"

func Recover(info string) {
	var i interface{}
	if i = recover(); i == nil {
		return
	}
	log.Printf("panic at: %s : %+v\n", info, i)
}
