// Package facade is an example of the Facade Pattern.
package main

import (
	"fmt"
	"strings"
)

/*
Фасад
*/

// Реализует мужчину и фасад
type Man struct {
	house *House
	tree  *Tree
	child *Child
}

// Конструктор Man
func NewMan() *Man {
	return &Man{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

// Возвращает что мужчина должен сделать
func (m *Man) todo(h, t, c string) string {
	result := []string{
		m.house.Build(h),
		m.tree.Grow(t),
		m.child.Born(c),
	}
	return strings.Join(result, "\n")
}

// Реализует подсистему "Дом"
type House struct {
}

// Реализуем постройку дома
func (h *House) Build(s string) string {
	return s
}

// Реализует подсистему "Дерево"
type Tree struct {
}

// Реализуем посадку дерева
func (t *Tree) Grow(s string) string {
	return s
}

// Реализуем подсистему "Ребенок"
type Child struct {
}

// Реализуем роды ребенка
func (c *Child) Born(s string) string {
	return s
}

func main() {
	man := NewMan().todo("построил", "сгнило", "посадили")
	man2 := NewMan().todo("построил", "посадил", "скоро родится")
	fmt.Println(man)
	fmt.Println(man2)
}
