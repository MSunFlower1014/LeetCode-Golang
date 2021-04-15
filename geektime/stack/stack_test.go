package stack

import "testing"

type Stack struct {
	List   [10]int
	length int
}

func (s *Stack) Push(i int) {
	if s.length == 9 {
		return
	}
	s.length++
	s.List[s.length] = i
}

func (s *Stack) Pop() int {
	if s.length == 0 {
		return 0
	}
	result := s.List[s.length]
	s.List[s.length] = 0
	s.length--
	return result
}
func TestStack(t *testing.T) {
	stack := &Stack{}

	stack.Push(7)

	i := stack.Pop()
	if i != 7 {
		t.Fatalf(" stack pop value is err")
	}

	i = stack.Pop()
	if i != 0 {
		t.Fatalf(" stack pop value is err , value should be 0")
	}
}

var d = map[int]int{40: 41, 91: 93}

//括号匹配
func TestParenthesisMatching(t *testing.T) {
	s := "(([]))"
	stack := &Stack{}
	for k, value := range []rune(s) {
		t.Logf("k: %v , value : %v , string value : %v", k, value, string(value))
	}
	for _, value := range []rune(s) {
		temp := stack.Pop()
		if temp != 0 {
			if int(value) == d[temp] {
				continue
			}
			stack.Push(temp)
			stack.Push(int(value))
		} else {
			stack.Push(int(value))
		}
	}
	temp := stack.Pop()
	if temp != 0 {
		t.Fatalf("stack not empty , temp value is %v", string(temp))
	}
}
