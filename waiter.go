package waiter

import (
	"os"
	"os/signal"
	"sync"
)

// WaitForSignal accepts an os.Signal and blocks until that signal is recieved
func WaitForSignal(sig os.Signal) {
	var waiter sync.WaitGroup
	waiter.Add(1)
	var signalChannel chan os.Signal
	signalChannel = make(chan os.Signal, 1)
	signal.Notify(signalChannel, sig)
	go func() {
		<-signalChannel
		waiter.Done()
	}()
	waiter.Wait()
}
