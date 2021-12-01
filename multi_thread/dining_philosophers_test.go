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
	The_dining_philosophers_v1(hungryPhis)
}

func Test_the_dining_philosophers_v2(t *testing.T) {
	count := 10
	hungryPhis := make([]int, count)
	for i := 0; i < count; i++ {
		hungryPhis[i] = rand.Intn(5)
	}
	fmt.Println("hungryPhis", hungryPhis)
	The_dining_philosophers_v2(hungryPhis)

}

func Test_the_dining_philosophers_v3(t *testing.T) {
	count := 100
	hungryPhis := make([]int, count)
	for i := 0; i < count; i++ {
		hungryPhis[i] = rand.Intn(5)
	}
	fmt.Println("hungryPhis", hungryPhis)
	The_dining_philosophers_v3(hungryPhis)
}

//not good as v3
func Test_the_dining_philosophers_v4(t *testing.T) {
	count := 100
	hungryPhis := make([]int, count)
	for i := 0; i < count; i++ {
		hungryPhis[i] = rand.Intn(5)
	}
	fmt.Println("hungryPhis", hungryPhis)
	The_dining_philosophers_v4(hungryPhis)
}
