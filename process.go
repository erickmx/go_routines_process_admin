package main

import (
	"time"
)

type Process struct {
	ID        uint64
	TimeUp    uint64
	IsRunning bool
}

func (self *Process) Start() {
	self.TimeUp = 0
	self.IsRunning = true

	for {
		self.TimeUp += 1
		time.Sleep(time.Millisecond * 500)
		if !self.IsRunning {
			break
		}
	}
}

func (self *Process) Stop() {
	self.IsRunning = false
}

func NewProcess(id uint64) *Process {
	return &Process{
		ID: id,
	}
}
