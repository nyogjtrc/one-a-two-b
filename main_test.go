package main

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

func TestInputCovert(t *testing.T) {
	t.Run("input should be 4-digit", func(t *testing.T) {
		input := "1234"
		ans, err := inputCovert(input)
		assert.NoError(t, err, "input should be 4-digit")
		assert.Len(t, ans, 4, "input should be 4-digit")
	})

	t.Run("input should not be 3-digit", func(t *testing.T) {
		input := "123"
		_, err := inputCovert(input)
		assert.Error(t, err, "input should not be 3-digit")
	})

	t.Run("input should not be 5-digit", func(t *testing.T) {
		input := "12345"
		_, err := inputCovert(input)
		assert.Error(t, err, "input should not be 5-digit")
	})

	t.Run("input should not be 4-letter", func(t *testing.T) {
		input := "abcd"
		_, err := inputCovert(input)
		assert.Error(t, err, "input should not be 4-letter")
	})

	t.Run("input should not be 4-symbol", func(t *testing.T) {
		input := "!@#$"
		_, err := inputCovert(input)
		assert.Error(t, err, "input should not be 4-symbol")
	})

	t.Run("input should not be 4-mixed", func(t *testing.T) {
		input := "a1b2"
		_, err := inputCovert(input)
		assert.Error(t, err, "input should not be 4-mixed")
	})
}

func TestCheckAnswer(t *testing.T) {
	// test case table
	type testCase struct {
		input  []int
		answer []int
		A      int
		B      int
	}
	testcases := []testCase{
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, 4, 0},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}, 0, 4},
		{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, 0, 0},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 5}, 3, 0},
		{[]int{1, 2, 3, 4}, []int{1, 2, 5, 3}, 2, 1},
		{[]int{1, 2, 3, 4}, []int{1, 5, 2, 3}, 1, 2},
		{[]int{1, 2, 3, 4}, []int{5, 1, 2, 3}, 0, 3},
		{[]int{1, 2, 3, 4}, []int{1, 2, 5, 4}, 3, 0},
		{[]int{1, 2, 3, 4}, []int{1, 5, 2, 4}, 2, 1},
		{[]int{1, 2, 3, 4}, []int{1, 5, 4, 3}, 1, 2},
		{[]int{1, 2, 3, 4}, []int{5, 1, 4, 3}, 0, 3},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("input %d, answer %d", tc.input, tc.answer), func(t *testing.T) {
			A, B := checkAnswer(tc.input, tc.answer)
			assert.Equal(t, tc.A, A, "A should be %d", tc.A)
			assert.Equal(t, tc.B, B, "B should be %d", tc.B)
		})
	}
}
