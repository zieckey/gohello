package main

import (
	"fmt"
	"github.com/kklis/gomemcache"
)

type Pool struct {
	stack chan *gomemcache.Memcache
	addr  string // the memcached server address
}

func New(addr string) *Pool {
	pool := new(Pool)
	StackSize := 128
	pool.addr = addr
	pool.stack = make(chan *gomemcache.Memcache, StackSize)
	for i := 0; i < StackSize; i++ {
		c, err := gomemcache.Dial(addr)
		if err != nil {
			fmt.Printf("Connect to %v failed\n", addr)
			return nil
		}
		pool.stack <- c
	}
	return pool
}

func (pool *Pool) Get() *gomemcache.Memcache {
	select {
	case conn := <-pool.stack:
		return conn
	default:
		conn, err := gomemcache.Dial(pool.addr)
		if err != nil {
			fmt.Printf("Connect to %v failed\n", pool.addr)
			return nil
		}
		return conn
	}
}

func (pool *Pool) Put(conn *gomemcache.Memcache) {
	select {
	case pool.stack <- conn:
	default:
	}
}
