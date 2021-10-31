/*
 * @Author: wenbin
 * @Date: 2021-10-31 21:19:35
 * @LastEditTime: 2021-11-01 00:22:11
 * @Description:
 * @FilePath: /crawler/crawler/engine/concurrent.go
 */
package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Init()
	Submit(Request)
	WorkerReady(chan Request)
	WorkerChan() chan Request
}

func (c *ConcurrentEngine) createOneWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			c.Scheduler.WorkerReady(in)
			req := <-in
			results, err := worker(req)
			if err != nil {
				continue
			}
			out <- results
		}
	}()
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	// 1. 初始化scheduler
	c.Scheduler.Init()

	// 2. 初始化worker
	out := make(chan ParseResult)
	for i := 0; i < c.WorkerCount; i++ {
		in := c.Scheduler.WorkerChan()
		c.createOneWorker(in, out)
	}

	// 3. 下发任务到scheduler
	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}
	// 4. 获取执行结果并处理
	for {
		results := <-out
		resultHandle(results.Items)

		for _, r := range results.Requests {
			c.Scheduler.Submit(r)
		}
	}
}
