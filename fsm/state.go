// state.go contains the State struct and its methods
//
// Copyright (C) 2024 Goutham Krishna K V
package fsm

type State[T comparable] struct {
	name T
}

func NewState[T comparable](name T) State[T] {
	return State[T]{name: name}
}

func (s *State[T]) Name() T {
	return s.name
}

func (s *State[T]) Copy() State[T] {
	return State[T]{name: s.name}
}
