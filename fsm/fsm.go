// fsm.go contains all methods related to FSM implementation
//
// NOTE: This is a thin FSM, and would only provide you the tools required to create
// FSMs, along with transitions and handlers.
//
// Copyright (C) 2024 Goutham Krishna K V
package fsm

import (
	"errors"
	"sync"
)

type FSM[ST comparable] struct {
	mutex        sync.RWMutex
	states       map[ST]State[ST]
	initialState *State[ST]
	currentState *State[ST]
	transitions  map[ST]Transition[ST]
}

func NewSimpleFSM[ST comparable](params InitializationParams[ST]) (*FSM[ST], error) {
	fsmStates := make(map[ST]State[ST])
	var initialState *State[ST]

	for _, fsmState := range params.States {
		newState := NewState(fsmState, params.OnEnterfunc, params.OnExitFunc)
		if params.Initial == fsmState {
			initialState = &newState
		}
		fsmStates[fsmState] = newState
	}

	if initialState == nil {
		return nil, errors.New("initial state not found")
	}

	return &FSM[ST]{
		states:       fsmStates,
		initialState: initialState,
		currentState: nil,
	}, nil
}

// -- methods --

func (f *FSM[ST]) Initialize() error {
	if f.currentState != nil {
		return errors.New("fsm already initialized")
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.currentState = f.initialState
	return nil
}

// - methods after initialization -

// Transition changes the state of the FSM
func (f *FSM[ST]) Transition(transition ST) error {
	if f.currentState == nil {
		return errors.New("fsm not initialized")
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()

	if transition, transitionExists := f.transitions[transition]; !transitionExists {
		return errors.New("transition not found")
	} else if transition.From != nil &&
		transition.From.Name() == f.currentState.Name() {

		if transition.Handler != nil {
			// TODO: write a better error-handler, with support to
			// return the following:
			// 1. error-code
			// 2. error-messages
			// 3. arguments (error-data)
			// 4. state-change override (if required)
			err := transition.Handler(f.Context(transition))
			if err != nil {
				return err
			}
		}

		// transitions only if handler is successful

		if transition.From.OnExit != nil {
			transition.From.OnExit(transition.Context())
		}
		f.currentState = transition.To
		if transition.To.OnEnter != nil {
			transition.To.OnEnter(transition.Context())
		}
	}

	return nil
}

// - methods before initialization -
func (f *FSM[ST]) AddTransition(name ST, from ST, to ST, handler func(*FSMContext[ST]) error) error {
	if f.currentState != nil {
		return errors.New("fsm already initialized, cannot add transition")
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	if f.transitions == nil {
		f.transitions = make(map[ST]Transition[ST])
	}

	fromState, fromStateExists := f.states[from]
	toState, toStateExists := f.states[to]

	if !fromStateExists {
		return errors.New("from state does not exist")
	}
	if !toStateExists {
		return errors.New("to state does not exist")
	}

	f.transitions[name] = Transition[ST]{
		Name:    name,
		From:    &fromState,
		To:      &toState,
		Handler: handler,
	}
	return nil
}

func (f *FSM[ST]) Context(transition Transition[ST]) *FSMContext[ST] {
	// copy the current state and initial state
	initialState := f.currentState.Copy()
	currentState := f.currentState.Copy()
	// return the FSM context for handler
	return &FSMContext[ST]{
		InitialState: &initialState,
		CurrentState: &currentState,
		states:       f.states,
		Transition:   transition.Context(),
	}
}

// -- getters --

func (f *FSM[ST]) IsInitialized() bool {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return f.currentState != nil
}

// InitialState returns the initial state
func (f *FSM[ST]) InitialState() ST {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return f.initialState.Name()
}

// CurrentState returns the state the FSM is currently in
func (f *FSM[ST]) CurrentState() ST {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return f.currentState.Name()
}
