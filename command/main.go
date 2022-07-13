package main

import "fmt"

//Интерфейс команды
type command interface {
	execute()
}

//Отправитель
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

//Добавить в команду
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

//Исключить из команды
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

//Интерфейс получателя
type device interface {
	on()
	off()
}

//Конкретный получатель
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &tv{}

	onCommand := &onCommand{
		device: tv,
	}

	offCommand := &offCommand{
		device: tv,
	}

	onButton := &button{
		command: onCommand,
	}
	onButton.press()

	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
