package main


import (
	"fmt"

)

func main(){

    menu()	
}

func menu(){

	var opcao int
	var deck *baralho

	  for {
		fmt.Println("Bem vindo ao Poker")
		fmt.Println("1- Criar baralho")
		fmt.Println("2- Embaralhar")
		fmt.Println("3- Mostrar baralho")
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