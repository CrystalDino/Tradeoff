package main

import "Tradeoff/core"

func main() {
	println("Tradeoff server starting")
	if err := core.DBLinkTest(); err == nil {
		println("db linked")
	} else {
		println("db ping error,", err)
	}
	core.WebRun()
}
