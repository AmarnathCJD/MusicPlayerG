package util

import (
	"os"
	"os/signal"
)

func GetSpecifiedPort() string {
	portENV := os.Getenv("PORT")
	if portENV == "" {
		return ":80"
	}
	return ":" + portENV
}

func InterruptHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	os.Exit(0)
}
