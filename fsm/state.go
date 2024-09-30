// state.go contains the State struct and its methods
//
// Copyright (C) 2024 Goutham Krishna K V
package fsm

type State[T comparable] struct {
	name    T
	OnEnter func(TransitionContext[T])
	OnExit  func(TransitionContext[T])
}

func NewState[T comparable](
	name T,
	onEnterFunc func(TransitionContext[T]),
	onExitFunc func(TransitionContext[T]),
) State[T] {
	return State[T]{
		name:    name,
		OnEnter: onEnterFunc,
		OnExit:  onExitFunc,
	}
}

func (s *State[T]) Name() T {
	return s.name
}

func (s *State[T]) Copy() State[T] {
	return State[T]{name: s.name}
}
