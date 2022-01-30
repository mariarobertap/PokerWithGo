package main


import (
	"fmt"
	"math/rand"
    "time"
    "os"
    "os/exec"
    "runtime"
)

var clear map[string]func() //create a map for storing clear funcs


func main(){

    rand.Seed(time.Now().UnixNano())
    menu()
	
}


func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func systemcls() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

func menu(){

	var opcao int
	var deck *baralho

	  for {
		systemcls()
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
			systemcls()
			fmt.Println("Baralho embaralhado com sucesso!.... Aperte qualquer tecla para voltar")
			fmt.Scanln(&opcao)
		}else if opcao == 3{
			systemcls()
			deck.print()
			fmt.Scanln(&opcao)
			
		}else if opcao == 4{
			systemcls()
			playPokerTest()
			fmt.Scanln(&opcao)
		}else {
			fmt.Print("Escolha uma opcao valida")
		}	

	
	}
}

func playPoker(){

	deck := newBaralho()			

	//deck.baralhoLen()
	//player1 := createHand(deck.generateHand(), 500)
	player1 := createPlayer(deck.generateHand(), 500, "Maria")
	//player1.getPlayerName()
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

	//player1.call(100)
	//player1.print()
	//print(len(cartas))
}


func playPokerTest() {
	deck := newBaralho()			

	
	player1 := createPlayer(deck.generateHand(), 500, "Maria")
	player2 := createPlayer(deck.generateHand(), 500, "Bot")
	table := createTable(deck.generateTableCards(3), 0)

	systemcls()

	for i:= 0; i<3; i++{
		
		player1.printChips()
		if(callOrBet(10, player1) == 0){
			systemcls()
			player2.call(10)
			table.setPot(20)
			fmt.Println("BOT JUST CALLED")
			player1.print()
			if(i > 0){
				cartass:= deck.generateTableCards(1)
				table.addCard(cartass[0])
			}
			
		
			table.print()
		}
		
		
	}
	player1.printChips()

	player1.print()

	
	table.print()
}

func callOrBet(amount int, playerHand *player) (int){

	var opcao int

	buffer := fmt.Sprintf("1 - CALL (%d) chips", amount)
	fmt.Println(buffer)
	fmt.Println("2 - BET")
	

	fmt.Scanln(&opcao)

	if(opcao == 1){
		playerHand.call(amount)
		return 0
	}


    return 0
}
