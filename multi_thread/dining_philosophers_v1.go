package multi_thread

import (
	"fmt"
	"sync"
	"time"
)

//this method will have dead lock issue. e.g all phi take left fork
func The_dining_philosophers_v1(hungryPhis []int) {
	//init forks
	initForks()
	//start eating
	startTime := time.Now()
	wg := &sync.WaitGroup{}
	wg.Add(len(hungryPhis))
	for _, p := range hungryPhis {
		go dining(wg, p)
	}
	wg.Wait()
	span := time.Now().Sub(startTime)
	fmt.Println("time cost: ", span)
}

func dining(wg *sync.WaitGroup, phi int) {
	defer wg.Done()
	//make forks ready
	lFork, lReady := pickLeftFork(phi)
	rFork, rReady := pickRightFork(phi)
	if lReady && rReady {
		time.Sleep(time.Microsecond * UnitDiningTime) //mock eating time
		Forks[lFork].Locker.Unlock()
		Forks[rFork].Locker.Unlock()
	}
}
