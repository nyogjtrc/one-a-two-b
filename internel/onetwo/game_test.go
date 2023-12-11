package onetwo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAnswer(t *testing.T) {
	t.Run("answer length should be 4", func(t *testing.T) {
		answer := generateAnswer()
		assert.Len(t, answer, 4, "length of answer is not 4")
	})

	t.Run("answer should be different each time", func(t *testing.T) {
		answer1 := generateAnswer()
		answer2 := generateAnswer()
		assert.NotEqual(t, answer1, answer2, "answer should be different each time")
	})

	t.Run("digit should not duplicate in answer", func(t *testing.T) {
		for x := 0; x < 10; x++ {
			answer := generateAnswer()
			t.Log("answer", answer)

			checker := func(ans []int) bool {
				for i := 0; i < 4; i++ {
					for j := i + 1; j < 4; j++ {
						if !assert.NotEqual(t, answer[i], answer[j], "digit should not duplicate in answer: %v", answer) {
							return false
						}
					}
				}
				return true
			}
			checker(answer)
		}
	})
}

func TestCheckAnswer(t *testing.T) {
	// test case table
	type testCase struct {
		game  *Game
		input []int
		A     int
		B     int
	}
	testcases := []testCase{
		{&Game{[]int{1, 2, 3, 4}}, []int{1, 2, 3, 4}, 4, 0},
		{&Game{[]int{1, 2, 3, 4}}, []int{4, 3, 2, 1}, 0, 4},
		{&Game{[]int{1, 2, 3, 4}}, []int{5, 6, 7, 8}, 0, 0},
		{&Game{[]int{1, 2, 3, 4}}, []int{1, 2, 3, 5}, 3, 0},
		{&Game{[]int{1, 2, 3, 4}}, []int{1, 2, 5, 3}, 2, 1},
		{&Game{[]int{1, 2, 3, 4}}, []int{1, 5, 2, 3}, 1, 2},
		{&Game{[]int{1, 2, 3, 4}}, []int{5, 1, 2, 3}, 0, 3},
		{&Game{[]int{1, 2, 3, 4}}, []int{1, 2, 5, 4}, 3, 0},
		{&Game{[]int{1, 2, 3, 4}}, []int{1, 5, 2, 4}, 2, 1},
		{&Game{[]int{1, 2, 3, 4}}, []int{1, 5, 4, 3}, 1, 2},
		{&Game{[]int{1, 2, 3, 4}}, []int{5, 1, 4, 3}, 0, 3},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("input %d, answer %d", tc.input, tc.game.answer), func(t *testing.T) {
			A, B := tc.game.CheckAnswer(tc.input)
			assert.Equal(t, tc.A, A, "A should be %d", tc.A)
			assert.Equal(t, tc.B, B, "B should be %d", tc.B)
		})
	}
}
