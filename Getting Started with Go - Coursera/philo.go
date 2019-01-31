package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var dining sync.WaitGroup

type signal struct{}

type Table struct {
	Host         chan bool
	Forks        [5]sync.Mutex
	Philosophers []*Philosopher
}

type Philosopher struct {
	Name        string
	Left, Right int
	Table       *Table
}

func (p *Philosopher) TakeForks() {
	p.Table.Forks[p.Left].Lock()
	p.Table.Forks[p.Right].Lock()
  fmt.Println("starting to eat ", p.Name)
}

func (p *Philosopher) PutForks() {
	p.Table.Forks[p.Left].Unlock()
	p.Table.Forks[p.Right].Unlock()
  fmt.Println("finishing eating ", p.Name)
	<-p.Table.Host
}

func (p *Philosopher) Think() {
	time.Sleep(50*time.Millisecond + time.Duration(rand.Intn(100)))
}

func (p *Philosopher) Eat() {
	time.Sleep(50*time.Millisecond + time.Duration(rand.Intn(100)))
}

func (p *Philosopher) AskPermission() {
  p.Table.Host <- true
}

func (p *Philosopher) Live() {
  for hunger := 3; hunger > 0; hunger-- {
		p.Think()
    p.AskPermission()
		p.TakeForks()
		p.Eat()
		p.PutForks()
	}
  dining.Done()
}

func main() {
  dining.Add(5)
	table := &Table{}
	table.Host = make(chan bool, 2)
	table.Philosophers = []*Philosopher{
		{"1", 0, 1, table},
		{"2", 1, 2, table},
		{"3", 2, 3, table},
		{"4", 3, 4, table},
		{"5", 4, 0, table},
	}

	for _, p := range table.Philosophers {
		go p.Live()
	}
  dining.Wait()
}
