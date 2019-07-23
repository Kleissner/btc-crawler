// This package describes the Bitcoin network crawler.
// The program's main entry point is in main.go, although the logic to dictate the running order of the crawler and listener is located in dispatcher.go
package main

import (
	_ "net/http/pprof"
)

func main() {
	dispatcher := NewDispatcher()
	_ = dispatcher.BuildImage(49000)

	return
}
