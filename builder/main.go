package main

import (
	"fmt"
)

/*
Паттерн «строитель».
*/

type Computer struct { // Строим копмьютер
	CPU string
	RAM int
	SSD int
}

type ComputerBuilderI interface {
	CPU(val string) ComputerBuilderI
	RAM(val int) ComputerBuilderI
	SSD(val int) ComputerBuilderI

	Build() Computer
}

type Builder struct {
	cpu string
	ram int
	ssd int
}

func NewComputerBuilder() ComputerBuilderI {
	return Builder{}
}

func (b Builder) CPU(val string) ComputerBuilderI {
	b.cpu = val
	return b
}
func (b Builder) RAM(val int) ComputerBuilderI {
	b.ram = val
	return b
}
func (b Builder) SSD(val int) ComputerBuilderI {
	b.ssd = val
	return b
}

func (b Builder) Build() Computer {
	return Computer{
		CPU: b.cpu,
		RAM: b.ram,
		SSD: b.ssd,
	}
}

func main() {
	fmt.Println("Часть первая builder")
	fmt.Println("")

	fmt.Println("Пустой строитель")
	Builder := NewComputerBuilder()
	fmt.Println(Builder.CPU("core 3").RAM(2).SSD(128).Build())
	fmt.Println("")

	fmt.Println("Строитель фронтендера")
	frontendComputer := NewFrontendComputerBuilder()
	fmt.Println(frontendComputer.Build())
	fmt.Println(frontendComputer.CPU("core i7").Build())
	fmt.Println("")

	fmt.Println("Строитель бэкэндера")
	compBuilder := NewBackendComputerBuilder()
	fmt.Println(compBuilder.Build())
	fmt.Println(compBuilder.CPU("core 3").RAM(2).SSD(128).Build())
	fmt.Println("")

	fmt.Println("Часть вторая director")
	fmt.Println("")

	compF := NewFrontendComputerBuilder()
	director := NewDirector(compF.SSD(2048))
	fmt.Println(director.BuildComputer())
	fmt.Println(director.BuildComputer())
	fmt.Println(director.BuildComputer())
	fmt.Println("")

	director = NewDirector(compF.RAM(64))
	fmt.Println(director.BuildComputer())
	fmt.Println(director.BuildComputer())
}

type frontendComputerBuilder struct {
	Builder
}

func NewFrontendComputerBuilder() ComputerBuilderI {
	return frontendComputerBuilder{}.CPU("core i5").RAM(8).SSD(256)
}

type backendComputerBuilder struct {
	Builder
}

func NewBackendComputerBuilder() ComputerBuilderI {
	return backendComputerBuilder{}.CPU("core i7").RAM(16).SSD(1024)
}

//Директор - это некто, кто занимается построением сложных объектов, т.е. может быть некий менеджер интерфейсов, который строит интерфейсы(строит диалоговые окны GUI), может быть построитель стратегий (может менять компоненты стратегии)

type Director struct {
	b ComputerBuilderI
}

func NewDirector(b ComputerBuilderI) *Director {
	return &Director{
		b: b,
	}
}

func (d *Director) BuildComputer() Computer {
	return d.b.Build()
}
