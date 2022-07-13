package main

import (
	"fmt"
	"log"
)

//Интерфейс состояния
type state interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

//Контекст
type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state

	currentState state

	itemCount int
	itemPrice int
}

//Конструктор контекста
func newVendingMachine(itemCount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &hasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &hasMoneyState{
		vendingMachine: v,
	}
	noItemState := &noItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

//Действие: выбрать предмет
func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

//Действие: добавить предмет
func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

//Действие: внести деньги
func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

//Действие: выдать предмет
func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

//Установить состояние
func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

//Увеличить количество предметов
func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}

//Состояние: не имеет предметов
type noItemState struct {
	vendingMachine *vendingMachine
}

//Действие: выбрать предмет
func (i *noItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

//Действие: добавить предмет
func (i *noItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

//Действие: внести деньги
func (i *noItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

//Действие: выдать предмет
func (i *noItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}

//Состояние: есть предметы
type hasItemState struct {
	vendingMachine *vendingMachine
}

//Действие: выбрать предмет
func (i *hasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item request\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

//Действие: добавить предмет
func (i *hasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

//Действие: внести деньги
func (i *hasItemState) insertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

//Действие: выдать предмет
func (i *hasItemState) dispenseItem() error {
	return fmt.Errorf("Please select item first")
}

//Состояние: выдаёт предметы
type itemRequestedState struct {
	vendingMachine *vendingMachine
}

//Действие: выбрать предмет
func (i *itemRequestedState) requestItem() error {
	return fmt.Errorf("Item already requested")
}

//Действие: добавить предмет
func (i *itemRequestedState) addItem(count int) error {
	return fmt.Errorf("Item Dispense in progress")
}

//Действие: внести деньги
func (i *itemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

//Действие: выдать предмет
func (i *itemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}

//Состояние: получил деньги
type hasMoneyState struct {
	vendingMachine *vendingMachine
}

//Действие: выбрать предмет
func (i *hasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

//Действие: добавить предмет
func (i *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

//Действие: внести деньги
func (i *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

//Действие: выдать предмет
func (i *hasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}

func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
