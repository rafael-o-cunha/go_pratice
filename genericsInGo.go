/**
*	GENERICS PERMITEM CRIAR COLLECTIONS PERSONALIZADAS QUE FUNCIONAM COM
*	DIFERENTES TIPOS.
*	EXEMPLO: UM STACK OU UMA LISTA GENÉRICA.
*/

package main

import "fmt"

type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

func genericsInGo() {
	var stack Stack[int]
    stack.Push(10)
    stack.Push(20)
    fmt.Println(stack.Pop()) // 20
    fmt.Println(stack.Pop()) // 10
}

