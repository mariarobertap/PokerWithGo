package main

import (
	"fmt"
)


type player struct {
	playerHand hand
	name string
}

type hand struct {
	carta []carta
	chips int
}


type handPairs struct {
	carta []cardPairs
}

type cardPairs struct {
	cartaa string
	amount int
}



func createHand(cartas []carta, chips int) (*hand){

	var cartinhas []carta

	for i:= 0; i<len(cartas); i++{
	
		cartinhas = append(cartinhas, cartas[i])
	}

	hand := hand{cartinhas, chips}

	return &hand
}

func createPlayer(cartas []carta, chips int, name string) (*player){

	var playerCards []carta

	for i:= 0; i<len(cartas); i++{
		playerCards = append(playerCards, cartas[i])
	}

	hand := hand{playerCards, chips}

    player := player{hand, name}

	return &player
}


func (p player) getPlayerName(){
	fmt.Println(p.name)
}


func (p *player) checkHandWithTableStatus(cartasMesa []carta){

	var playerHand []carta

	
	if (len(p.playerHand.carta) < 0){
		fmt.Println("Seems like you dont have a hand yet")
		return	
	}
	if (len(cartasMesa) < 3){
		fmt.Println("Seems like you dont have a table yet")
		return	
	}


	playerHand = p.appendWithTableCards(cartasMesa)

	

	if(isFlush(playerHand) == true){
		return
	}

    listOfPairs := generateListOfPairs(playerHand)


	myPairs := handPairs{listOfPairs}

	if(len(myPairs.carta) > 0){                                                            

		if(IsTreeOfaKind(myPairs) == false){

			if(IsPair(myPairs) == false){

				isHighCard(myPairs)
			}
		}	
	}
}

func generateListOfPairs(playerHand []carta) ([]cardPairs){

	var listPairs []cardPairs
	var amountPair int

	for i := 0; i < len(playerHand); i++{
		amountPair = 0

		for j := i+1; j < len(playerHand); j++{
			
				if(playerHand[i].cartaa == playerHand[j].cartaa){
					amountPair += 1
				}
	
		}

		pair := cardPairs{playerHand[i].cartaa, amountPair}
		listPairs = append(listPairs, pair)
		
	}

	return listPairs

}

func IsTreeOfaKind(deck handPairs) (bool) {

	var tree int
	if(len(deck.carta) <= 1){
		return false
	}

	for i := 0; i < len(deck.carta); i++{

		for j := i+1; j < len(deck.carta); j++{
			if(deck.carta[i].cartaa == deck.carta[j].cartaa){
				tree = deck.carta[i].amount + deck.carta[j].amount
				if(tree == 3){
					fmt.Println("You have tree of a kind")
					return true
				}
			}

		}
	}

	return false
}

func IsPair(deck handPairs) (bool) {

	var tree int
	if(len(deck.carta) == 1){
		tree = deck.carta[0].amount
		
	}else{
		for i := 0; i < len(deck.carta); i++{

			for j := i+1; j < len(deck.carta); j++{
				if(deck.carta[i].cartaa == deck.carta[j].cartaa){
					tree += deck.carta[i].amount + deck.carta[j].amount
					
				}
	
			}
		}
	
	}

	if(tree == 0){
		return false
	}

	s:=fmt.Sprintf("You have %d pairs",tree)
	fmt.Println(s)

	
	return true
}

func isHighCard(deck handPairs) (bool) {

	for i := 0; i < len(deck.carta); i++{
		if(deck.carta[i].cartaa == "A" || 
		deck.carta[i].cartaa == "J" ||
		deck.carta[i].cartaa == "Q" ||
		deck.carta[i].cartaa == "K"){
			fmt.Println("You have a high card!!")
			return true
		}
	}
	return false
}

func isFlush(deck []carta) (bool) {
	var amount int  

	for i := 0; i < len(deck); i++{
		amount = 1

		for j := i+1; j < len(deck); j++{
			if(deck[i].nipe == deck[j].nipe){
				amount += 1		
			}
		}

		if(amount == 5){
			fmt.Println("You have a flush")
			return true
		}
	}

	return false
}

func (p player) print(){



	fmt.Println("--------------Player---------------")
	fmt.Println("PLAYER: [", p.name, " ]")
	fmt.Println("CHIPS: [", p.playerHand.chips, " ]")
	fmt.Println("-----------------------------------")

}

func (p player) printChips(){

	fmt.Println("-----------------------------------")
	fmt.Println("PLAYER: [", p.name, " ]")
	fmt.Println("CHIPS: [", p.playerHand.chips, " ]")
	fmt.Println("-----------------------------------")

}


func (p player) appendWithTableCards(cardsAux []carta)  ([]carta){
	
	for i := 0; i < len(p.playerHand.carta); i++{
		cardsAux = append(cardsAux, p.playerHand.carta[i])
	}

	return cardsAux
}



/*
func (h hand) foldHand(){

}

func (h hand) bet(betAmount){

}

*/
func (p *player) call(amount int){

	p.playerHand.chips -= amount
	buffer := fmt.Sprintf("%s Call [%d] chips", p.name, amount)
    fmt.Println(buffer)

}



