package onetwo

import (
	"math/rand"
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

type Game struct {
	answer []int
}

func NewGame() *Game {
	return &Game{
		answer: generateAnswer(),
	}
}

func (g *Game) GetAnswer() []int {
	return g.answer
}

func (g *Game) CheckAnswer(inputAns []int) (int, int) {
	A := 0
	B := 0
	for i, v := range inputAns {
		for j, w := range g.answer {
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
