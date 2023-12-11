package onetwo

import (
	"bufio"
	"fmt"
	"os"
)

func InputCovert(input string) ([]int, error) {
	if len(input) != 4 {
		return nil, fmt.Errorf("input should be 4-digit")
	}
	ans := []int{}
	for _, v := range input {
		inputToInt := int(v - '0')
		if inputToInt < 0 || inputToInt > 9 {
			return nil, fmt.Errorf("input should be 4-digit")
		}
		ans = append(ans, inputToInt)
	}
	return ans, nil
}

func RunCliGame() {
	game := NewGame()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Please input your guess: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err.Error())
		}

		inputAns, err := InputCovert(text[:len(text)-1])
		if err != nil {
			fmt.Println(err)
			continue
		}

		A, B := game.CheckAnswer(inputAns)
		fmt.Printf("%dA%dB\n", A, B)
		if A == 4 {
			fmt.Println("You Win!")
			break
		}
	}
}
