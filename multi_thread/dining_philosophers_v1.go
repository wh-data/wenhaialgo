package multi_thread

import (
	"fmt"
	"sync"
	"time"
)

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
		time.Sleep(time.Microsecond * UnitDiningTime) //eating
		Forks[lFork].Locker.Unlock()
		Forks[rFork].Locker.Unlock()
	}
}

//todo: address dead lock issue. e.g all phi take left fork
func pickLeftFork(phi int) (int, bool) {
	var lFork int
	//check fork num
	if phi == 0 {
		lFork = len(Philosophers) - 1
	} else {
		lFork = phi - 1
	}
	//check fork lock
	fmt.Println("phi", phi, "lFork", Forks[lFork].Locker)
	Forks[lFork].Locker.Lock()
	return lFork, true
}

func pickRightFork(phi int) (int, bool) {
	var rFork int
	//check fork num
	rFork = phi
	//check fork lock
	Forks[rFork].Locker.Lock()
	return rFork, true
}
