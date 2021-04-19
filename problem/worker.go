package problem

/*
问题：
给定k个worker以及n个job，其中job会以放入队列的时间putInQueueAt排列，同一时间一个
worker只能执行一个任务，每个任务的耗时为duration，求出每个job的执行workerID以及
执行完成时间。
要求：当多个worker同时空闲时，选择编号最小的offer执行
*/

type job struct {
	ID int
	// time to put it into queue
	putInQueueAt int
	// time that the job will consume
	duration int
}

type worker struct {
	ID int
	// time that work complete
	finishedAt int
}

type res struct {
	jobID         int
	workerID      int
	jobCompleteAt int
}

// O(nlogk)
func getJobScheduleInfo(jobs []job, k int) []res {
	if len(jobs) == 0 || k == 0 {
		return nil
	}
	ret := make([]res, 0, len(jobs))
	// 使用free和busy两个堆来维护顺序
	// 根据id最小原则维护空闲worker堆
	// 根据执行完成时间最小原则维护工作中worker堆
	free := make([]worker, k)
	busy := make([]worker, 0, k)
	for i := range free {
		free[i] = worker{ID: i + 1}
	}
	for _, j := range jobs {
		// 把在当前时间内busy中空闲的worker放到free中
		for len(busy) > 0 && busy[0].finishedAt <= j.putInQueueAt {
			busy[0], busy[len(busy)-1] = busy[len(busy)-1], busy[0]
			last := busy[len(busy)-1]
			busy = busy[:len(busy)-1]
			down(busy, 0, compareByFinishTime)
			free = append(free, last)
			up(free, len(free)-1, compareByID)
		}

		activeWorker := worker{}
		// 若空闲worker为0，则需要等待busy中最快完成任务的worker执行完毕后再执行当前任务
		if len(free) == 0 {
			activeWorker = busy[0]
			busy[0], busy[len(busy)-1] = busy[len(busy)-1], busy[0]
			busy = busy[:len(busy)-1]
			down(busy, 0, compareByFinishTime)
			activeWorker.finishedAt += j.duration
		} else {
			activeWorker = free[0]
			free[0], free[len(free)-1] = free[len(free)-1], free[0]
			free = free[:len(free)-1]
			down(free, 0, compareByID)
			activeWorker.finishedAt = j.putInQueueAt + j.duration
		}

		// 把执行当前任务的worker放到busy堆中
		busy = append(busy, activeWorker)
		up(busy, len(busy)-1, compareByFinishTime)

		ret = append(ret, res{
			jobID:         j.ID,
			workerID:      activeWorker.ID,
			jobCompleteAt: activeWorker.finishedAt,
		})
	}
	return ret
}

func compareByID(a, b worker) bool {
	return a.ID < b.ID
}

func compareByFinishTime(a, b worker) bool {
	return a.finishedAt < b.finishedAt
}

func down(data []worker, x int, compare func(worker, worker) bool) {
	for i := x; i*2+1 < len(data); {
		flag := i*2 + 1
		if flag+1 < len(data) && compare(data[flag+1], data[flag]) {
			flag++
		}
		if compare(data[flag], data[i]) {
			data[flag], data[i] = data[i], data[flag]
			i = flag
		} else {
			break
		}
	}
}

func up(data []worker, x int, compare func(worker, worker) bool) {
	for i := x; (i-1)/2 >= 0 && compare(data[i], data[(i-1)/2]); i = (i - 1) / 2 {
		data[i], data[(i-1)/2] = data[(i-1)/2], data[i]
	}
}
