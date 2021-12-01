package multi_thread

import (
	"fmt"
	"sync"
	"time"
)

//v4 is not good as v3,
//v4 allow at most certain thread at the same time, maybe they are same resource requester, in this case, it is not parallel.

type Semaphore struct {
	permits int      // 许可数量
	channel chan int // 通道
}

/* 创建信号量 */
func NewSemaphore(permits int) *Semaphore {
	return &Semaphore{channel: make(chan int, permits), permits: permits}
}

/* 获取许可 */
func (s *Semaphore) Acquire() {
	s.channel <- 0
}

/* 释放许可 */
func (s *Semaphore) Release() {
	<-s.channel
}

func The_dining_philosophers_v4(hungryPhis []int) {
	//init forks
	initForks()
	semaphore := NewSemaphore(len(hungryPhis) - 2)
	//start eating
	startTime := time.Now()
	wg := &sync.WaitGroup{}
	wg.Add(len(hungryPhis))
	for _, p := range hungryPhis {
		semaphore.Acquire() // before go to a new thread, check whether have slot
		go diningV4(wg, semaphore, p)
	}
	wg.Wait()
	span := time.Now().Sub(startTime)
	fmt.Println("time cost: ", span)
}

func diningV4(wg *sync.WaitGroup, semaphore *Semaphore, phi int) {
	defer wg.Done()
	defer semaphore.Release()
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
