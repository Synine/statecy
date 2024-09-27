// FSMContext is the context to be sent to transition function to decide the next state
//
// Copyright (C) 2024 Goutham Krishna K V
package fsm

// FSM Context to be sent to transition function to decide the next state
type FSMContext[ST comparable] struct {
	Transition   TransitionContext[ST]
	CurrentState *State[ST]
	InitialState *State[ST]
	states       map[ST]State[ST]
}

func (fsmc *FSMContext[ST]) GetState(stateRef ST) (State[ST], bool) {
	state, ok := fsmc.states[stateRef]
	if !ok {
		return State[ST]{}, false
	}
	return state.Copy(), true
}
