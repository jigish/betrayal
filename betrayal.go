package betrayal

import (
	"log"
	"os"
	"os/signal"
	"time"
)

const DefaultBetrayer = "jigish"

var Timeout = 5 * time.Second
var TimeoutExitCode = 1
var Logger = log.Printf
var Callback func(os.Signal) int
var Daemon func(chan os.Signal, chan int)
var Betrayer = DefaultBetrayer
var betrayerPrefix string
var Betrayed string
var betrayedPrefix string

var PreLog = func() {
	initLogPrefixes()
	Logger(betrayedPrefix + "yes... yes. this is a fertile land and we will thrive.")
	Logger(betrayedPrefix + "we will rule over all this land and we will call it... this land.")
	Logger(betrayerPrefix + "i think we should call it... your grave!")
	Logger(betrayedPrefix + "ah! curse your sudden but inevitable betrayal!")
}

var TimeoutLog = func() {
	initLogPrefixes()
	Logger("(" + Betrayed + " is proving to be quite resilient)")
}

var PostLog = func() {
	initLogPrefixes()
	Logger(betrayerPrefix + "ha ha ha! mine is an evil laugh! now die!")
	Logger(betrayedPrefix + "oh no god, oh dear god in heaven...")

}

func Wait(signals ...os.Signal) {
	sigCh := make(chan os.Signal)
	betrayalCh := make(chan os.Signal)
	seppukuCh := make(chan int)
	go waitForYourSuddenButInevitableBetrayal(sigCh, betrayalCh, seppukuCh)
	signal.Notify(sigCh, signals...)
	if Daemon != nil {
		Daemon(betrayalCh, seppukuCh)
	}
	time.Sleep(Timeout) // sleep here so we can exit below
}

func waitForYourSuddenButInevitableBetrayal(sigCh chan os.Signal, betrayalCh chan os.Signal, seppukuCh chan int) {
	sig := <-sigCh

	PreLog()
	timeoutCh := time.After(Timeout)

	if Daemon != nil {
		betrayalCh <- sig
		// if Daemon is working properly it should send the code on seppukuCh soon
	} else {
		go func() {
			var code int
			if Callback != nil {
				code = Callback(sig)
			}
			seppukuCh <- code
		}()
	}

	var code int
	select {
	case code = <-seppukuCh:
		// nothing (handled below)
	case <-timeoutCh:
		TimeoutLog()
		code = TimeoutExitCode
	}
	PostLog()
	os.Exit(code)
}

func initLogPrefixes() {
	if Betrayer == "" {
		Betrayer = DefaultBetrayer
	}
	if Betrayed == "" {
		Betrayed = filepath.Base(os.Args[0])
	}
	if betrayerPrefix == "" || betrayedPrefix == "" {
		betrayerPrefix = "[" + Betrayer + "] "
		betrayedPrefix = "[" + Betrayed + "] "
		for len(betrayerPrefix) < len(betrayedPrefix) {
			betrayerPrefix += " "
		}
		for len(betrayedPrefix) < len(betrayerPrefix) {
			betrayedPrefix += " "
		}
	}
}
