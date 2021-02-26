package TaskQueue

import (
	"sync"
	"sync/atomic"
)

type TaskFun func()
type taskQueue struct {
	taskList []TaskFun
	taskFront int64
	taskTail int64
	taskCount int64
	taskMaxCount int64

}

func NewTaskQueue(taskMaxCount int64) *taskQueue {
	return &taskQueue{taskCount: 0,taskMaxCount: taskMaxCount,taskList: make([]TaskFun,taskMaxCount)}
}
func (queue *taskQueue) isEmpty()bool  {
	 return atomic.CompareAndSwapInt64(&queue.taskCount,0,0)
}

func (queue *taskQueue) isFull()bool  {
	return atomic.CompareAndSwapInt64(&queue.taskCount,queue.taskMaxCount,queue.taskMaxCount)
}

func (queue *taskQueue) push(fun TaskFun)  {
     queue.taskCount ++;
     temp := queue.taskTail
	 queue.taskTail++
	 queue.taskList[temp%queue.taskMaxCount] = fun
}

func (queue *taskQueue) pop() TaskFun {

	queue.taskCount --;
	temp:=queue.taskFront
	queue.taskFront++
	return queue.taskList[temp%queue.taskMaxCount]
}

type ThreadPoolInterface interface {
	AddTask(fun TaskFun)
	Start()
	End()
}

func  NewThreadTool(maxTaskQueueCount int64) ThreadPoolInterface{
	temp := ThreadPool{}
	temp.exitCount.Store( 0)
	temp.cond = sync.Cond{L: &temp.locker}
	temp.stop = false
	temp.mTaskQueue = NewTaskQueue(maxTaskQueueCount)
	temp.Start()
	return &temp
}
type ThreadPool struct {
	mTaskQueue *taskQueue
	exitCount atomic.Value
	stop bool
	cond sync.Cond
	locker sync.Mutex
}

func (t* ThreadPool) AddTask(fun TaskFun) {
	t.locker.Lock()
	for t.mTaskQueue.isFull() {
		t.cond.Wait()
	}
	t.mTaskQueue.push(fun)
	t.locker.Unlock()
	t.cond.Signal()

}

func (t* ThreadPool) Start() {
	for i:=(int64)(0);i<t.mTaskQueue.taskMaxCount;i++{
		go func() {
			for {
				t.locker.Lock()
				for t.mTaskQueue.isEmpty() {
					if t.stop{
						temp :=t.exitCount.Load().(int)
						temp ++;
						t.exitCount.Store(temp)
						return
					}
					t.cond.Wait()
				}
				t.mTaskQueue.pop()()
				t.locker.Unlock()
				t.cond.Signal()
				if t.stop{
					temp :=t.exitCount.Load().(int)
					temp ++;
					t.exitCount.Store(temp)
					return
				}
			}
		}()
	}
}

func (t* ThreadPool) End() {
	t.stop = true
	for t.exitCount.Load() != t.mTaskQueue.taskMaxCount{

	}
}

