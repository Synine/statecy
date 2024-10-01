// Transition changes the state of the FSM
//
// Copyright (C) 2024 Goutham Krishna K V
package fsm

type TransitionContext[ST comparable] struct {
	Name              ST
	From              ST
	To                ST
	CheckCurrentState bool
}

type Transition[ST comparable] struct {
	Name              ST
	From              *State[ST]
	To                *State[ST]
	CheckCurrentState bool
	Handler           func(*FSMContext[ST]) error
}

func (t *Transition[ST]) context() TransitionContext[ST] {
	return TransitionContext[ST]{
		Name:              t.Name,
		From:              t.From.Name(),
		To:                t.To.Name(),
		CheckCurrentState: t.CheckCurrentState,
	}
}
