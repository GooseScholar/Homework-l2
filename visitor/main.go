package main

import "fmt"

/*
Посетитель
*/

// Интерфейс посетителя
type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerBar(p *BurgerBar) string
}

// Интерфейс места, которое должен посетить посетитель
type Place interface {
	Accept(v Visitor) string
}

// Реализует интерфейс посетителя
type People struct {
}

// Посещение сушибара
func (v *People) VisitSushiBar(p *SushiBar) string {
	return p.BuySushi()
}

// Посещение пиццерии
func (v *People) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}

// Посещение бургерной
func (v *People) VisitBurgerBar(p *BurgerBar) string {
	return p.BuyBurger()
}

// Коллекция мест (город) для посещения
type City struct {
	places []Place
}

// Добавления места в коллекцию
func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

// Посещение всех мест в городе
func (c *City) Accept(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}

// Реализует интерфейс места (сушибар)
type SushiBar struct {
}

// Посещение места
func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}

// Покупка суши
func (s *SushiBar) BuySushi() string {
	return "Buy sushi..."
}

// Реализует интерфейс места (пиццерия)
type Pizzeria struct {
}

// Посещение места
func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}

// Покупка пиццы
func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

// Реализует интерфейс места (бургерная)
type BurgerBar struct {
}

// Посещение места
func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(b)
}

// Покупка бургера
func (b *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}

func main() {
	city := new(City)

	city.Add(&SushiBar{})
	city.Add(&Pizzeria{})
	city.Add(&BurgerBar{})

	c := city.Accept(&People{})
	fmt.Println(c)
}
