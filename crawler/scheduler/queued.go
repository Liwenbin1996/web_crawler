/*
 * @Author: wenbin
 * @Date: 2021-11-01 00:25:25
 * @LastEditTime: 2021-11-01 00:45:09
 * @Description:
 * @FilePath: /crawler/crawler/scheduler/queued.go
 */
package scheduler

import "crawler/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) Init() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	var requestQueue []engine.Request
	var workerQueue []chan engine.Request

	go func() {
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}
			select {
			case r := <-q.requestChan:
				requestQueue = append(requestQueue, r)
			case w := <- q.workerChan:
				workerQueue = append(workerQueue, w)
			case activeWorker <- activeRequest:
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()
}

func (q *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) Submit(req engine.Request) {
	q.requestChan <- req
}

func (q *QueuedScheduler) WorkerReady(worker chan engine.Request) {
	q.workerChan <- worker
}
