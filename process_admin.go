package main

import "fmt"

type ProcessAdmin struct {
	Procecess     []*Process
	ProcessLength uint64
}

func (self *ProcessAdmin) AddProcess(process *Process) {
	self.Procecess = append(self.Procecess, process)
	self.ProcessLength += 1
}

func (self *ProcessAdmin) KillProcess(processId uint64) bool {
	newProcecess := []*Process{}
	deleted := false

	for _, process := range self.Procecess {
		if process.ID != processId {
			newProcecess = append(newProcecess, process)
		}

		if process.ID == processId {
			deleted = true
			process.Stop()
			self.ProcessLength -= 1
		}
	}

	self.Procecess = newProcecess

	return deleted
}

func (self *ProcessAdmin) ShowProcecess() {
	for _, process := range self.Procecess {
		fmt.Printf("ID %d: %d \n", process.ID, process.TimeUp)
	}
}

func (self *ProcessAdmin) KillAllProcecess() {
	for _, process := range self.Procecess {
		process.Stop()
	}
}
