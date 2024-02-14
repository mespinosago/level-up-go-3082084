package main

import (
	"flag"
	"log"
)

type Stack []rune

func (s *Stack) Last() rune {
	if s== nil || len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Push(r rune) {
	if s== nil {
		stack := Stack{}
		s = &stack
	}
	*s = append(*s, r)
}

func (s *Stack) Pop() rune {
	var r rune
	if s == nil || len(*s) == 0 {
		return 0
	}
	r = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return r
}

func (s *Stack) IsEmpty() bool {
	return s == nil || len(*s) == 0
}

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	stack := Stack{}
	for _, r := range expr {
		switch r {
		case '{', '[', '(':
			stack.Push(r)
		case ')':
			if stack.Pop() != '(' {
				return false
			}
		case ']':
			if stack.Pop() != '[' {
				return false
			}
		case '}':
			if stack.Pop() != '{' {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
