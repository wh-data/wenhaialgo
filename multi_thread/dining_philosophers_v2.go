package multi_thread

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

func The_dining_philosophers_v2(hungryPhis []int) {
	//init forks
	initForks()
	//start eating
	startTime := time.Now()
	wg := &sync.WaitGroup{}
	wg.Add(len(hungryPhis))
	for _, p := range hungryPhis {
		go diningV2(wg, p)
	}
	wg.Wait()
	span := time.Now().Sub(startTime)
	fmt.Println("time cost: ", span)
}

func diningV2(wg *sync.WaitGroup, phi int) {
	defer wg.Done()
	//make forks ready
	lFork, rFork := pickForks(phi, Retry)
	time.Sleep(time.Millisecond * UnitDiningTime) //mock eating time
	if lFork != -1 && rFork != -1 {
		Forks[lFork].Locker.Unlock()
		Forks[rFork].Locker.Unlock()
		fmt.Println("phi ", phi, "takes fork: ", lFork, rFork)
	} else {
		fmt.Println("phi ", phi, "failed to eat")
		fmt.Println(Forks[lFork].Locker)
		fmt.Println(Forks[rFork].Locker)
	}
}

//lock both left and right together to avoid dead lock
//add retry to reduce the possibility of no resource
//bas solution: even serial run (串行) is faster than this
func pickForks(phi, retry int) (int, int) {
	if retry < 0 {
		return -1, -1
	}
	var lFork, rFork int
	//check fork num
	if phi == 0 {
		lFork = len(Philosophers) - 1
	} else {
		lFork = phi - 1
	}
	rFork = phi
	//lock left and right together or do not lock
	lState := reflect.ValueOf(Forks[lFork].Locker).Elem().FieldByName("state")
	rState := reflect.ValueOf(Forks[rFork].Locker).Elem().FieldByName("state")
	if lState.Int() == 0 && rState.Int() == 0 {
		Forks[lFork].Locker.Lock()
		Forks[rFork].Locker.Lock()
	} else {
		time.Sleep(time.Second * 3) //bad solution, if retry time set not proper, it aways retry fail.
		return pickForks(phi, retry-1)
	}
	return lFork, rFork
}
