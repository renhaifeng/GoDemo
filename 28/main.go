package main

import (
	"fmt"
	"sync"
	"time"
)

type Money struct {
	lock   sync.Mutex
	amount int64
}

// Add 加钱
func (m *Money) Add(i int64) {
	m.amount = m.amount + i
}

// Sub 减钱
func (m *Money) Sub(i int64) {
	m.lock.Lock()
	defer m.lock.Unlock()

	// 钱足才能减
	if m.amount >= i {
		m.amount = m.amount - i
	}
}

// Get 查看还有多少钱
func (m *Money) Get() int64 {
	return m.amount
}

func main() {
	m := new(Money)
	m.Add(10000)
	var wg sync.WaitGroup

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			time.Sleep(500 * time.Millisecond)
			m.Sub(5)
			fmt.Println(m.Get())
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("结束：", m.Get())

}
