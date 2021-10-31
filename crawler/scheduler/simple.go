/*
 * @Author: wenbin
 * @Date: 2021-11-01 00:09:36
 * @LastEditTime: 2021-11-01 00:23:01
 * @Description:
 * @FilePath: /crawler/crawler/scheduler/simple.go
 */
package scheduler

import "crawler/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Init() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(req engine.Request) {
	go func() { s.workerChan <- req }()
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}
