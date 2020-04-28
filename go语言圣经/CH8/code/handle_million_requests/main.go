package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	MaxWorker = 10
	maxQueue  = 1000
	MaxLength = 2048
)

func main() {
	dispatcher := NewDispatcher(maxQueue)
	dispatcher.Run()

	http.HandleFunc("/", PayloadHandler)
	fmt.Println("Listen at 8090")
	http.ListenAndServe(":8090", nil)
}

func PayloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var content = &PayloadCollection{}
	// err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
	// 我们大部分时间是读取全部的body
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	///逐个检查每个payload和queue，以便将它们发送到S3
	for _, payload := range content.Payloads {
		// 创建一个任务，并放到queue
		job := Job{Playload: payload}
		JobQueue <- job
	}

	w.WriteHeader(http.StatusOK)
	return
}

type Dispatcher struct {
	WorkerPool chan chan Job
}

func NewDispatcher(maxQueue int) *Dispatcher {
	pool := make(chan chan Job, maxQueue)
	return &Dispatcher{WorkerPool: pool}
}

func (d *Dispatcher) Run() {
	for i := 0; i < MaxWorker; i++ {
		worker := NewWork(d.WorkerPool)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// 尝试获取一个可用的work pool, 该操作会阻塞
			jobChannel := <-d.WorkerPool
			jobChannel <- job
			// 接收到一个任务
			// go func(job Job) {
			// 	// 尝试获取一个可用的work pool, 该操作会阻塞
			// 	jobChannel := <-d.WorkerPool

			// 	jobChannel <-job
			// }(job)
		}
	}
}
