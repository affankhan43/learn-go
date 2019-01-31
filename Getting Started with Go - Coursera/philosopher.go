/** 
WEEK 4 - Concurrency in GO
**/

package main

import(
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type ChopSt struct {
	sync.Mutex
}

type Philo struct {
	philo_no int
	leftCS, rightCS *ChopSt
}

func (ph Philo) eatting(){
	ph.leftCS.Lock()
	ph.rightCS.Lock()
	fmt.Printf("starting to eat %d\n", ph.philo_no)
	fmt.Printf("finishing eating %d\n", ph.philo_no)
	ph.rightCS.Unlock()
	ph.leftCS.Unlock()
	wg.Done()
}

func host(philos []*Philo) {
	philospher_nums := make([]int, 0,0)
	for i :=0; i < 5 ;i++ {
		philospher_nums = append(philospher_nums, i) // 0 , 1 , 2, 3, 4
		philospher_nums = append(philospher_nums, (i + 2) % 5) // 2, 3, 4, 0, 1
		philospher_nums = append(philospher_nums, (i + 4) % 5) // 4, 0, 1, 2, 3
	}
	//philospher_nums [0 2 4 1 3 0 2 4 1 3 0 2 4 1 3]
	for len(philospher_nums) > 0 {
		ph_no1 := philospher_nums[0] // 0 ,4 ,3 ,2 ,1 ,0 ,4
		if len(philospher_nums) == 1 {
			wg.Add(1)
			philospher_nums = nil
			go philos[ph_no1].eatting()
		} else {
			wg.Add(2)
			ph_no2 := philospher_nums[1] // 2 ,1 ,0 ,4 ,3, 2 ,1
			philospher_nums = philospher_nums[2:] // [len 13] ,[len 11], [len 9] ,[len 7], [len 5], [len 3], [len 1]
			go philos[ph_no1].eatting()
			go philos[ph_no2].eatting()
		}
		wg.Wait()
	}
}

func main(){
	ChopSticks := make([]*ChopSt,5)
	for i := 0; i < 5; i++ {
		ChopSticks[i] = new(ChopSt)
	}
	philos := make([]*Philo,5)
	for i := 0; i<5 ; i++ {
		philos[i] = &Philo { i, ChopSticks[i],ChopSticks[(i + 1) % 5] }
	}
	host(philos)
}

