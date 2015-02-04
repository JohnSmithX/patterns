package main

import (
	"fmt"
	c "github.com/JohnSmithX/patterns/command_pattern"
)

type Light struct {
	light bool
}

func (self *Light) on() {
	//do something
	self.light = true
	fmt.Println("light on")
}

type LightOnCommand struct {
	light *Light
}

//constructor for create a LightOnCommand instance
func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{
		light: light,
	}
}

func (self *LightOnCommand) Execute() {
	self.light.on()
}

func main() {
	cp := c.NewCP()
	loc := NewLightOnCommand(&Light{})
	cp.SetCommand(loc)
	cp.Exe()
}
