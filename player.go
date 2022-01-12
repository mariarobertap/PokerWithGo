package main

import (
	"fmt"
)



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

func (h hand) print() (hand){
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


