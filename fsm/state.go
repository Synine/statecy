// state.go contains the State struct and its methods
//
// Copyright (C) 2024 Goutham Krishna K V
package fsm

type State[T comparable] struct {
	name    T
	onEnter func(TransitionContext[T])
	onExit  func(TransitionContext[T])
}

func NewState[T comparable](
	name T,
	onEnterFunc func(TransitionContext[T]),
	onExitFunc func(TransitionContext[T]),
) State[T] {
	return State[T]{
		name:    name,
		onEnter: onEnterFunc,
		onExit:  onExitFunc,
	}
}

func (s *State[T]) Name() T {
	return s.name
}

func (s *State[T]) Copy() State[T] {
	return State[T]{name: s.name}
}
