package lock

import "sync"

/*
	竞争条件
		* 锁
			一个函数在并发调用时没法工作的原因太多了，比如死锁(deadlock)，活锁(livelock)、饿死resource starvation
 */

// 惯例，被mutex所保护的变量是在mutex变量声明之后立刻声明的(如下)。如果你的做法和惯例不符，确保在文档里对你的做法进行说明
var (
	mu sync.Mutex
	RWmu sync.RWMutex
	balance int
)
/*

// version1
func Deposit(amout int)  {
	mu.Lock()
	defer mu.Unlock()

	balance += amout
}


func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	b := balance
	return b
}*/

// 取款功能
// 下面代码报错。会产生死锁。因为Deposit()会调用mu.Lock第二次去获取互斥锁，但因为mutex在Withdraw()已经锁上，无法被重入
// 也就是说没法已经锁上的mutext再次上锁
// 解决上面的方案：将一个函数分离为多个函数，比如我们把Deposit分离成两个：比如我们把Deposit分离成两个：一个不导出的函数deposit，这个函数假设
// 锁总是会被保持并去做实际的操作，另一个导出的函数Deposit，这个函数会调用deposit，但在调用前会先去获取锁
/*func Withdraw(amount int) bool  {
	mu.Lock()
	defer mu.Unlock()

	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false
	}

	return true
}*/

func deposit(amount int)  {
	balance += amount
}

func Deposit(amount int)  {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}

	return true
}

func Balance() int {
	RWmu.RLock()
	defer RWmu.RUnlock()
	return balance
}