package multi_thread

import (
	"fmt"
	"sync"
	"time"
)

//v2 need to retry and there is possibility that retry fail, and it can hardly parallel run
//v3 no need retry
//solution: use lock wait directly; how to avoid dead lock: some phi pick left first, others pick right fist
func The_dining_philosophers_v3(hungryPhis []int) {
	//init forks
	initForks()
	//start eating
	startTime := time.Now()
	wg := &sync.WaitGroup{}
	wg.Add(len(hungryPhis))
	for _, p := range hungryPhis {
		go diningV3(wg, p)
	}
	wg.Wait()
	span := time.Now().Sub(startTime)
	fmt.Println("time cost: ", span)
}

func diningV3(wg *sync.WaitGroup, phi int) {
	defer wg.Done()

	//make sure half take left first, another half take right first
	lFork, rFork := -1, -1
	if phi/2 == 0 {
		lFork, _ = pickLeftFork(phi)
		rFork, _ = pickRightFork(phi)
	} else {
		rFork, _ = pickRightFork(phi)
		lFork, _ = pickLeftFork(phi)
	}
	time.Sleep(time.Millisecond * UnitDiningTime) //mock eating time
	//after eating, release forks
	Forks[lFork].Locker.Unlock()
	Forks[rFork].Locker.Unlock()
}
