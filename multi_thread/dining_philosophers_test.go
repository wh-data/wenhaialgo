package multi_thread

import (
	"fmt"
	"math/rand"
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
	count := 100
	hungryPhis := make([]int, count)
	for i := 0; i < count; i++ {
		hungryPhis[i] = rand.Intn(5)
	}
	fmt.Println(hungryPhis)
	for count > 0 {
		The_dining_philosophers_v1(hungryPhis)
		count--
	}
}

func Test_the_dining_philosophers_v2(t *testing.T) {
	count := 100
	hungryPhis := make([]int, count)
	for i := 0; i < count; i++ {
		hungryPhis[i] = rand.Intn(5)
	}
	for count > 0 {
		The_dining_philosophers_v2(hungryPhis)
		count--
	}
}
