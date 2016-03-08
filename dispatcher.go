package crawler

import (
	"time"
)

type Dispatcher struct {
}

func (d *Dispatcher) BuildImage(workers int) *Image {
	crawler := NewCrawler(workers)

	crawler.Start()
	image := <-crawler.Done

	listener := NewListener(image, 24*time.Hour)
	go listener.Listen()
	<-listener.DoneC

	return image
}

func NewDispatcher() *Dispatcher {
	d := new(Dispatcher)

	return d
}
