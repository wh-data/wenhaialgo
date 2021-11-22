package multi_thread

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLocker(t *testing.T) {
	fmt.Println("start")
	var a sync.Mutex
	a.Lock()
	go func() {
		time.Sleep(time.Second * 4)
		fmt.Println("unlock")
		a.Unlock()
	}()
	a.Lock() //cannot lock util it is unlocked
	fmt.Println("lock again")
}

func Test_the_dining_philosophers_v1(t *testing.T) {
	hungryPhis := []int{0, 1, 2, 3, 4}
	count := 1000
	for count > 0 {
		The_dining_philosophers_v1(hungryPhis)
		count--
	}
}

func Test_the_dining_philosophers_v2(t *testing.T) {
	hungryPhis := []int{0, 1, 2, 3, 4}
	count := 10
	for count > 0 {
		The_dining_philosophers_v2(hungryPhis)
		count--
	}
}
