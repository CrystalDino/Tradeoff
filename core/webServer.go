package core

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	WebEngine *gin.Engine
	webAddr   string
)

func init() {
	webAddr = GetConfigString("addr")
	if webAddr == "" {
		log.Fatalln("web engine init error, web addr not set.")
		return
	}
	WebEngine = gin.New()
	// Global middleware
	WebEngine.Use(gin.Logger())
	WebEngine.Use(gin.Recovery())
	//other configure
}

func WebRun() (err error) {
	if WebEngine == nil {
		return errors.New("web engine not running")
	}
	return WebEngine.Run(webAddr)
}
