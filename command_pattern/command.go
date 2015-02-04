package command_pattern

import (
	"sync"
)

type Command interface {
	Execute()
}


type Cpattern struct {
	mu sync.Mutex
	commands []Command
}

func NewCP() *Cpattern {
	return &Cpattern{}
}


func (self *Cpattern) SetCommand(command Command) {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.commands = append(self.commands, command)
}

func (self *Cpattern) DelCommand(command Command, offset int, num int) {
	switch {
	case offset <= 0:
		offset = 1
		fallthrough
	case num <= 0:
		return
	case offset > len(self.commands):
		return
	}
	for k, com := range self.commands {
		if com == command {
			if k >= offset {
				self.mu.Lock()
				self.commands = append(self.commands[:k], self.commands[k+1:]...)
				self.mu.Unlock()
				num -= 1
				if num <= 0 {
					return
				}
			}
		}
	}
}

func (self *Cpattern) Exe() {
	for _, com := range self.commands {
		com.Execute()
	}
}
