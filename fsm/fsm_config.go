// FSM is a simple finite state machine implementation
//
// This file contains the FSM configuration and context required for transition
// (both current-state and initial-state is copies, but the states are a shared value)
//
// Copyright (C) 2024 Goutham Krishna K V
package fsm

// InitializationParams set up the FSM with the states and the initial state. Consider
// this ONLY the initialization methods, not the
type InitializationParams[T comparable] struct {
	// States is a list of states that the FSM is allowed to be in
	States  []T
	Initial T
}
