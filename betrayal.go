package betrayal

import (
	"log"
	"os"
	"syscall"
	"time"
)

const DefaultBetrayer = "jigish"

var Timeout = 5 * time.Second
var TimeoutExitCode = 1
var Logger = log.Printf
var Callback func(os.Signal) int
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

func Wait(betrayalCh chan os.Signal, signals ...os.Signal) {
	sigCh := make(chan os.Signal)
	go WaitForYourSuddenButInevitableBetrayal(sigCh)
	signal.Notify(sigCh, signals...)
}

func WaitForYourSuddenButInevitableBetrayal(sigCh chan os.Signal) {
	signal := <-sigCh
	PreLog()
	seppukuCh := make(chan int)
	timeoutCh := time.After(Timeout)
	go func() {
		var code int
		if Callback != nil {
			code = Callback()
		}
		seppukuCh <- code
	}()
	var code int
	select {
	case code = <-seppukuCh:
		// nothing (handled below)
	case timeoutCh:
		KillLog()
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
		betrayerPrefix = "[" + betrayer + "] "
		betrayedPrefix = "[" + betrayed + "] "
		for len(betrayer) < len(betrayed) {
			betrayer += " "
		}
		for len(betrayed) < len(betrayer) {
			betrayed += " "
		}
	}
}
