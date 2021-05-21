package new_and_old_models_task

import (
	"fmt"
	"sync"
	"time"
)

/*
	使用锁还是通道的普遍经验法则：
	* 使用锁的场景
		* 访问共享数据结构中的缓存信息
		* 保存应用程序上下文和状态信息数据
	* 使用通道的情景
		* 与异步操作的结果进行交互
		* 分发任务
		* 传递数据所有权
 */

var wg sync.WaitGroup

type Task struct {
	id int
	name string
}

type Pool struct {
	Mu sync.Mutex
	Tasks []Task

}

func Worker()  {
	
}

func Doworker()  {

	task1 := Task{1, "jack"}
	task2 := Task{2, "mary"}
	task1.name = "jack2"

	p := &Pool{Tasks: []Task{task1, task2}}
	fmt.Println(p)

	wg.Add(2)
	// 将Tasks中所有的name都加上 班级：
	go func() {

		defer wg.Done()
		for k, task := range p.Tasks {
			p.Mu.Lock()

			task.name = "a"
			p.Tasks[k] = task

			p.Mu.Unlock()
		}
	}()

	// 将Tasks中所有的name都加上 班级：
	go func() {
		time.Sleep(1e9)
		defer wg.Done()

		for k, task := range p.Tasks {

			p.Mu.Lock()

			task.name = "班级：" + task.name
			p.Tasks[k] = task

			p.Mu.Unlock()
		}
	}()



	wg.Wait()

	fmt.Println(p)
}

