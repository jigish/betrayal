package betrayal

import (
	"os"
	"testing"
)

func TestBetrayal(t *testing.T) {
	Daemon = func(sigCh chan os.Signal, dieCh chan int) {
		sig := <-sigCh
		if sig == os.Interrupt {
			dieCh <- 0
		} else {
			dieCh <- 1
		}
	}

	sigCh := make(chan os.Signal)
	go func() {
		sigCh <- os.Interrupt
	}()
	if code := Test(sigCh); code != 0 {
		t.Fatal("should have been 0")
	}

	go func() {
		sigCh <- os.Kill
	}()
	if code := Test(sigCh); code != 1 {
		t.Fatal("should have been 1")
	}
}
