//GO里面MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全，且需要实现以下接口：
package main

import (
	"fmt"
	"sync"
	"time"
)

type sp interface {
	Out(key string, val interface{})                  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type Map struct {
	c   map[string]*entry
	rmx *sync.RWMutex
}
type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

//
func (m *Map) Out(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()
	item, ok := m.c[key]
	if !ok {
		m.c[key] = &entry{
			ch:      make(chan struct{}),
			value:   val,
			isExist: true,
		}
		close(m.c[key].ch)
		return
	}
	item.value = val
	item.isExist = true
	close(item.ch)
	return
}

//1.key存在 直接返回
//2.key不存在 等待
//
func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	m.rmx.Lock()

	if item, ok := m.c[key]; ok && item.isExist {
		m.Unlock()
		return item.value
	} else if !ok {
		item = &entry{
			ch:      make(chan struct{}),
			isExist: false,
		}
		m.c[key] = item
		m.rmt.Unlock()
		fmt.Println("协程阻塞 -> ", key)
		select {
		case <-item.ch:
			return item.value
		case <-time.After(timeout):
			fmt.Println("协程超时 -> ", key)
			return nil
		}
	} else {
		m.rmx.Unlock()
		fmt.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			fmt.Println("协程超时 -> ", key)
			return nil
		}
	}
}
