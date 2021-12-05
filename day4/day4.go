package main

import (
	advent "advent/reuse"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	var rawData []string
	var wonBoards []int
	var winBoards []int
	var lastNumber int
	inputFile, err := filepath.Abs("input")
	if err != nil {
		log.Fatal(err)
	}
	rawData = advent.LoadInput(inputFile)

	bingoMatrices, bingoNumbers := buildBingo(rawData)
	lastMatrices, _ := buildBingo(rawData)
	for _, number := range bingoNumbers {
		fillNumber(&bingoMatrices, number, &wonBoards)
		isLine, matrix, _ := checkLines(&bingoMatrices)
		if isLine {
			matrixSum := sumMatrix(matrix) * number
			fmt.Println(matrixSum)
			break
		}
		isColumn, matrix, _ := checkColumns(&bingoMatrices)
		if isColumn {
			matrixSum := sumMatrix(matrix) * number
			fmt.Println(matrixSum)
			break
		}
	}
	for _, number := range bingoNumbers {
		fillNumber(&lastMatrices, number, &winBoards)
		lastNumber = number
		allWinners(&lastMatrices, &winBoards, &lastNumber)
	}

	fmt.Println(sumMatrix(lastMatrices[winBoards[len(winBoards)-1]]))
}

func allWinners (board *[100][5][]int, winners *[]int, lastNumber *int) {
	var counter int
	// Check all lines on all boards
	for k, b := range board {
		counter = 0
		for i, _ := range b {
			counter = board[k][i][1] + board[k][i][3] + board[k][i][5] + board[k][i][7] + board[k][i][9]
			if counter == 5 && !checkBoard(winners, k){
				*winners = append(*winners, k)
				fmt.Println(*lastNumber)
			}
		}
	}
	// check all columns on all boards
	for k, _ := range board {
		for j:=1; j<=9; j+=2 {
			counter = board[k][0][j] + board[k][1][j] + board[k][2][j] + board[k][3][j] + board[k][4][j]
			if counter == 5 && !checkBoard(winners, k) {
				*winners = append(*winners, k)
				fmt.Println(*lastNumber)
			}
		}
	}
}
func checkBoard(wonBoardList *[]int, boardId int) bool {
	for _,v := range *wonBoardList {
		if v == boardId {
			return true
		}
	}
	return false
}

func sumMatrix (matrix [5][]int)  int {
	var sum int
	for k, _ := range matrix {
		for j:=0; j<10; j+=2 {
			if matrix[k][j+1] == 0 {
				sum+=matrix[k][j]
			}
		}
	}
	return sum
}

func checkLines(board *[100][5][]int) (bool, [5][]int, int) {
	var counter int
	for k, b := range board {
		counter = 0
		for i, _ := range b {
			counter = board[k][i][1] + board[k][i][3] + board[k][i][5] + board[k][i][7] + board[k][i][9]
			if counter == 5 {
				return true, board[k], k
			}
		}
	}
	return false, board[0], 1000
}

func checkColumns(board *[100][5][]int) (bool, [5][]int, int) {
	counter:=0
	for k, _ := range board {
			for j:=1; j<=9; j+=2 {
				counter = board[k][0][j] + board[k][1][j] + board[k][2][j] + board[k][3][j] + board[k][4][j]
				if counter == 5 {
					return true, board[k], k
				}
		}
	}
	return false, board[0], 1000
}

func fillNumber (board *[100][5][]int, number int, wonBoardList *[]int) {
	for k, b := range board {
		if !checkBoard(wonBoardList, k) {
			for i, boardLine := range b {
				for j, boardNumber := range boardLine {
					if j%2 == 0 {
						if number == boardNumber {
							board[k][i][j+1] = 1
						}
					}
				}
			}
		}
	}
}

func buildBingo(data []string) ([100][5][]int, []int) {
	var bingoNumbers []int
	var matrixCounter int
	var boardCounter int
	bigMatrix := [100][5][]int{}

	for k, v := range data {
		if k == 0 {
			for _, z := range strings.Split(v, ",") {
				num, err := strconv.Atoi(z)
				if err != nil {
					log.Fatal(err)
				}
				bingoNumbers = append(bingoNumbers, num)
				continue
			}
		} else {
			if len(v) == 0 && boardCounter == 0 && matrixCounter == 0 {
				continue
			}
			if len(v) == 0 && matrixCounter != 0 {
				matrixCounter = 0
				boardCounter++
			}
			if len(v) != 0 {
				for _, v := range strings.Split(v, " ") {
					if v == "" {
						continue
					}
					num, err := strconv.Atoi(v)
					if err != nil {
						log.Fatal(err)
					}
					bigMatrix[boardCounter][matrixCounter] = append(bigMatrix[boardCounter][matrixCounter], num)
					bigMatrix[boardCounter][matrixCounter] = append(bigMatrix[boardCounter][matrixCounter], 0)
				}
				matrixCounter++
			}
		}
	}
	return bigMatrix, bingoNumbers
}