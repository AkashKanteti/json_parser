package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Stack struct {
	data []string
}

func (s Stack) pop() {
	len := len(s.data)
	if len == 0 {
		return
	}
	s.data = s.data[:len-1]
}

func (s Stack) push(st string) {
	s.data = append(s.data, st)
}

func (s Stack) peek() string {
	len := len(s.data)
	if len == 0 {
		return ""
	}
	return s.data[len-1]
}

func (s Stack) isEmpty() bool {
	len := len(s.data)
	if len == 0 {
		return true
	}
	return false
}

func (s Stack) validate(lis []string) bool {
	if len(lis) == 0 {
		return false
	}
	for _, v := range lis {
		if v == "{" {
			s.push(v)
		} else if v == "}" {
			top := s.peek()
			if top == "{" {
				s.pop()
			}
		}
	}
	return s.isEmpty()
}

func main() {
	s := Stack{}
	files := tests()
	for _, v := range files {
		file, err := os.ReadFile("tests/step1/" + v)
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		data := string(file)
		data = strings.TrimSpace(data)
		lis := strings.Fields(data)
		if s.validate(lis) {
			fmt.Printf("successfully parsed\n")
		} else {
			err := errors.New("invalid json")
			fmt.Printf("%v\n", err)
			// os.Exit(1)
		}
	}
}

func tests() []string {
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString(' ')
	if err != nil && err != io.EOF {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	files := strings.Fields(line)
	return files
}
