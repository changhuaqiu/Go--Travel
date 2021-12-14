package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

//定义的错误语句
//errors.New return errorString{text}对象
var ErrTimeOut = errors.New("执行者执行超时")
var ErrInterrupt = errors.New("执行者被中断")

type Runner struct {
	tasks     []func(int)      //任务列表
	complete  chan error       //完成判断 分两种情况
	timeout   <-chan time.Time //超时判断
	interrupt chan os.Signal   //信号
}

func New(t time.Duration) *Runner {
	return &Runner{
		complete:  make(chan error),
		timeout:   time.After(t),           //time.After 返回 <-chan time.Time 对象
		interrupt: make(chan os.Signal, 1), //信号
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.isInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}

//测试
func main() {
	log.Printf("...开始执行任务...")
	timeout := 3 * time.Second
	r := New(timeout)
	r.Add(createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeOut:
			log.Println(err)
			os.Exit(1)
		case ErrInterrupt:
			log.Println(err)
			os.Exit(1)
		}
	}
	log.Println("...任务执行结束")
}

//测试添加任务
func createTask() func(int) {
	return func(id int) {
		log.Printf("正在执行任务%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
