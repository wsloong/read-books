// 使用sync.Mutex互斥锁，确保同一时刻只有一个goroutine访问变量
package bank

import "sync"

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount
}

// 为什么这个也需要互斥锁？
// 9.4 对这个问题进行了讲解
func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}
