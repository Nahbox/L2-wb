package commandpattern

import "fmt"

// Интерфейс Команды
type Command interface {
	Execute()
}

// Конкретная команда - включить свет
type LightOnOffCommand struct {
	Light *Light
}

func (c *LightOnOffCommand) Execute() {
	c.Light.TurnOnOff()
}

// Получатель команды - свет
type Light struct {
	isOn bool
}

func (l *Light) TurnOnOff() {
	if l.isOn {
		l.isOn = false
		fmt.Println("Свет выключен")
	} else {
		l.isOn = true
		fmt.Println("Свет включен")
	}
}

// Инициатор - пульт управления
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

// CommandPattern точка входа
func CommandPattern() {
	// Создаем объект света
	light := &Light{}

	// Создаем команды
	lightOnCommand := &LightOnOffCommand{Light: light}

	// Создаем пульт управления
	remoteControl := &RemoteControl{lightOnCommand}

	// Используем пульт для включения света
	remoteControl.PressButton()

	// Используем пульт для выключения света
	remoteControl.PressButton()
}
