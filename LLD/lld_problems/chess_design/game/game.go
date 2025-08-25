package game

import (
	"chess_design/board"
	"chess_design/cell"
	"chess_design/enum"
	"chess_design/player"
	"container/list"
	"fmt"
)


type Game struct {
	playerA player.Player
	playerB player.Player
	playerQueue *list.List
	board board.Board
}

func NewGame(playerA player.Player,playerB player.Player) *Game{ 
	b := board.NewBoard(make(map[player.Player][]cell.Piece,0))
	q := list.New()
	q.PushBack(playerA)
	q.PushBack(playerB)
	return &Game{playerA: playerA,playerB: playerB,playerQueue: q,board: b}
}


func (g Game) getWinner() {

   if g.board.GetNoOfPieces(enum.WHITE) > g.board.GetNoOfPieces(enum.BLACK) {
	   fmt.Println("Player A ",g.playerA.GetName(), " won the game")
	   return 
   }

    if g.board.GetNoOfPieces(enum.BLACK) > g.board.GetNoOfPieces(enum.WHITE) {
		 fmt.Println("Player B ",g.playerA.GetName(), " won the game")
	  return 
   }
   
   fmt.Println("Game tied between both players ",g.playerA.GetName(), " : ",g.playerB.GetName())
	
}



func (g *Game) StartGame() {

	fmt.Println("Welcome to Sawlanis Chess Game Blitz 20 iteration game haha")
	
	for i:=0;i<10;i++ {
		// Assuming plyaer A has white pieces and plyaer B has black pieces
		g.board.PrintBoard()
		
		flag := false 

		for ;!flag; {

			var pr,pc,dr,dc int
			var player player.Player
			var piece string
			if i%2 == 0 {
				player = g.playerA
				piece = "BLUE"
			} else {
				player = g.playerB
				piece = "GREEN"
			}

			fmt.Println("\n Enter row , col of the piece",player.GetName(), " for color ", piece)
			fmt.Scanf("%d %d",&pr,&pc)

			fmt.Println("\n Enter row , col of the destination")
			fmt.Scanf("%d %d",&dr,&dc)
			cells := g.board.GetCells()
			if g.valid(pr,pc) && g.valid(dr,dc) {
				if i%2 == 0 && cells[pr][pc].Piece.GetPieceColor() == enum.WHITE {
				   flag = g.board.CanMove(*cells[pr][pc],*cells[dr][dc])
				} else if i%2 !=0 && cells[pr][pc].Piece.GetPieceColor() == enum.BLACK {
					flag = g.board.CanMove(*cells[pr][pc],*cells[dr][dc]) 
				}
			}

			if flag {
				g.board.MovePiece(*cells[pr][pc],*cells[dr][dc])
			} else {
				fmt.Println("Invalid Move Man")
			}
		}
	}

	g.getWinner()
    
}

func (g *Game) valid(row,col int) bool {
	cells := g.board.GetCells()
	n,m := len(cells),len(cells[0])

	return row >= 0 && row < n && col >= 0 && col < m
}