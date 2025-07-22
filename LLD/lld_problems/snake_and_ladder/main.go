package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"snake_and_ladder/lld"
)


func main() {

	// cells addition

	cells := make([]lld.Cell,101)

	has := make(map[int]bool,0)

	for i:=1;i<=100;i++ {
		cells[i] = lld.NewCell(i,lld.NewJump(-1,-1))
	}
	// snake addition
	for i:=0;i<10;i++ {
		start,end := 1+rand.Intn(30),30+rand.Intn(65)
       
		if has[start] || has[end] {
			continue
		}
		has[start] = true 
		has[end] = true 
		cells[end].GetJump().SetStart(end)
		cells[end].GetJump().SetEnd(start)

		fmt.Println("Added snake at position ", end, "from ",start, " to ", end)
	}



	// ladder addition 

	for i:=0;i<10;i++ {
		start,end := 1+rand.Intn(30),30+rand.Intn(65)
		
		if has[start] || has[end] {
			continue	
		}
		has[start] = true 
		has[end] = true
		cells[start].GetJump().SetStart(start)
		cells[start].GetJump().SetEnd(end)
		
		fmt.Println("Added ladder at position ", end, "from ",start, " to ", end)
	}


	// player addition

	l := list.New()
	l.PushBack(lld.NewPlayer(1,"Shubham"))
	l.PushBack(lld.NewPlayer(2,"Nishi"))

	// dice addition
    dice := lld.NewDice(2)


	// board craetion
	board := lld.NewBoard(cells)
	
	game := lld.NewGameManager(board,l,dice)
	game.StartGame()
}