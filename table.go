package main

import (
	"fmt"
)

type table struct {
	carta []carta
	pot int
}

func createTable(cartas []carta, pot int) (*table){

	var cartinhas []carta

	for i:= 0; i<len(cartas); i++{
	
		cartinhas = append(cartinhas, cartas[i])
	}

	table := table{cartinhas, pot}

	return &table
}

func (t *table) addCard(cartas carta){
	t.carta = append(t.carta, cartas)
}

func (t *table) print(){
	fmt.Println("--------------Table---------------")
	fmt.Println("CARDS: [", t.carta, " ]")
	fmt.Println("POT AMOUNT: [", t.pot, " ]")
	fmt.Println("-----------------------------------")
}

func (t *table) getCards() ([]carta){
	return t.carta
}

func (t *table) setPot(amount int) {

	t.pot += amount 
}


