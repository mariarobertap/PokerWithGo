package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type carta struct {
	cartaa string
	nipe string
}
type baralho struct {
	carta []carta
	embarahado bool
}

func newBaralho() (*baralho) {

	var cartinhas []carta
    cards2 := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	nipes := []string{"CORACAO", "ESPADA", "PAUS", "OURO"}
		
	for i:= 0; i< len(nipes); i++ {
		
		for j:= 0; j < len(cards2); j++ {
			testando := carta{cards2[j], nipes[i]}
			cartinhas = append(cartinhas, testando)
		}
		
	}

	baralho := baralho{cartinhas, false}

	return &baralho
}

func  (b *baralho) baralhoLen() {

	fmt.Println("LEN:",len(b.carta))
	

	
}

func (b *baralho) print(){

	for i := 0; i< len(b.carta); i++{
		fmt.Println(b.carta[i].cartaa +" DE "+ b.carta[i].nipe)
	}
	s := strconv.FormatBool(b.embarahado)
	fmt.Println("BARALHO EMBARALHADO:" + s)
	
}

func (b *baralho) embaralhar(){

	for i := 0; i< len(b.carta); i++{
		random := rand.Intn(len(b.carta))
		b.carta[random], b.carta[i] = b.carta[i], b.carta[random]
		
	}

	b.embarahado = true

	b.print()
	
}

func (b *baralho) generateHand()  ([]carta){

	var handCards []carta
	

	for i := 0; i<2; i++{
		random := rand.Intn(len(b.carta))
		testando := carta{b.carta[random].cartaa, b.carta[random].nipe}
		b.RemoveCard(random)
		handCards = append(handCards, testando)

	}


	return handCards

}

func (b *baralho) generateTableCards(total int) ([]carta){

	var handCards []carta
	
	if(total > 3){
		print("No")
		return handCards
	}
	
	for i := 0; i<total; i++{
		random := rand.Intn(len(b.carta))
		testando := carta{b.carta[random].cartaa, b.carta[random].nipe}
		b.RemoveCard(random)
		handCards = append(handCards, testando)

	}

	return handCards

}

func (b *baralho) RemoveCard(index int) {
    b.carta = append(b.carta[:index], b.carta[index+1:]...)
}