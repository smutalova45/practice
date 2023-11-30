package main

import (
	"fmt"
	"math/rand"
)
// const chances=3
func main(){

player:=Player{} //struct of player
game:=Game{} // struct of game

name,age,favouriteNumber,chances:=getPlayerName() // getplayer digan funksiya orqali oyinchi ismi olinadi
newPlayer:=player.NewPlayer(name,age,favouriteNumber,chances) // newplayer PLayer structni methodi , bu method orqali yengi oyinchi create qlinadi va olingan ism va cons chance beriladi
newGame:=game.NewGame(newPlayer)// newGame game ni methodi bu orqali yengi game create qilinadi newplayer bilan
newGame.StartGame()


}

type Game struct {
	RandomNumber int 
	isPrime bool
	Player   Player //game ni ichidegi player 
}

func (g Game) NewGame(player Player) Game{ // Game structga newGame methodi ochildi bu player uchun

	return Game {
		RandomNumber: rand.Intn(10),
		Player:   player,
		isPrime: isPrime(g.RandomNumber),
	}
}

func isPrime(num int) bool {
	if num <= 1 { 
	 return false
	}
   
	for i := 2; i*i <= num; i++ {
	 if num%i == 0 {
	  return false
	 }else{
		fmt.Println("random number is prime ")
	 }
	}
   
	return true
}
func (g Game) StartGame(){ //game ni methodi oyin boshlash uchun
	fmt.Println("welcome ",g.Player.Name,", age is: ",g.Player.Age,", you have :",g.Player.Chances,"chances only, "," Favourite number is: ",g.Player.favouriteNumber)
	fmt.Println("THIS IS GUESSING GAME")
    fmt.Println(g.isPrime)
	for i:=0;i<g.Player.Chances;i++{ 
		var n int 
	
        fmt.Print("ENTER YOUR NUMBER ")
		fmt.Scan(&n)
		if g.Player.Chances >5 {

			rand.Intn(20)
		}
		if n==g.RandomNumber {
			if  g.Player.favouriteNumber==g.RandomNumber {
				fmt.Println("and your favourite number was random number")
			}
			fmt.Println("you won! in " ,i+1,"tries")
			return
        }else {
			fmt.Println("incorrect")
			if g.RandomNumber>1 && g.RandomNumber<5{
				fmt.Println("it is between 1 and 5")
			}else if g.RandomNumber>5 && g.RandomNumber<10{
	
				fmt.Println("it is between 5 and 10")
			}else if g.RandomNumber>1 && g.RandomNumber<5{
				fmt.Println("between 1 and 5")
			}
		}
		

	}
	fmt.Println("you lost!\n the random number was ",g.RandomNumber)
}





type Player struct{
	Name string 
	Age int 
	favouriteNumber int 
	Chances int
}

func(p Player) NewPlayer( name string  ,age int ,favouriteNum int, chans int)Player{ // method of player //newplayer ochildi player orqali
	return Player{
		Name: name,
		Chances: chans,
		favouriteNumber:favouriteNum ,
		Age: age,

		
	}
}

func getPlayerName()(string,int,int,int){ //funksiya umumiy
	var(
		name string
		age int 
		favouriteNumber int 
		chances int 
	)
	fmt.Print("enter your name ")
	fmt.Scan(&name)
	fmt.Println()
	fmt.Print("enter your age ")
	fmt.Scan(&age)
	fmt.Println()
	fmt.Print("enter your favourite number : ")
	fmt.Scan(&favouriteNumber)
	fmt.Print("how many chances do you want to try: ")
	fmt.Scan(&chances)
	return name,age,favouriteNumber,chances
}