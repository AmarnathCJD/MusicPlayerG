package util

import (
	"os"
	"os/signal"
	"strconv"
)

// returns the port to listen on, defaults to 8080 if not specified
func GetSpecifiedPort() string {
	portENV := os.Getenv("PORT")
	if portENV == "" {
		return ":80"
	}
	return ":" + portENV
}

// waits for an interrupt signal and gracefully exits the server
func InterruptHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	os.Exit(0)
}

func Itoa(i int) string {
	return strconv.Itoa(i)
}

func ItoaF(i int) string {
	if i >= 1000000000 {
		return strconv.Itoa(i/1000000000) + "B"
	} else if i >= 1000000 {
		return strconv.Itoa(i/1000000) + "M"
	} else if i >= 1000 {
		return strconv.Itoa(i/1000) + "K"
	}
	return strconv.Itoa(i)
}
