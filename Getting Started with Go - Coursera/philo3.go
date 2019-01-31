package main

import (
  "fmt"
  "sync"
)

var wg sync.WaitGroup

type ChopS struct {
  sync.Mutex
}

type Philo struct {
  name  int
  leftCS, rightCS *ChopS
}

func (p Philo) eat() {
  p.leftCS.Lock()
  p.rightCS.Lock()
  fmt.Printf("starting to eat %d\n", p.name)
  fmt.Printf("finishing eating %d\n", p.name)
  p.rightCS.Unlock()
  p.leftCS.Unlock()
  wg.Done()
}

func sum( a [5]int) int {
  var t int
  for i,_ := range a {
   t +=a[i]
  }
  return t
}

func host(philos []*Philo) {
  portions := make([]int, 0,0)
  for i :=0; i < 5 ;i++ {
    portions = append(portions, i) // 0 , 1 , 2, 3, 4
    portions = append(portions, (i + 2) % 5) // 2, 3, 4, 0, 1
    portions = append(portions, (i + 4) % 5) // 4, 0, 1, 2, 3
   }
  for len(portions) > 0 {
    pn1 := portions[0]
    if len(portions) == 1 {
      wg.Add(1)
      portions = nil
      go philos[pn1].eat()
    } else {
      wg.Add(2)
      pn2 := portions[1]
      portions = portions[2:]
      go philos[pn1].eat()
      go philos[pn2].eat()
    }
    wg.Wait()
  }
}

func main() {

  CSticks := make([]*ChopS,5)
  for i := 0; i < 5; i++ {
  CSticks[i] = new(ChopS)
  }
  philos := make([]*Philo,5) 
  for i := 0; i<5 ; i++ {
    philos[i] = &Philo { i, CSticks[i],
                  CSticks[(i + 1) % 5] }
  }
  host(philos)
}