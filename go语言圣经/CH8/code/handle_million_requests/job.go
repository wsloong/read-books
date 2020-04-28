package main

import (
	"fmt"
	"log"
)

type Payload int

// UploadToS3 模拟任务
func (p Payload) UploadToS3() error {
	fmt.Println("doing job:", p)
	return nil
}

// PayloadCollection request发过来的数据
type PayloadCollection struct {
	WindowsVersion string    `json:"version"`
	Token          string    `json:"token"`
	Payloads       []Payload `json:"data"`
}

// Job 要运行的任务
type Job struct {
	Playload Payload
}

// JobQueue 任务的通道
// 这里作者文章中并没有初始化,所以在 PayloadHandler 中会阻塞
var JobQueue = make(chan Job)

// Worker 执行任务的worker
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

// NewWork 实例 Worker
func NewWork(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// Start 启动Worker
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				// 接收到一个job
				if err := job.Playload.UploadToS3(); err != nil {
					log.Printf("error uploading to s3:%s", err)
				}
			case <-w.quit:
				// 接收到了退出的信号
				return
			}
		}
	}()
}

// Stop 停止Worker
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
