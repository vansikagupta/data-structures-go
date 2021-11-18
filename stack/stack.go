package main

import (
	"errors"
	"fmt"
)

// write an stack interface which have 3 methods
// 1. push : args -> num, return -> no return
// 2. pop : args -> no args, return -> num
// 3. size : args -> no args, return -> total no of elements present in stack

//stack interface
type Stack interface {
	size() int
	pop() (interface{}, error)
	push(interface{}) error
}

// implement this interface for int and Employee enum data type
type IntStack struct {
	stack []int
	limit int
}

//IntStack implements Stack interface
func (i IntStack) size() int {
	return len(i.stack)
}

func (i *IntStack) push(v interface{}) error {
	if i.size() == i.limit {
		return errors.New("Cannot push more elements: Stack limit reached")
	}
	//Type assertion; getting underlying value of an interface
	val := v.(int)
	i.stack = append(i.stack, val)
	return nil
}

func (i *IntStack) pop() (interface{}, error) {
	if i.size() == 0 {
		return nil, errors.New("Cannot pop elements: Stack is empty")
	}
	v := i.stack[len(i.stack)-1]
	i.stack = i.stack[:len(i.stack)-1]
	return v, nil
}

type Employee struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

type EmployeeStack []Employee

//Employee stack implements Stack interface
func (s EmployeeStack) size() int {
	return len(s)
}

func (s *EmployeeStack) push(v interface{}) error {
	val := v.(Employee)
	*s = append(*s, val)
	return nil
}

func (s *EmployeeStack) pop() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("Cannot pop elements: Stack is empty")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}

func main() {

	var stackAPI Stack = &IntStack{limit: 2}
	fmt.Println("Stack size: ", stackAPI.size())

	stackAPI.push(1)
	stackAPI.push(2)
	fmt.Println("Stack size: ", stackAPI.size())

	if err := stackAPI.push(3); err != nil {
		fmt.Println(err)
	}

	fmt.Println(stackAPI.pop())
	fmt.Println(stackAPI.pop())
	fmt.Println("Stack size: ", stackAPI.size())

	_, err := stackAPI.pop()
	if err != nil {
		fmt.Println(err)
	}

	/*********************************/
	fmt.Println("Demonstration on Employee Stack")
	var empStack EmployeeStack
	stackAPI = &empStack

	fmt.Println("Stack size: ", stackAPI.size())

	stackAPI.push(Employee{"E00", "Jon", 200})
	stackAPI.push(Employee{"E01", "Jake", 250})
	fmt.Println("Stack size: ", stackAPI.size())

	fmt.Println(stackAPI.(*EmployeeStack))
}
