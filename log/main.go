package main

import (
	log "mylog/log"
)

// use log.Logger
func main() {
	log.Logger.Info("test info")
	log.Logger.Debug("test debug")
	log.Logger.Error("test error")
}


