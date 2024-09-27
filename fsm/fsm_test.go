// Tests for the FSM package
//
// Copyright (C) 2024 Goutham Krishna K V
package fsm

import (
	"errors"
	"testing"
)

func TestFSM(t *testing.T) {
	{
		// success case
		successFsmConfig := InitializationParams[string]{
			States:  []string{"foo", "bar", "baz"},
			Initial: "foo",
		}
		// create a new FSM
		fsm, fsmErr := NewFSM(successFsmConfig)
		if fsmErr != nil {
			t.Errorf("creation error: %v", fsmErr)
		}
		if initErr := fsm.Initialize(); initErr != nil {
			t.Errorf("initialization error: %v", initErr)
		}
	}

	// failure case
	{
		failureFsmConfig := InitializationParams[string]{
			States:  []string{"foo", "bar", "baz"},
			Initial: "qux",
		}
		// create a new FSM
		_, fsmErr := NewFSM(failureFsmConfig)
		if fsmErr == nil {
			t.Errorf("creation error not caught!")
		}

	}
}

func TestFSMInitialization(t *testing.T) {
	{
		// success case
		successFsmConfig := InitializationParams[string]{
			States:  []string{"foo", "bar", "baz"},
			Initial: "foo",
		}
		// create a new FSM
		fsm, fsmErr := NewFSM(successFsmConfig)
		if fsmErr != nil {
			t.Errorf("creation error: %v", fsmErr)
		}
		if initErr := fsm.Initialize(); initErr != nil {
			t.Errorf("initialization error: %v", initErr)
		}
		if initErr := fsm.Initialize(); initErr == nil {
			t.Errorf("no initialization error received, unexpected")
		} else {
			t.Logf("initialization error, expected: %v", initErr)
		}
	}
}

func TestFSMUnknownTransition(t *testing.T) {
	successFsmConfig := InitializationParams[string]{
		States:  []string{"foo", "bar", "baz"},
		Initial: "bar",
	}

	fsm, fsmErr := NewFSM(successFsmConfig)
	if fsmErr != nil {
		t.Errorf("creation error: %v", fsmErr)
	}
	if initErr := fsm.Initialize(); initErr != nil {
		t.Errorf("initialization error: %v", initErr)
	}
	if !fsm.IsInitialized() {
		t.Errorf("fsm not initialized, expected initialized")
	}
	if transitionErr := fsm.Transition("yug"); transitionErr == nil {
		t.Errorf("transition error not caught!")
	}
}

func TestFSMTransition(t *testing.T) {
	successFsmConfig := InitializationParams[string]{
		States:  []string{"foo", "bar", "baz"},
		Initial: "bar",
	}

	fsm, fsmErr := NewFSM(successFsmConfig)
	if fsmErr != nil {
		t.Errorf("creation error: %v", fsmErr)
	}
	if addTransitionErr := fsm.AddTransition("yug", "bar", "foo", func(f *FSMContext[string]) error {
		t.Log("transitioning from bar to foo")
		return nil
	}); addTransitionErr != nil {
		t.Errorf("add transition error: %v", addTransitionErr)
	}
	if transitionErr := fsm.Transition("yug"); transitionErr == nil {
		t.Errorf("transition error not received, unexpected")
	} else {
		t.Logf("transition error expected: %v", transitionErr)
	}
	if initErr := fsm.Initialize(); initErr != nil {
		t.Errorf("initialization error: %v", initErr)
	}
	if !fsm.IsInitialized() {
		t.Errorf("fsm not initialized, expected initialized")
	}
	if fsm.CurrentState() != fsm.InitialState() {
		t.Errorf("current state is not bar")
	}
	if transitionErr := fsm.Transition("yug"); transitionErr != nil {
		t.Errorf("transition error: %v", transitionErr)
	}
	if fsm.CurrentState() == fsm.InitialState() {
		t.Errorf("current state is not foo")
	}
}

func TestFSMErrorTransition(t *testing.T) {
	successFsmConfig := InitializationParams[string]{
		States:  []string{"foo", "bar", "baz"},
		Initial: "bar",
	}

	fsm, fsmErr := NewFSM(successFsmConfig)
	if fsmErr != nil {
		t.Errorf("creation error: %v", fsmErr)
	}
	if addTransitionErr := fsm.AddTransition("yug", "bar", "foo", func(f *FSMContext[string]) error {
		return errors.New("lorem ipsum")
	}); addTransitionErr != nil {
		t.Errorf("add transition error: %v", addTransitionErr)
	}
	if initErr := fsm.Initialize(); initErr != nil {
		t.Logf("initialization error expected: %v", initErr)
	}
	if !fsm.IsInitialized() {
		t.Errorf("fsm not initialized, expected initialized")
	}
	if fsm.CurrentState() != "bar" {
		t.Errorf("current state is not bar")
	}
	if transitionErr := fsm.Transition("yug"); transitionErr != nil {
		t.Logf("transition error expected: %v", transitionErr)
	}
	if fsm.CurrentState() != "bar" {
		t.Errorf("current state changed from bar, unexpected")
	}
}
