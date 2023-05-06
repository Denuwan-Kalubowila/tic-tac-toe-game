package main

import (
	"fmt"
)

var boardVal = [9]string{"_", "_", "_",
	"_", "_", "_",
	"_", "_", "_"}

var playing bool = true
var currentPlayer string = "X"

var winner string

func main() {

	for playing {
		printBoard(boardVal)
		getInput(boardVal[:])
		checkHorizontal(boardVal)
		checkVertical(boardVal)
		checkDiaganal(boardVal)
		changePlayer()

	}
	printBoard(boardVal)

}

func printBoard(board [9]string) {
	fmt.Println("|" + board[0] + "|" + board[1] + "|" + board[2] + "|")
	fmt.Println("-------")
	fmt.Println("|" + board[3] + "|" + board[4] + "|" + board[5] + "|")
	fmt.Println("-------")
	fmt.Println("|" + board[6] + "|" + board[7] + "|" + board[8] + "|")

}

func getInput(boardVal []string) {
	var pos int
	fmt.Println("Enter the Posission:")
	fmt.Scan(&pos)
	if pos >= 1 && pos <= 9 && boardVal[pos-1] == "_" {
		boardVal[pos-1] = currentPlayer
	} else {
		fmt.Println("Oops is already in that spot.")
	}
}

func checkHorizontal(board [9]string) {
	if board[0] == board[1] && board[1] == board[2] && board[0] != "_" {
		winner = board[0]
		fmt.Printf("The winner is %s\n", winner)
		playing = false

	} else if board[3] == board[4] && board[4] == board[5] && board[3] != "_" {
		winner = board[3]
		fmt.Printf("The winner is %s\n", winner)
		playing = false

	} else if board[6] == board[7] && board[7] == board[8] && board[6] != "_" {
		winner = board[6]
		fmt.Printf("The winner is %s\n", winner)
		playing = false

	}

}
func checkVertical(board [9]string) {
	if board[0] == board[3] && board[3] == board[6] && board[0] != "_" {
		winner = board[0]
		fmt.Printf("The winner is %s\n", winner)
		playing = false

	} else if board[1] == board[4] && board[4] == board[7] && board[1] != "_" {
		winner = board[1]
		fmt.Printf("The winner is %s\n", winner)
		playing = false

	} else if board[2] == board[5] && board[5] == board[8] && board[2] != "_" {
		winner = board[2]
		fmt.Printf("The winner is %s\n", winner)
		playing = false

	}
}

func checkDiaganal(board [9]string) {
	if board[0] == board[4] && board[4] == board[8] && board[0] != "_" {
		winner = board[0]
		fmt.Printf("The winner is %s\n", winner)
		playing = false

	} else if board[2] == board[4] && board[4] == board[6] && board[2] != "_" {
		winner = board[2]
		fmt.Printf("The winner is %s\n", winner)
		playing = false

	}

}

func changePlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}
