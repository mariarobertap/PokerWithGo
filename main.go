package main


import (
	"fmt"
	"math/rand"
    "time"

)

func main(){

    rand.Seed(time.Now().UnixNano())

	deck := newBaralho()

	//deck.baralhoLen()
	player1 := createHand(deck.generateHand(), 500)
	player1.print()
	//player2 := createHand(deck.generateHand(), 500)
	//player2.print()
	//player1.checkHandStatus()
	///deck.baralhoLen()
	table := createTable(deck.generateTableCards(3), 100)
	//table.print()
	//cartas := table.getCards()
	cartass:= deck.generateTableCards(1)
	table.addCard(cartass[0])
	//table.print()
	cartass = deck.generateTableCards(1)
	table.addCard(cartass[0])
	table.print()
	deck.baralhoLen()

	player1.checkHandWithTableStatus(table.getCards())
	//print(len(cartas))
	
	
}

func menu(){

	var opcao int
	var deck *baralho

	  for {
		fmt.Println("Bem vindo ao Poker")
		fmt.Println("1- Criar baralho")
		fmt.Println("2- Embaralhar")
		fmt.Println("3- Mostrar baralho")
		fmt.Println("4- Jogar")
		fmt.Scanln(&opcao)
		if opcao == 1{
			deck = newBaralho()
		}else if opcao == 2{
			deck.embaralhar()
		}else if opcao == 3{
			deck.print()
		}else {
			fmt.Println("This is not a valid option")
		}	

		systemCls()
	}
}

func systemCls(){
	fmt.Println("000000000")
}