package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
mainLoop:
	for {
	startLoop:
		fmt.Println("Enter chessboard size:")
		sizeInput := getInput(reader)
		size, err := strconv.Atoi(sizeInput)

		if err != nil {
			fmt.Println("It is not a valid size, try again.")
			fmt.Println()
			goto startLoop
		} else {
			drawBoard(size)
		}

		fmt.Println()

	retryRequest:
		fmt.Println("Try again? (y/n)")
		retry := getInput(reader)

		switch retry {
		case "y":
			fmt.Println()
			continue
		case "n":
			break mainLoop
		default:
			fmt.Println("Answer is not valid. Try again")

			fmt.Println()
			goto retryRequest
		}
	}
}
func getInput(b *bufio.Reader) string {
	fmt.Print("-> ")
	s, _ := b.ReadString('\n')
	s = strings.Replace(s, "\n", "", -1)
	return s
}

func drawBoard(n int) {
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if (i+j)%2 != 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
