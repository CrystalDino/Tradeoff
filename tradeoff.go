package main

import (
	"Tradeoff/core"
	"Tradeoff/models"
)

func main() {
	println("Tradeoff server starting")
	if err := core.DBLinkTest(); err != nil {
		println("db ping error,", err)
	}
	models.Show()
	core.WebRun()
}
