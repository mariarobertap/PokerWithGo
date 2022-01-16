package main

import (
	"fmt"
)

type handPairs struct {
	carta []cardPairs
}

type cardPairs struct {
	cartaa string
	amount int
}

type hand struct {
	carta []carta
	chips int
}

func createHand(cartas []carta, chips int) (*hand){

	var cartinhas []carta

	for i:= 0; i<len(cartas); i++{
	
		cartinhas = append(cartinhas, cartas[i])
	}

	hand := hand{cartinhas, chips}

	return &hand
}

func (h *hand)checkHandStatus(){

	if (len(h.carta) < 0){
		print("Seems like you dont have a hand yet")
		return	
	}

	if(h.carta[0].cartaa == h.carta[1].cartaa){
		if(h.carta[0].cartaa == "A"){
			print("You have the best pair in poker")
		}
		print("You have a pair!!")
	}

	for i := 0; i < 2; i++{
		if(h.carta[i].cartaa == "A" || 
		   h.carta[i].cartaa == "J" ||
		   h.carta[i].cartaa == "Q" ||
		   h.carta[i].cartaa == "K"){
			print("You have a high card!!")
		}
	}	
}


func (h *hand)checkHandWithTableStatus(cartasMesa []carta){


	var cartinhas []cardPairs
	pair := 0
	if (len(h.carta) < 0){
		fmt.Println("Seems like you dont have a hand yet")
		return	
	}
	if (len(cartasMesa) < 3){
		fmt.Println("Seems like you dont have a table yet")
		return	
	}

	for i := 0; i < len(h.carta); i++{
		cartasMesa = append(cartasMesa, h.carta[i])
	}
	

	for i := 0; i < len(cartasMesa); i++{
		pair= 0
		for j := i+1; j < len(cartasMesa); j++{
			
				if(cartasMesa[i].cartaa == cartasMesa[j].cartaa){
					pair += 1
				}
	
		}

		
			testando := cardPairs{cartasMesa[i].cartaa, pair}
		    cartinhas = append(cartinhas, testando)
		
	}


	myPairs := handPairs{cartinhas}
	fmt.Println(myPairs)
	if(len(myPairs.carta) > 0){
		if(IsTreeOfaKind(myPairs) == false){
			if(IsPair(myPairs) == false){
				isHighCard(myPairs)
			}
		}
	
		
	}
	




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

	if(tree == 3){
		fmt.Println("You have tree of a kind")
		return true
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

func (h hand) print() (hand){

	fmt.Println("Player")

	fmt.Println(h)

	return h
}

/*
func (h hand) foldHand(){

}

func (h hand) bet(betAmount){

}

func (h hand) call(){

}
*/


