package main

import (
	log "./log"
)

// test log.Logger
func main() {
	logger := log.Logger
	logger.Info("test info")
	logger.Debug("test debug")
	logger.Error("test error")
}



