package main

import (
    "testing"
)

func TestCleanInput(t *testing.T) {
    testCases := []struct {
        input      string
        expected []string
    }{
        {
            input: "   ",
            expected: []string{},
        },
        {
            input: " hello  ",
            expected: []string{"hello"},
        },
        {
            input: " hello world  ",
            expected: []string{"hello", "world"},
        },
        {
            input: "  heLLo WorlD  ",
            expected: []string{"hello", "world"},
        },
    }

    for _, c := range testCases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("length don't match '%v' vs '%v'", actual, len(c.expected))
            continue
        }
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
            }
        }
    }
}
