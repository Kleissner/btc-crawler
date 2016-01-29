package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Image struct {
	ID         int
	Seed       *Node
	StartedAt  time.Time
	FinishedAt time.Time
	Nodes      []*Node
	seen       map[string]bool
}

func (i *Image) Build() {
	jobs := make(chan *Node, 1000000)
	results := make(chan *Node, 100)
	i.seen = make(map[string]bool)

	seed := i.Seed

	jobs <- seed
	i.seen[seed.Address] = true

	i.StartedAt = time.Now()

	for i := 0; i < 200; i++ {
		go searcher(jobs, results)
	}

	count := 0
	for {
		select {
		case node := <-results:
			count++
			for _, neighbour := range node.Neighbours() {
				if !i.seen[neighbour.Address] {
					i.seen[neighbour.Address] = true
					jobs <- neighbour
				}
			}
			fmt.Println("Jobs length:", len(jobs))
			fmt.Println("Results length:", len(results))
			fmt.Println("Processed:", count)
		case <-time.After(20 * time.Second):
			close(jobs)
			close(results)
			i.FinishedAt = time.Now()
			return
		}
	}
}

func NewImage(seed *Node) *Image {
	i := new(Image)
	i.Seed = seed
	return i
}