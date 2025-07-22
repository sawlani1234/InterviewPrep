package main

import (
	"bufio"
	"fmt"
	"os"
	"solid_design/tic_tac_toe/common"
	"solid_design/tic_tac_toe/entity"
	"solid_design/tic_tac_toe/enum"
	"strconv"
	"strings"
)


func main() {

    board := entity.NewBoard(3)

    player1 := entity.NewPlayer("shubham",enum.O)
    player2 := entity.NewPlayer("swarnim",enum.X)

    g := entity.NewGame([]entity.Player{player1,player2},board)

    fmt.Println("Welcome to SHUBHAM SAWLANI TIC TAC TOE GAMING\n Player 1: ", player1.GetName(),"\n Player 2: ",player2.GetName())
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for ;; {

        fmt.Println("\nCurrent Board is \n\n")
        board.PrintBoard()
        fmt.Println("\n\nIts turn of player :",g.GetCurrentPlayerTurn()+1)
        scanner.Scan()
        input := scanner.Text()

        coord := strings.Split(input,",")

        i,err := strconv.Atoi(coord[0])
        if err != nil {
            fmt.Println("invalid input retry")
            continue
        }
        j,err := strconv.Atoi(coord[1])
        
        if err != nil {
            fmt.Println("invalid input retry")
            continue
        }

        err = g.Play(i,j)
        if err != nil {
            fmt.Println(err.Error())
            continue
        }

        ans := g.GetWinner()
        
        if ans == common.NO_RESULT {
            continue
        }

        if ans == common.TIE {
            fmt.Println("game is TIED")
            break
        }

        fmt.Println("WINNER IS ",ans)
        break;

                
        
    }



}