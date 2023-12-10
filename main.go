package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

// generateAnswer generates a random 4-digit answer
func generateAnswer() []int {
	source := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ans := []int{}
	// random generate integer
	for i := 10; i > 0; i-- {
		newIndex := rand.Intn(i)
		ans = append(ans, source[newIndex])
		source = append(source[:newIndex], source[newIndex+1:]...)
		if len(ans) == 4 {
			break
		}
	}
	return ans
}

func inputCovert(input string) ([]int, error) {
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

func checkAnswer(inputAns []int, answer []int) (int, int) {
	A := 0
	B := 0
	for i, v := range inputAns {
		for j, w := range answer {
			if v == w {
				if i == j {
					A++
				} else {
					B++
				}
			}
		}
	}
	return A, B
}

func main() {
	fmt.Println("welcome to 1A2B game")

	answer := generateAnswer()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Please input your guess: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err.Error())
		}

		inputAns, err := inputCovert(text[:len(text)-1])
		if err != nil {
			fmt.Println(err)
			continue
		}

		A, B := checkAnswer(inputAns, answer)
		fmt.Printf("%dA%dB\n", A, B)
		if A == 4 {
			fmt.Println("You Win!")
			break
		}
	}

}
